package services

import (
	"github.com/gin-gonic/gin"
	"github.com/kbfu/pegasus/core"
	pegasusHttp "github.com/kbfu/pegasus/components/http"
	"net/http"
	"fmt"
)

var ammos = []pegasusHttp.RequestData{}

func Load(c *gin.Context)  {
	var data pegasusHttp.RequestData
	err := c.BindJSON(&data)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("%s", err))
		return
	}
	ammos = append(ammos, data)
}

func Fire(c *gin.Context)  {
	defer func() {
		ammos = ammos[:0]
	}()
	for _, v := range ammos {
		var (
			body        string
			queryParams map[string]string
			pathParams  []string
			headers     map[string]string
			file        map[string]string
			form        map[string]string
		)
		url := v.Url
		workers := v.Workers
		duration := v.Duration
		rate := v.Rate
		method := v.Method
		if v.Headers != nil {
			headers = v.Headers
		}
		if v.QueryParams != nil {
			queryParams = v.QueryParams
		}
		if v.Body != "" {
			body = v.Body
		}
		if v.PathParams != nil {
			pathParams = v.PathParams
		}
		if v.File != nil {
			file = v.File
		}
		if v.Form != nil {
			form = v.Form
		}

		tasks := duration * rate
		jobs := make(chan func(), workers)
		results := make(chan map[string]interface{}, tasks)

		r := pegasusHttp.RequestData{
			Url:         url,
			Method:      method,
			Body:        body,
			QueryParams: queryParams,
			PathParams:  pathParams,
			Headers:     headers,
			File:        file,
			Form:        form,
		}

		core.InitWorkerPool(jobs, workers)
		core.InitJobs(tasks, rate, jobs, &r, results)

		for a := 0; a < tasks; a++ {
			<-results
			//fmt.Println(<-results)
		}
	}

}
