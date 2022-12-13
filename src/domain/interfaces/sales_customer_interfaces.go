package interfaces

import (
	"go-pipeliner/src/domain/model"
	"time"
)

type DevSalesCustomerRepositoryInterface interface {
	FindAllByCreateDate(periodCreateDate time.Time) ([]*model.SalesCustomerModel, error)
}

type ProdSalesCustomerRepositoryInterface interface {
	Insert(data []*model.SalesCustomerModel) error
}
