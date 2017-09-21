package core

import (
	httpPegasus "github.com/kbfu/pegasus/components/http"
	"time"
)

func InitWorkerPool(jobs chan func(), workers int) {
	for w := 0; w < workers; w++ {
		go worker(jobs)
	}
}

func worker(jobs chan func()) {
	for j := range jobs {
		go j()
	}
}

func InitJobs(tasks int, rate int, jobs chan func(), r *httpPegasus.RequestData, results chan map[string]interface{}) {
	limiter := time.Tick(time.Duration(float64(time.Second) / float64(rate)))
	for j := 0; j < tasks; j++ {
		<-limiter
		jobs <- func() {
			r.Request(*Client, results)
		}
	}
}
