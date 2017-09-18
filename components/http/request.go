package http

import (
	"fmt"
	"github.com/kbfu/pegasus/utils"
	"github.com/valyala/fasthttp"
)

type RequestData struct {
	Loop            int
	ResponseChannel chan map[string]interface{}
	Client          *fasthttp.Client
	Url             string
	Headers         []map[string]string
	PathParams      []string
	QueryParams     map[string]string
	Body            []byte
}

type Requester interface {
	Request() (*RequestData, fasthttp.Request)
}

func (r *RequestData) Request() (*RequestData, fasthttp.Request) {
	request := fasthttp.AcquireRequest()
	var headers fasthttp.RequestHeader
	var url string
	if r.Headers != nil {
		for _, v := range r.Headers {
			for k := range v {
				headers.Add(k, v[k])
			}
		}
		request.Header = headers
	}
	if r.PathParams != nil {
		url = fmt.Sprintf(r.Url, utils.UnpackString(r.PathParams)...)
	}
	if r.QueryParams != nil {
		params := ""
		for k, v := range r.QueryParams {
			params = params + fmt.Sprintf("%s=%s&", k, v)
		}
		url = url + "?" + params[:len(params)-1]
	}
	if url == "" {
		url = r.Url
	}
	request.SetRequestURI(url)
	if r.Body != nil {
		request.SetBody(r.Body)
	}
	return r, *request
}
