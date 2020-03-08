package apiclient

import (
	"errors"

	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type ApiParams struct {
	method     string
	url        string
	params     map[string]string
	queryParam map[string]string
	headers    map[string]string
	req        *http.Request
	rawBody    string
}

func NewApiClient(urlstr string, method string) *ApiParams {
	return &ApiParams{
		strings.ToUpper(method),
		urlstr,
		make(map[string]string),
		make(map[string]string),
		make(map[string]string),
		nil,
		"",
	}
}

func (ap *ApiParams) SetRequestParams(params map[string]string) {
	ap.params = params

}

func (ap *ApiParams) SetRawBody(str string) {
	ap.rawBody = str
}
func (ap *ApiParams) SetQueryParams(qp map[string]string) {
	ap.queryParam = qp
}
func (ap *ApiParams) SetHeaders(h map[string]string) {
	ap.headers = h
}

func (ap *ApiParams) setHeader() {
	for key, val := range ap.headers {
		ap.req.Header.Add(key, val)
	}

}

func (ap *ApiParams) setQueryParams() {
	q := ap.req.URL.Query()
	for key, val := range ap.queryParam {

		q.Add(key, val)

	}
	ap.req.URL.RawQuery = q.Encode()

}

func (ap *ApiParams) setRequestParams() {
	q := ap.req.URL.Query()
	for key, val := range ap.params {

		q.Add(key, val)

	}
	ap.req.URL.RawQuery = q.Encode()
}
func (ap *ApiParams) getParams() url.Values {
	keyVal := url.Values{}
	if ap.method == "POST" {
		if ap.rawBody == "" {
			for key, val := range ap.params {
				keyVal.Add(key, val)
			}

			ap.headers["Content-Type"] = "application/x-www-form-urlencoded"
		} else {
			ap.headers["Content-Type"] = "text/plain"
		}

	}
	return keyVal
}
func (ap *ApiParams) Do() (int, []byte, error) {
	client := &http.Client{}
	//client.Timeout = 0
	paramVal := ap.getParams()

	ap.req, _ = http.NewRequest(ap.method, ap.url, strings.NewReader(paramVal.Encode()))
	if ap.rawBody != "" {
		//var jsonStr = ap.rawBody
		ap.req, _ = http.NewRequest("POST", ap.url, strings.NewReader(ap.rawBody))
	}
	ap.setHeader()
	ap.setQueryParams()
	if ap.method == "GET" {
		ap.setRequestParams()
	}

	resp, err := client.Do(ap.req)
	if err != nil {
		log.Println("API request failed (client.Do)")
		log.Println(err)
		return http.StatusInternalServerError, make([]byte, 0), errors.New("API request failed at source (client.Do)")
	}

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		log.Println("Data read from server failed (iotuil.Readall)")
		log.Println(readErr)
		return http.StatusInternalServerError, make([]byte, 0), errors.New("Data reading from server failed (ioutil.ReadAll)")
	}
	if (resp.StatusCode >= 200) && (resp.StatusCode < 300) {
		return resp.StatusCode, body, nil
	}

	return resp.StatusCode, body, errors.New("Non 2xx response returned")

}
