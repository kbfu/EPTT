package core

import (
	httpPegasus "github.com/kbfu/pegasus/components/http"
	"time"
)

func InitWorkerPool(jobs chan func(), rate int, workers int) {
	for w := 0; w < workers; w++ {
		go worker(jobs, rate, workers)
	}
}

func worker(jobs chan func(), rate int, workers int) {
	limiter := time.Tick(time.Second / time.Duration(rate/workers))
	for j := range jobs {
		<-limiter
		go j()
	}
}

func InitJobs(tasks int, jobs chan func(), r *httpPegasus.RequestData, results chan map[string]interface{}) {
	for j := 0; j < tasks; j++ {
		jobs <- func() {
			r.Request(results)
		}
	}
}
