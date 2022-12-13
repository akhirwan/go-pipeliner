package model

import "time"

type SalesCustomerModel struct {
	ID                string    `json:"customer_id" db:"id"`
	CompanyID         string    `json:"company_id" db:"company_id"`
	SourceID          string    `db:"source_id"`
	PlantID           string    `json:"plant_id" db:"plant_id"`
	Name              string    `json:"customer_name" db:"name"`
	AliasName         string    `json:"customer_name2" db:"alias_name"`
	CustomerTypeID    string    `json:"customer_group_id" db:"customer_type_id"`
	Street            string    `json:"street" db:"street"`
	City              string    `json:"city" db:"city"`
	KeyAccountID      string    `json:"key_account_id" db:"key_account_id"`
	KeyAccountGroupID string    `json:"key_account_group_id" db:"key_account_group_id"`
	TopID             string    `json:"top_id" db:"top_id"`
	Status            string    `json:"status" db:"status"`
	IsDeleted         bool      `json:"is_deleted" db:"is_deleted"`
	CreateDate        string    `json:"create_date" db:"create_date" time_format:"sql_date"`
	PriceGroupID      string    `json:"price_group" db:"price_group_id"`
	InsertDate        time.Time `db:"insert_date"`
}
