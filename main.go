package main

import (
	"fmt"
	"github.com/kbfu/pegasus/core"
	pegasusHttp "github.com/kbfu/pegasus/components/http"
	"io/ioutil"
	"github.com/go-yaml/yaml"
	"github.com/kbfu/pegasus/utils"
)

func main() {
	var yamlFile map[string]interface{}
	file, _ := ioutil.ReadFile("test.yaml")
	err := yaml.Unmarshal(file, &yamlFile)
	utils.Check(err)
	requests := yamlFile["requests"]

	for _, request := range requests.([]interface{}) {
		var (
			body []byte
			queryParams map[interface{}]interface{}
			pathParams []interface{}
			headers map[interface{}]interface{}
			file map[interface{}]interface{}
			form map[interface{}]interface{}
		)
		url := request.(map[interface{}]interface{})["url"].(string)
		workers := request.(map[interface{}]interface{})["workers"].(int)
		duration := request.(map[interface{}]interface{})["duration"].(int)
		rate := request.(map[interface{}]interface{})["rate"].(int)
		method := request.(map[interface{}]interface{})["method"].(string)
		if request.(map[interface{}]interface{})["headers"] != nil {
			headers = request.(map[interface{}]interface{})["headers"].(map[interface{}]interface{})
		}
		if request.(map[interface{}]interface{})["queryParams"] != nil {
			queryParams = request.(map[interface{}]interface{})["queryParams"].(map[interface{}]interface{})
		}
		if request.(map[interface{}]interface{})["body"] != nil {
			body = []byte(request.(map[interface{}]interface{})["body"].(string))
		}
		if request.(map[interface{}]interface{})["pathParams"] != nil {
			pathParams = request.(map[interface{}]interface{})["pathParams"].([]interface{})
		}
		if request.(map[interface{}]interface{})["file"] != nil {
			file = request.(map[interface{}]interface{})["file"].(map[interface{}]interface{})
		}
		if request.(map[interface{}]interface{})["form"] != nil {
			form = request.(map[interface{}]interface{})["form"].(map[interface{}]interface{})
		}

		tasks := duration * rate
		jobs := make(chan func(), workers)
		results := make(chan map[string]interface{}, tasks)

		r := pegasusHttp.RequestData{
			Client: *core.Client(),
			Url:    url,
			Method: method,
			Body: body,
			QueryParams: queryParams,
			PathParams: pathParams,
			Headers: headers,
			File: file,
			Form: form,
		}

		core.InitWorkerPool(jobs, rate, workers)
		core.InitJobs(tasks, jobs, &r, results)

		for a := 0; a < tasks; a++ {
			fmt.Println(<-results)
		}
	}
}
