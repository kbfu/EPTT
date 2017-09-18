package main

import (
	"fmt"
	"github.com/go-yaml/yaml"
	"github.com/kbfu/pegasus/components/http"
	"github.com/kbfu/pegasus/utils"
	"github.com/valyala/fasthttp"
	"io/ioutil"
	"math"
	"strings"
	"time"
)

var responseChannel = make(chan map[string]interface{}, 200)
var yamlFile map[string]interface{}

func main() {
	file, _ := ioutil.ReadFile("test.yaml")
	err := yaml.Unmarshal(file, &yamlFile)
	utils.Check(err)
	requests := yamlFile["requests"]
	for _, request := range requests.([]interface{}) {
		var body []byte
		var queryParams map[interface{}]interface{}
		var pathParams []interface{}
		url := request.(map[interface{}]interface{})["url"].(string)
		loop := request.(map[interface{}]interface{})["loop"].(int)
		method := request.(map[interface{}]interface{})["method"].(string)
		goroutineCount := request.(map[interface{}]interface{})["goroutine"].(int)
		headers := request.(map[interface{}]interface{})["headers"].(map[interface{}]interface{})
		if request.(map[interface{}]interface{})["queryParams"] != nil {
			queryParams = request.(map[interface{}]interface{})["queryParams"].(map[interface{}]interface{})
		}
		if request.(map[interface{}]interface{})["body"] != nil {
			body = []byte(request.(map[interface{}]interface{})["body"].(string))
		}
		if request.(map[interface{}]interface{})["pathParams"] != nil {
			pathParams = request.(map[interface{}]interface{})["pathParams"].([]interface{})
		}

		switch strings.ToLower(method) {
		case "post":
			r := &http.RequestData{
				Loop:            loop,
				ResponseChannel: responseChannel,
				Client:          &fasthttp.Client{},
				Url:             url,
				Headers:         headers,
				QueryParams:     queryParams,
				Body:            body,
				PathParams:      pathParams,
			}
			startTime := time.Now().UnixNano()
			for i := 0; i < goroutineCount; i++ {
				go http.DoPost(r)
			}
			for i := 0; i < loop*goroutineCount; i++ {
				data := <-responseChannel
				fmt.Println(data["statusCode"])
			}
			endTime := time.Now().UnixNano()
			fmt.Println(float64(endTime-startTime) / math.Pow10(9))
		case "get":
			r := &http.RequestData{
				Loop:            loop,
				ResponseChannel: responseChannel,
				Client:          &fasthttp.Client{},
				Url:             url,
				Headers:         headers,
				QueryParams:     queryParams,
				Body:            body,
				PathParams:      pathParams,
			}
			startTime := time.Now().UnixNano()
			for i := 0; i < goroutineCount; i++ {
				go http.DoGet(r)
			}
			for i := 0; i < loop*goroutineCount; i++ {
				data := <-responseChannel
				fmt.Println(data["statusCode"])
			}
			endTime := time.Now().UnixNano()
			fmt.Println(float64(endTime-startTime) / math.Pow10(9))
		case "put":
			r := &http.RequestData{
				Loop:            loop,
				ResponseChannel: responseChannel,
				Client:          &fasthttp.Client{},
				Url:             url,
				Headers:         headers,
				QueryParams:     queryParams,
				Body:            body,
				PathParams:      pathParams,
			}
			startTime := time.Now().UnixNano()
			for i := 0; i < goroutineCount; i++ {
				go http.DoPut(r)
			}
			for i := 0; i < loop*goroutineCount; i++ {
				data := <-responseChannel
				fmt.Println(data["statusCode"])
			}
			endTime := time.Now().UnixNano()
			fmt.Println(float64(endTime-startTime) / math.Pow10(9))
		case "delete":
			r := &http.RequestData{
				Loop:            loop,
				ResponseChannel: responseChannel,
				Client:          &fasthttp.Client{},
				Url:             url,
				Headers:         headers,
				QueryParams:     queryParams,
				Body:            body,
				PathParams:      pathParams,
			}
			startTime := time.Now().UnixNano()
			for i := 0; i < goroutineCount; i++ {
				go http.DoDelete(r)
			}
			for i := 0; i < loop*goroutineCount; i++ {
				data := <-responseChannel
				fmt.Println(data["statusCode"])
			}
			endTime := time.Now().UnixNano()
			fmt.Println(float64(endTime-startTime) / math.Pow10(9))
		}
	}
}
