package repository

import (
	"fmt"
	"go-pipeliner/src/domain/model"
	"log"
	"sync"
	"time"

	"github.com/jmoiron/sqlx"
)

// Global repo for any source system
type salesCustomerRepository struct {
	DB *sqlx.DB
}

func NewSalesCustomerRepository(db *sqlx.DB) *salesCustomerRepository {
	return &salesCustomerRepository{DB: db}
}

func (s *salesCustomerRepository) Insert(data []*model.SalesCustomerModel) error {
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

func (s *salesCustomerRepository) batchInsert(
	id int,
	data []*model.SalesCustomerModel,
	ch chan<- error,
	wg *sync.WaitGroup) {

	defer wg.Done()

	_, err := s.DB.NamedExec(`
		INSERT INTO sales_customer (
			id,
			company_id, 
			source_id, 
			plant_id, 
			name, 
			alias_name, 
			customer_type_id,
			street,
			city, 
			key_account_id, 
			key_account_group_id,
			top_id, 
			status, 
			is_deleted, 
			create_date,
			price_group_id,
			insert_date
		)
		VALUES(
			:id,
			:company_id, 
			:source_id, 
			:plant_id, 
			:name, 
			:alias_name, 
			:customer_type_id,
			:street,
			:city, 
			:key_account_id, 
			:key_account_group_id,
			:top_id, 
			:status, 
			:is_deleted, 
			:create_date,
			:price_group_id,
			:insert_date
		) 
		ON DUPLICATE KEY UPDATE 
			source_id 				= VALUES(source_id),
			name 					= VALUES(name),
			alias_name 				= VALUES(alias_name),
			customer_type_id 		= VALUES(customer_type_id),
			street 					= VALUES(street),
			city 					= VALUES(city),
			key_account_id 			= VALUES(key_account_id),
			key_account_group_id 	= VALUES(key_account_group_id),
			top_id 					= VALUES(top_id),
			status 					= VALUES(status),
			is_deleted 				= VALUES(is_deleted),
			create_date 			= VALUES(create_date),
			price_group_id 			= VALUES(price_group_id)
		`, data)

	ch <- err
}

func (s *salesCustomerRepository) FindAllByCreateDate(periodCreateDate time.Time) ([]*model.SalesCustomerModel, error) {
	var data []*model.SalesCustomerModel

	err := s.DB.Select(
		&data,
		`SELECT 
			id,
			company_id,
			source_id,
			plant_id,
			name,
			alias_name,
			customer_type_id,
			street,
			city,
			key_account_id,
			key_account_group_id,
			price_group_id,
			top_id,
			status,
			is_deleted,
			LEFT(create_date, 10) create_date,
			insert_date
		FROM 
			sales_customer 
		WHERE 
			LEFT(create_date, 7) = ? 
		ORDER BY id;`,
		periodCreateDate.Format("2006-01"))

	if err != nil {
		log.Printf("[FATAL] get sales_customer error: %s", err.Error())
		return nil, err
	}

	return data, nil
}
