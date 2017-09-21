package http

import (
	"bytes"
	"fmt"
	"github.com/kbfu/pegasus/utils"
	"io"
	"io/ioutil"
	"math"
	"mime/multipart"
	"net/http"
	"os"
	"time"
)

type RequestData struct {
	Url    string
	Method string
	// Content-Type will be auto reset if Form or File has values
	Headers map[interface{}]interface{}
	// Body should be nil if Form or File has values
	Body        []byte
	QueryParams map[interface{}]interface{}
	PathParams  []interface{}
	Client      http.Client
	File        map[interface{}]interface{}
	Form        map[interface{}]interface{}
	Name        string
}

func (r *RequestData) Request(results chan map[string]interface{}) {
	var url string
	var err error
	var req *http.Request

	if r.PathParams != nil {
		url = fmt.Sprintf(r.Url, utils.UnpackString(r.PathParams)...)
	} else {
		url = r.Url
	}

	if r.Form != nil || r.File != nil {
		body, contentType := multipartForm(r.File, r.Form)
		req, err = http.NewRequest(r.Method, url, bytes.NewBuffer(body.Bytes()))
		q := req.URL.Query()
		utils.Check(err)
		req.TransferEncoding = []string{"UTF-8"}
		for k, v := range r.Headers {
			req.Header.Add(k.(string), v.(string))
		}
		for k, v := range r.QueryParams {
			q.Add(k.(string), v.(string))
		}
		req.URL.RawQuery = q.Encode()
		req.Header.Add("Content-Type", contentType)
	} else {
		req, err = http.NewRequest(r.Method, url, bytes.NewBuffer(r.Body))
		q := req.URL.Query()
		utils.Check(err)
		req.TransferEncoding = []string{"UTF-8"}
		for k, v := range r.Headers {
			req.Header.Add(k.(string), v.(string))
		}
		for k, v := range r.QueryParams {
			q.Add(k.(string), v.(string))
		}
		req.URL.RawQuery = q.Encode()
	}
	startTime := time.Now().UnixNano()
	resp, err := r.Client.Do(req)
	elapsedTime := float64(time.Now().UnixNano()-startTime) / math.Pow10(9)
	defer resp.Body.Close()
	data := make(map[string]interface{})
	body, err := ioutil.ReadAll(resp.Body)
	utils.Check(err)
	data["statusCode"] = resp.StatusCode
	data["body"] = body
	data["elapsed"] = elapsedTime
	data["startTime"] = startTime
	data["error"] = err
	data["name"] = r.Name
	results <- data
}

func multipartForm(file map[interface{}]interface{}, form map[interface{}]interface{}) (body bytes.Buffer, contentType string) {
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	// add form file
	for k, v := range file {
		f, err := os.Open(v.(string))
		utils.Check(err)
		fw, err := w.CreateFormFile(k.(string), v.(string))
		utils.Check(err)
		_, err = io.Copy(fw, f)
		f.Close()
		utils.Check(err)
	}
	// add form data
	for k, v := range form {
		fw, err := w.CreateFormField(k.(string))
		utils.Check(err)
		_, err = fw.Write([]byte(v.(string)))
		utils.Check(err)
	}
	w.Close()

	return b, w.FormDataContentType()
}
