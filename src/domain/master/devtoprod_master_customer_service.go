package master

import (
	"go-pipeliner/src/domain/interfaces"
	"log"
	"time"
)

type DevToProdMasterCustomerService struct {
	DevRepo  interfaces.DevSalesCustomerRepositoryInterface
	ProdRepo interfaces.ProdSalesCustomerRepositoryInterface
}

// Pass repository
func NewDevToProdMasterCustomerService(
	devRepo interfaces.DevSalesCustomerRepositoryInterface,
	prodRepo interfaces.ProdSalesCustomerRepositoryInterface,
) *DevToProdMasterCustomerService {
	return &DevToProdMasterCustomerService{
		DevRepo:  devRepo,
		ProdRepo: prodRepo,
	}
}

func (s *DevToProdMasterCustomerService) LoadMasterCustomer(periodCreateDate time.Time) error {
	log.Printf("[%s] Start Extract ...", periodCreateDate.Format("200601"))
	result, err := s.DevRepo.FindAllByCreateDate(periodCreateDate)
	if err != nil {
		log.Printf("[%s][ERROR] extract %s", periodCreateDate.Format("200601"), err.Error())
		return err
	}

	if len(result) == 0 {
		log.Printf("[%s][ERROR] empty data", periodCreateDate.Format("200601"))
		return nil
	}

	log.Printf("[%s] End Extract with %v rows", periodCreateDate.Format("200601"), len(result))

	log.Printf("[%s] Start Load with %v rows", periodCreateDate.Format("200601"), len(result))
	err = s.ProdRepo.Insert(result)
	if err != nil {
		log.Printf("[%s][ERROR] load %s", periodCreateDate.Format("200601"), err.Error())
		return err
	}

	log.Printf("[%s] End Load ", periodCreateDate.Format("200601"))

	return nil
}
