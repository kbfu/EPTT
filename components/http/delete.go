package http

import (
	"github.com/valyala/fasthttp"
	"log"
	"math"
	"time"
)

func DoDelete(r Requester) {
	var response fasthttp.Response
	requestData, request := r.Request()
	request.Header.SetMethod("DELETE")

	for i := 0; i < requestData.Loop; i++ {
		data := make(map[string]interface{})
		startTime := time.Now().UnixNano()
		err := requestData.Client.Do(&request, &response)
		if err != nil {
			log.Fatal(err)
		}
		elapsedTime := float64(time.Now().UnixNano()-startTime) / math.Pow10(9)
		data["statusCode"], data["body"], data["elapsed"], data["startTime"] = response.StatusCode(),
			response.Body(), elapsedTime, float64(startTime) / math.Pow10(9)
		requestData.ResponseChannel <- data
	}
	fasthttp.ReleaseRequest(&request)
	fasthttp.ReleaseResponse(&response)
}
