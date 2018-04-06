package rest

import (
	_ "bytes"
	_ "encoding/json"
	_ "fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type RequestData struct {
	EndPoint    string
	Headers     map[string]string
	RequestData map[string]string
	Supplier    string
	ReqType     string
	Options     map[string]string
}

var default_request_timeout = 90000

func MakeRequest(method string, req_data RequestData) string {
	options := req_data.Options
	if options["recv_timeout"] == "" {
		options["recv_timeout"] = strconv.Itoa(default_request_timeout)
	}
	if options["ssl"] == "" {
		options["ssl"] = "tlsv1.2"
	}

	client := &http.Client{}
	req, _ := http.NewRequest(method, req_data.EndPoint, nil)
	for key, header := range req_data.Headers {
		req.Header.Add(key, header)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("The HTTP request failed with error %s\n", err)
		return ""
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		return string(data)
	}
}
