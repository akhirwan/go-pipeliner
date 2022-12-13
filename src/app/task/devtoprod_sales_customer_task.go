package task

import (
	"go-pipeliner/src/app/task/worker"
	"go-pipeliner/src/domain/master"
	"go-pipeliner/src/infrastructure/config"
	"go-pipeliner/src/infrastructure/repository"
	"go-pipeliner/src/infrastructure/repository/db"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
)

type devToProdSalesCustomerPipelineTask struct {
	DBDEV  *sqlx.DB
	DBPROD *sqlx.DB
}

func NewDevToProdSalesCustomerPipelineTask(config config.Config) *devToProdSalesCustomerPipelineTask {
	return &devToProdSalesCustomerPipelineTask{
		DBDEV:  db.NewMySQLDBConnection(createDWHDevConfig(config)),
		DBPROD: db.NewMySQLDBConnection(createDWHProdConfig(config)),
	}
}

func (s *devToProdSalesCustomerPipelineTask) Execute(createDateFrom, createDateTo time.Time) {
	s.executeInParallel(createDateFrom, createDateTo)

	defer s.DBDEV.Close()
	defer s.DBPROD.Close()

	log.Println("Pipeline execution done")
}

func (s *devToProdSalesCustomerPipelineTask) executeInParallel(createDateFrom, createDateTo time.Time) {
	devRepo := repository.NewSalesCustomerRepository(s.DBDEV)
	prodRepo := repository.NewSalesCustomerRepository(s.DBPROD)

	service := master.NewDevToProdMasterCustomerService(devRepo, prodRepo)
	pool := worker.NewDevToProdSalesCustomerWorkerPool(10, service)

	for period := createDateFrom; !period.After(createDateTo); period = period.AddDate(0, 1, 0) {
		pool.JobsCh <- worker.MasterCustomerJob{
			Period: period,
		}
	}

	close(pool.JobsCh)
	<-pool.DoneCh

	log.Println("All Workers closed")
}
