package task

import (
	"go-pipeliner/src/app/task/worker"
	"go-pipeliner/src/domain/entity"
	"go-pipeliner/src/domain/sales"
	"go-pipeliner/src/infrastructure/config"
	"go-pipeliner/src/infrastructure/repository"
	"go-pipeliner/src/infrastructure/repository/db"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/k0kubun/pp"
)

type prodToDevSalesDatamartDailyPipelineTask struct {
	DBPROD *sqlx.DB
	DBDEV  *sqlx.DB
}

func NewProdToDevSalesDatamartDailyPipelineTask(config config.Config) *prodToDevSalesDatamartDailyPipelineTask {
	return &prodToDevSalesDatamartDailyPipelineTask{
		DBPROD: db.NewMySQLDBConnection(createDWHProdConfig(config)),
		DBDEV:  db.NewMySQLDBConnection(createDWHDevConfig(config)),
	}
}

func (s *prodToDevSalesDatamartDailyPipelineTask) Execute(periodFrom, periodTo time.Time) {
	plants := []*entity.SalesPlantEntity{} // all plants
	// plants = append(plants, &entity.SalesPlantEntity{ID: "P101", Name: "Test Plant"})

	s.executeInParallel(plants, periodFrom, periodTo)

	defer s.DBPROD.Close()
	defer s.DBDEV.Close()

	log.Println("Pipeline execution done")
}

func (s *prodToDevSalesDatamartDailyPipelineTask) executeInParallel(plants []*entity.SalesPlantEntity, periodFrom, periodTo time.Time) {
	plantRepo := repository.NewMasterPlantRepository(s.DBPROD)
	prodRepo := repository.NewSalesDatamartDailyRepository(s.DBPROD)
	devRepo := repository.NewSalesDatamartDailyRepository(s.DBDEV)

	service := sales.NewProdToDevSalesDatamartDailyService(prodRepo, devRepo)
	pool := worker.NewProdToDevSalesDatamartDailyWorkerPool(10, service)

	var err error
	if len(plants) == 0 {
		plants, err = plantRepo.FindAll()
		if err != nil {
			pp.Println("Failed to get plant query")
			return
		}
	}

	pp.Println("total plant : ", len(plants))
	for _, plant := range plants {
		for period := periodFrom; !period.After(periodTo); period = period.AddDate(0, 1, 0) {
			pool.JobsCh <- worker.SalesDatamartDailyJob{
				PlantID: plant.ID,
				Period:  period,
			}
		}
	}

	close(pool.JobsCh)
	<-pool.DoneCh

	log.Println("All Workers closed")
}
