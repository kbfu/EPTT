package main

import (
	"fmt"
	"github.com/kbfu/pegasus/components/http"
	"github.com/valyala/fasthttp"
	"math"
	"time"
)

var responseChannel = make(chan map[string]interface{}, 200)

func main() {
	count := 10000
	goroutineCount := 20
	r := &http.RequestData{
		Loop:    count / goroutineCount,
		ResponseChannel:       responseChannel,
		Client:  &fasthttp.Client{},
		Url:     "http://localhost:60005/hello",
	}

	startTime := time.Now().UnixNano()
	for i := 0; i < goroutineCount; i++ {
		go http.DoPost(r)
	}
	for i := 0; i < count; i++ {
		data := <-responseChannel
		fmt.Println(data["statusCode"])
	}
	endTime := time.Now().UnixNano()
	fmt.Println(float64(endTime-startTime) / math.Pow10(9))
}
