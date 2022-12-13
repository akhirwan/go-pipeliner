package repository

import (
	"fmt"
	"go-pipeliner/src/domain/model"
	"log"
	"sync"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/k0kubun/pp"
)

// Global repo for any source system
type salesDatamartDailyRepository struct {
	DB *sqlx.DB
}

func NewSalesDatamartDailyRepository(db *sqlx.DB) *salesDatamartDailyRepository {
	return &salesDatamartDailyRepository{DB: db}
}

func (s *salesDatamartDailyRepository) Insert(data []*model.SalesDatamartDailyModel) error {
	batchSize := 30 // Smaller batch is faster
	batchID := 0
	total := len(data) // align with maxconns of database

	ch := make(chan error, total)
	var wg sync.WaitGroup

	for i := 0; i < total; i += batchSize {
		j := i + batchSize
		if j > total {
			j = total
		}

		batchID++

		wg.Add(1)
		// log.Printf("Batch insert %v to %v of %v\n", i+1, j, total)
		go s.batchInsert(batchID, data[i:j], ch, &wg)

	}

	wg.Wait()
	close(ch)

	// read Values from the channel
	for err := range ch {
		if err != nil {
			return fmt.Errorf("[FATAL] query error: %s", err.Error()) // return from func not for loop
		}
	}

	return nil
}

func (s *salesDatamartDailyRepository) batchInsert(
	id int,
	data []*model.SalesDatamartDailyModel,
	ch chan<- error,
	wg *sync.WaitGroup) {

	defer wg.Done()

	_, err := s.DB.NamedExec(
		`INSERT INTO sales_datamart_daily (
			company_id, 
			source_id, 
			plant_id, 
			billing_date, 
			billing_number, 
			period, 
			billing_type_id, 
			billing_type, 
			salesman_id, 
			customer_id, 
			customer_type_id, 
			item_number, 
			product_id, 
			qty_in_pcs, 
			qty_in_ctn, 
			gross_value, 
			regular_discount_amount, 
			cod_discount_amount, 
			pc_discount_amount, 
			bs_discount_amount, 
			total_discount_amount, 
			ipt_discount_amount, 
			non_ipt_discount_amount, 
			total_tpr_discount_amount, 
			vat_amount, 
			net_value, 
			currency_id, 
			return_reason_id, 
			is_cancelled, 
			create_date, 
			insert_date,
			modify_date
		) VALUES (
			:company_id, 
			:source_id, 
			:plant_id, 
			:billing_date, 
			:billing_number, 
			:period, 
			:billing_type_id, 
			:billing_type, 
			:salesman_id, 
			:customer_id, 
			:customer_type_id, 
			:item_number, 
			:product_id, 
			:qty_in_pcs, 
			:qty_in_ctn, 
			:gross_value, 
			:regular_discount_amount, 
			:cod_discount_amount, 
			:pc_discount_amount, 
			:bs_discount_amount, 
			:total_discount_amount, 
			:ipt_discount_amount, 
			:non_ipt_discount_amount, 
			:total_tpr_discount_amount, 
			:vat_amount, 
			:net_value, 
			:currency_id, 
			:return_reason_id, 
			:is_cancelled, 
			:create_date, 
			:insert_date,
			:modify_date
		) ON DUPLICATE KEY UPDATE
		company_id					= VALUES(company_id), 
		source_id					= VALUES(source_id), 
		plant_id					= VALUES(plant_id), 
		billing_date				= VALUES(billing_date), 
		billing_number				= VALUES(billing_number), 
		period						= VALUES(period), 
		billing_type_id				= VALUES(billing_type_id), 
		billing_type				= VALUES(billing_type), 
		salesman_id					= VALUES(salesman_id), 
		customer_id					= VALUES(customer_id), 
		customer_type_id			= VALUES(customer_type_id), 
		item_number					= VALUES(item_number), 
		product_id					= VALUES(product_id), 
		qty_in_pcs					= VALUES(qty_in_pcs), 
		qty_in_ctn					= VALUES(qty_in_ctn), 
		gross_value					= VALUES(gross_value), 
		regular_discount_amount		= VALUES(regular_discount_amount), 
		cod_discount_amount			= VALUES(cod_discount_amount), 
		pc_discount_amount			= VALUES(pc_discount_amount), 
		bs_discount_amount			= VALUES(bs_discount_amount), 
		total_discount_amount		= VALUES(total_discount_amount), 
		ipt_discount_amount			= VALUES(ipt_discount_amount), 
		non_ipt_discount_amount		= VALUES(non_ipt_discount_amount), 
		total_tpr_discount_amount	= VALUES(total_tpr_discount_amount), 
		vat_amount					= VALUES(vat_amount), 
		net_value					= VALUES(net_value), 
		currency_id					= VALUES(currency_id), 
		return_reason_id			= VALUES(return_reason_id), 
		is_cancelled				= VALUES(is_cancelled), 
		create_date					= VALUES(create_date),
		modify_date					= VALUES(modify_date);`,
		data)

	ch <- err
}

