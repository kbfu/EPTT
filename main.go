package main

import (
	"fmt"
	"github.com/go-yaml/yaml"
	pegasusHttp "github.com/kbfu/pegasus/components/http"
	"github.com/kbfu/pegasus/core"
	"github.com/kbfu/pegasus/utils"
	"io/ioutil"
)

func main() {
	var yamlFile map[string]interface{}
	file, _ := ioutil.ReadFile("test.yaml")
	err := yaml.Unmarshal(file, &yamlFile)
	utils.Check(err)
	requests := yamlFile["requests"]

	for _, request := range requests.([]interface{}) {
		for k, v := range request.(map[interface{}]interface{}) {
			var (
				body        []byte
				queryParams map[interface{}]interface{}
				pathParams  []interface{}
				headers     map[interface{}]interface{}
				file        map[interface{}]interface{}
				form        map[interface{}]interface{}
			)
			url := v.(map[interface{}]interface{})["url"].(string)
			workers := v.(map[interface{}]interface{})["workers"].(int)
			duration := v.(map[interface{}]interface{})["duration"].(int)
			rate := v.(map[interface{}]interface{})["rate"].(int)
			method := v.(map[interface{}]interface{})["method"].(string)
			if v.(map[interface{}]interface{})["headers"] != nil {
				headers = v.(map[interface{}]interface{})["headers"].(map[interface{}]interface{})
			}
			if v.(map[interface{}]interface{})["queryParams"] != nil {
				queryParams = v.(map[interface{}]interface{})["queryParams"].(map[interface{}]interface{})
			}
			if v.(map[interface{}]interface{})["body"] != nil {
				body = []byte(v.(map[interface{}]interface{})["body"].(string))
			}
			if v.(map[interface{}]interface{})["pathParams"] != nil {
				pathParams = v.(map[interface{}]interface{})["pathParams"].([]interface{})
			}
			if v.(map[interface{}]interface{})["file"] != nil {
				file = v.(map[interface{}]interface{})["file"].(map[interface{}]interface{})
			}
			if v.(map[interface{}]interface{})["form"] != nil {
				form = v.(map[interface{}]interface{})["form"].(map[interface{}]interface{})
			}

			tasks := duration * rate
			jobs := make(chan func(), workers)
			results := make(chan map[string]interface{}, tasks)

			r := pegasusHttp.RequestData{
				Client:      *core.Client(),
				Url:         url,
				Method:      method,
				Body:        body,
				QueryParams: queryParams,
				PathParams:  pathParams,
				Headers:     headers,
				File:        file,
				Form:        form,
				Name:        k.(string),
			}

			core.InitWorkerPool(jobs, rate, workers)
			core.InitJobs(tasks, jobs, &r, results)

			for a := 0; a < tasks; a++ {
				fmt.Println(<-results)
			}
		}
	}
}
