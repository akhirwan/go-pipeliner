package entity

import "time"

type SalesPlantEntity struct {
	ID               string    `db:"id"`
	CompanyID        string    `db:"company_id"`
	PlantSecondaryID string    `db:"plant_secondary_id"`
	SourceID         string    `db:"source_id"`
	Name             string    `db:"name"`
	ProfitCenterID   string    `db:"profit_center_id"`
	ProfitCenterName string    `db:"profit_center_name"`
	SalesOrgID       string    `db:"sales_org_id"`
	AreadID          string    `db:"area_id"`
	City             string    `db:"city"`
	Province         string    `db:"province"`
	InsertDate       time.Time `db:"insert_date"`
}

type SalesPlantAttributeEntity struct {
	CompanyID       string    `db:"company_id"`
	PlantID         string    `db:"plant_id"`
	PlantCustomerID string    `db:"plant_customer_id"`
	PlantTypeID     string    `db:"plant_type_id"`
	PlantGroup      string    `db:"plant_group"`
	Region          string    `db:"region"`
	RegionGroup     string    `db:"region_group"`
	Island          string    `db:"island"`
	GeoLatitude     string    `db:"geo_latitude"`
	GeoLongitude    string    `db:"geo_longitude"`
	IsActive        string    `db:"is_active"`
	IsDirectSelling string    `db:"is_direct_selling"`
	InsertDate      time.Time `db:"insert_date"`
}

type SalesPlantBOneEntity struct {
	CompanyID      string    `db:"company_id"`
	SourceID       string    `db:"source_id"`
	PlantID        string    `db:"plant_id"`
	PlantName      string    `db:"plant_name"`
	WarehouseID    string    `db:"warehouse_id"`
	ProfitCenterID string    `db:"profit_center_id"`
	SalesOrgID     string    `db:"sales_org_id"`
	AreadID        string    `db:"area_id"`
	City           string    `db:"city"`
	Province       string    `db:"province"`
	InsertDate     time.Time `db:"insert_date"`
}