func (s *salesDatamartDailyRepository) FindAllByBillingDate(plantID string, billingDateFrom, billingDateTo time.Time) ([]*model.SalesDatamartDailyModel, error) {
	var data []*model.SalesDatamartDailyModel

	err := s.DB.Select(
		&data,
		`SELECT 
			company_id,
			source_id,
			plant_id,
			LEFT(billing_date, 10) billing_date,
			billing_number,
			period,
			billing_type_id,
			billing_type,
			salesman_id,
			customer_id,
			customer_type_id,
			item_number,
			product_id,
			qty_in_pcs,
			qty_in_ctn,
			gross_value,
			regular_discount_amount,
			cod_discount_amount,
			pc_discount_amount,
			bs_discount_amount,
			total_discount_amount,
			ipt_discount_amount,
			non_ipt_discount_amount,
			total_tpr_discount_amount,
			vat_amount,
			net_value,
			currency_id,
			return_reason_id,
			is_cancelled,
			LEFT(create_date, 10) create_date,
			insert_date,
			modify_date
		FROM 
			sales_datamart_daily 
		WHERE 
			plant_id = ?
			AND LEFT(billing_date, 10) BETWEEN ? AND ?;`,
		plantID,
		billingDateFrom.Format("2006-01-02"),
		billingDateTo.Format("2006-01-02"))

	if err != nil {
		log.Printf("[FATAL] get sales_datamart_daily error: %s", err.Error())
		return nil, err
	}

	return data, nil
}

func (s *salesDatamartDailyRepository) FindAllByPeriod(plantID string, period time.Time) ([]*model.SalesDatamartDailyModel, error) {
	var data []*model.SalesDatamartDailyModel

	// pp.Println(s)
	pp.Println(plantID, period.Format("200601"))
	err := s.DB.Select(
		&data,
		`SELECT 
			company_id,
			source_id,
			plant_id,
			LEFT(billing_date, 10) billing_date,
			billing_number,
			period,
			billing_type_id,
			billing_type,
			salesman_id,
			customer_id,
			customer_type_id,
			item_number,
			product_id,
			qty_in_pcs,
			qty_in_ctn,
			gross_value,
			regular_discount_amount,
			cod_discount_amount,
			pc_discount_amount,
			bs_discount_amount,
			total_discount_amount,
			ipt_discount_amount,
			non_ipt_discount_amount,
			total_tpr_discount_amount,
			vat_amount,
			net_value,
			currency_id,
			return_reason_id,
			is_cancelled,
			LEFT(create_date, 10) create_date,
			insert_date,
			modify_date
		FROM 
			sales_datamart_daily 
		WHERE 
			plant_id = ?
			AND period = ?;`,
		plantID,
		period.Format("200601"))

	if err != nil {
		log.Printf("[FATAL] get sales_datamart_daily error: %s", err.Error())
		return nil, err
	}

	pp.Println(len(data))

	return data, nil
}
