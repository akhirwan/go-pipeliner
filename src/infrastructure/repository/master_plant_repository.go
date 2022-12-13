package repository

import (
	"go-pipeliner/src/domain/entity"

	"github.com/jmoiron/sqlx"
	"github.com/k0kubun/pp"
)

type masterPlantRepository struct {
	DB *sqlx.DB
}

func NewMasterPlantRepository(db *sqlx.DB) *masterPlantRepository {
	return &masterPlantRepository{DB: db}
}

func (s *masterPlantRepository) FindAll() ([]*entity.SalesPlantEntity, error) {
	plants := []*entity.SalesPlantEntity{}
	err := s.DB.Select(&plants, "SELECT * FROM sales_plant ORDER BY plant_secondary_id, id ASC")
	if err != nil {
		pp.Println(err)
		return nil, err
	}

	return plants, nil
}
