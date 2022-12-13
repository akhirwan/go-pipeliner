package worker

import (
	"go-pipeliner/src/domain/sales"
	"log"
	"sync"
	"time"
)

type SalesDatamartDailyJob struct {
	PlantID string
	Period  time.Time
}

type prodToDevSalesDatamartDailyWorkerPool struct {
	WorkerCap int
	JobsCh    chan SalesDatamartDailyJob
	DoneCh    chan bool
	Service   *sales.ProdToDevSalesDatamartDailyService
}

func NewProdToDevSalesDatamartDailyWorkerPool(workerCap int, service *sales.ProdToDevSalesDatamartDailyService) *prodToDevSalesDatamartDailyWorkerPool {
	pool := &prodToDevSalesDatamartDailyWorkerPool{
		WorkerCap: workerCap,
		JobsCh:    make(chan SalesDatamartDailyJob, workerCap),
		DoneCh:    make(chan bool),
		Service:   service,
	}

	go pool.createPool()

	return pool
}

func (s *prodToDevSalesDatamartDailyWorkerPool) createPool() {
	var wg sync.WaitGroup

	for i := 0; i < s.WorkerCap; i++ {
		wg.Add(1)
		go s.doWork(i+1, s.JobsCh, &wg)
	}

	log.Println("Worker pools is ready with", s.WorkerCap, " workers.")

	wg.Wait()

	s.DoneCh <- true
}

func (s *prodToDevSalesDatamartDailyWorkerPool) doWork(id int, jobCh <-chan SalesDatamartDailyJob, wg *sync.WaitGroup) {

	log.Printf("Workers %v is running", id)
	// fmt.Println(jobCh)

	for j := range jobCh {
		log.Printf("🔴 [%v][%s][%s] Receive Job\n", id, j.PlantID, j.Period.Format("200601"))

		// We don't care the return here yet
		if err := s.Service.LoadSales(j.PlantID, j.Period); err != nil {
			log.Printf("❎ [%v][%s][%s] Done Job With Error %s\n", id, j.PlantID, j.Period.Format("200601"), err.Error())
		} else {
			log.Printf("✅ [%v][%s][%s] Done Job\n", id, j.PlantID, j.Period.Format("200601"))
		}
	}

	wg.Done()

	log.Printf("🏁 Workers %v stopped", id)
}
