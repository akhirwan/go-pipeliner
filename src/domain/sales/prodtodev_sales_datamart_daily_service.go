package sales

import (
	"go-pipeliner/src/domain/interfaces"
	"go-pipeliner/src/domain/utility"
	"log"
	"time"
)

type ProdToDevSalesDatamartDailyService struct {
	ProdRepo interfaces.ProdSalesDatamartDailyRepositoryInterface
	DevRepo  interfaces.DevSalesDatamartDailyRepositoryInterface
}

// Pass repository
func NewProdToDevSalesDatamartDailyService(
	prodRepo interfaces.ProdSalesDatamartDailyRepositoryInterface,
	devRepo interfaces.DevSalesDatamartDailyRepositoryInterface,
) *ProdToDevSalesDatamartDailyService {
	return &ProdToDevSalesDatamartDailyService{
		DevRepo:  devRepo,
		ProdRepo: prodRepo,
	}
}

func (s *ProdToDevSalesDatamartDailyService) LoadSales(plantID string, period time.Time) error {
	log.Printf("[%s][%s] Start Extract ...", plantID, period.Format("200601"))

	billingDateFrom, billingDateTo := utility.FirstLastInPeriod(period)
	result, err := s.ProdRepo.FindAllByBillingDate(plantID, billingDateFrom, billingDateTo)
	if err != nil {
		log.Printf("[%s][%s][ERROR] extract %s", plantID, period.Format("200601"), err.Error())
		return err
	}

	if len(result) == 0 {
		log.Printf("[%s][%s][ERROR] empty data", plantID, period.Format("200601"))
		return nil
	}

	log.Printf("[%s][%s] End Extract with %v rows", plantID, period.Format("200601"), len(result))

	log.Printf("[%s][%s] Start Load with %v rows", plantID, period.Format("200601"), len(result))
	err = s.DevRepo.Insert(result)
	if err != nil {
		log.Printf("[%s][%s][ERROR] load %s", plantID, period.Format("200601"), err.Error())
		return err
	}

	log.Printf("[%s][%s] End Load ", plantID, period.Format("200601"))

	return nil
}
