package interfaces

import (
	"go-pipeliner/src/domain/model"
	"time"
)

type DevSalesDatamartDailyRepositoryInterface interface {
	Insert(data []*model.SalesDatamartDailyModel) error
}

type ProdSalesDatamartDailyRepositoryInterface interface {
	FindAllByBillingDate(plantID string, billingDateFrom, billingDateTo time.Time) ([]*model.SalesDatamartDailyModel, error)
	FindAllByPeriod(plantID string, period time.Time) ([]*model.SalesDatamartDailyModel, error)
}
