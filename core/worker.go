package core

import (
	httpPegasus "git.jiayincloud.com/TestDev/pegasus.git/components/http"
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

func InitJobs(rate int, jobs chan func(), r []httpPegasus.RequestData, results chan map[string]interface{}) {
	limiter := time.Tick(time.Duration(float64(time.Second) / float64(rate)))
	for _, v := range r {
		task := v
		<-limiter
		jobs <- func() {
			task.Request(*Client, results)
		}
	}
}
