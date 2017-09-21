package services

import (
	"github.com/gin-gonic/gin"
	"github.com/kbfu/pegasus/core"
	pegasusHttp "github.com/kbfu/pegasus/components/http"
	"net/http"
	"fmt"
)

func Overload(c *gin.Context)  {
	var data pegasusHttp.RequestData
	err := c.BindJSON(&data)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("%s", err))
		return
	}
			var (
				body        string
				queryParams map[string]string
				pathParams  []string
				headers     map[string]string
				file        map[string]string
				form        map[string]string
			)
			url := data.Url
			workers := data.Workers
			duration := data.Duration
			rate := data.Rate
			method := data.Method
			if data.Headers != nil {
				headers = data.Headers
			}
			if data.QueryParams != nil {
				queryParams = data.QueryParams
			}
			if data.Body != "" {
				body = data.Body
			}
			if data.PathParams != nil {
				pathParams = data.PathParams
			}
			if data.File != nil {
				file = data.File
			}
			if data.Form != nil {
				form = data.Form
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

			core.InitWorkerPool(jobs, rate, workers)
			core.InitJobs(tasks, jobs, &r, results)

			for a := 0; a < tasks; a++ {
				<-results
				//fmt.Println(<-results)
	}
}
