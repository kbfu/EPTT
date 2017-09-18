package http

import (
	"github.com/valyala/fasthttp"
	"math"
	"time"
	"github.com/kbfu/pegasus/utils"
)

func DoPost(r Requester) {
	var response fasthttp.Response
	requestData, request := r.Request()
	request.Header.SetMethod("POST")

	for i := 0; i < requestData.Loop; i++ {
		data := make(map[string]interface{})
		startTime := time.Now().UnixNano()
		err := requestData.Client.Do(&request, &response)
		utils.Check(err)
		elapsedTime := float64(time.Now().UnixNano()-startTime) / math.Pow10(9)
		data["statusCode"], data["body"], data["elapsed"], data["startTime"] = response.StatusCode(),
			response.Body(), elapsedTime, float64(startTime) / math.Pow10(9)
		requestData.ResponseChannel <- data
	}
	fasthttp.ReleaseRequest(&request)
	fasthttp.ReleaseResponse(&response)
}
