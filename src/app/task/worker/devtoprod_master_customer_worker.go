package worker

import (
	"go-pipeliner/src/domain/master"
	"log"
	"sync"
	"time"
)

type MasterCustomerJob struct {
	Period time.Time
}

type devToProdSalesCustomerWorkerPool struct {
	WorkerCap int
	JobsCh    chan MasterCustomerJob
	DoneCh    chan bool
	Service   *master.DevToProdMasterCustomerService
}

func NewDevToProdSalesCustomerWorkerPool(workerCap int, service *master.DevToProdMasterCustomerService) *devToProdSalesCustomerWorkerPool {
	pool := &devToProdSalesCustomerWorkerPool{
		WorkerCap: workerCap,
		JobsCh:    make(chan MasterCustomerJob, workerCap),
		DoneCh:    make(chan bool),
		Service:   service,
	}

	go pool.createPool()

	return pool
}

func (s *devToProdSalesCustomerWorkerPool) createPool() {
	var wg sync.WaitGroup

	for i := 0; i < s.WorkerCap; i++ {
		wg.Add(1)
		go s.doWork(i+1, s.JobsCh, &wg)
	}

	log.Println("Worker pools is ready with", s.WorkerCap, " workers.")

	wg.Wait()

	s.DoneCh <- true
}

func (s *devToProdSalesCustomerWorkerPool) doWork(id int, jobCh <-chan MasterCustomerJob, wg *sync.WaitGroup) {

	log.Printf("Workers %v is running", id)
	// fmt.Println(jobCh)

	for j := range jobCh {
		log.Printf("🔴 [%v][%s] Receive Job\n", id, j.Period.Format("200601"))

		// We don't care the return here yet
		if err := s.Service.LoadMasterCustomer(j.Period); err != nil {
			log.Printf("❎ [%v][%s] Done Job With Error %s\n", id, j.Period.Format("200601"), err.Error())
		} else {
			log.Printf("✅ [%v][%s] Done Job\n", id, j.Period.Format("200601"))
		}
	}

	wg.Done()

	log.Printf("🏁 Workers %v stopped", id)
}
