package model

import "time"

type SalesDatamartDailyModel struct {
	CompanyID              string    `db:"company_id"`
	SourceID               string    `db:"source_id"`
	PlantId                string    `db:"plant_id"`
	BillingDate            string    `db:"billing_date" time_format:"sql_date"`
	BillingNumber          string    `db:"billing_number"`
	Period                 string    `db:"period"`
	BillingTypeID          string    `db:"billing_type_id"`
	BillingType            string    `db:"billing_type"`
	SalesmanID             string    `db:"salesman_id"`
	CustomerID             string    `db:"customer_id"`
	CustomerTypeID         string    `db:"customer_type_id"`
	ItemNumber             int       `db:"item_number"`
	ProductID              string    `db:"product_id"`
	QtyInPCS               float64   `db:"qty_in_pcs"`
	QtyInCTN               float64   `db:"qty_in_ctn"`
	GrossValue             float64   `db:"gross_value"`
	RegularDiscountAmount  float64   `db:"regular_discount_amount"`
	CODDiscountAmount      float64   `db:"cod_discount_amount"`
	PCDiscountAmount       float64   `db:"pc_discount_amount"`
	BSDiscountAmount       float64   `db:"bs_discount_amount"`
	TotalDiscountAmount    float64   `db:"total_discount_amount"`
	IPTDiscountAmount      float64   `db:"ipt_discount_amount"`
	NonIPTDiscountAmount   float64   `db:"non_ipt_discount_amount"`
	TotalTPRDiscountAmount float64   `db:"total_tpr_discount_amount"`
	VATAmount              float64   `db:"vat_amount"`
	NetValue               float64   `db:"net_value"`
	CurrencyID             string    `db:"currency_id"`
	ReturnReasonID         string    `db:"return_reason_id"`
	IsCancelled            bool      `db:"is_cancelled"`
	CreateDate             string    `db:"create_date" time_format:"sql_date"`
	CreateTime             string    `db:"create_time"`
	InsertDate             time.Time `db:"insert_date"`
	ModifyDate             time.Time `db:"modify_date"`
}
