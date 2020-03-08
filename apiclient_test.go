package apiclient

import (
	//"encoding/json"
	"fmt"
	"testing"
)

func TestHttpClient(t *testing.T) {
	url := "https://httpbin.org/get"
	qp := map[string]string{
		"name":  "sanjay",
		"name1": "dhiraj",
	}
	h := map[string]string{
		"authorization": "just a test authorization header",
		"some-other":    "some other value header",
		"user-agent":    "sand client",
	}
	p := map[string]string{
		"first_name": "sand12345",
		"name2":      "dhiraj",
		"phone":      "123123123"}

	fmt.Println("Test started with following parameters:-")
	fmt.Println("Query strings set:")
	fmt.Println(qp)
	fmt.Println("Request  parametes set:")
	fmt.Println(p)
	fmt.Println("Headers set:")
	fmt.Println(h)

	fmt.Println("Making request to url: ")
	fmt.Println(url)

	c := NewApiClient(url, "get")

	c.SetRequestParams(p)

	c.SetQueryParams(qp)
	c.SetHeaders(h)
	rc, body, err := c.Do()
	fmt.Print("Response code: ")
	fmt.Println(rc)
	fmt.Println("Response body: ")
	fmt.Println(string(body))
	fmt.Print("Error generated: ")
	fmt.Println(err)
	fmt.Println("Testing package generated messages are as below")

}

func TestRaw(t *testing.T) {
	url := "https://www.vvkicrm.org/sites/vvkicrm.org/modules/civicrm/extern/rest.php"
	qp := map[string]string{
		"json":    "1",
		"entity":  "course",
		"action":  "getContactDetailsByAolid",
		"api_key": "VvkiDN@2015PS",
		"key":     "2c8f9b72e546c30d6bb1030a9d47626e",
	}
	h := map[string]string{
		"Content-Type": "text/plain",
	}
	p := map[string]string{"aol_Id": "10000011"}
	//b,_ := json.Marshal(p)
	fmt.Println("Test started with following parameters:-")
	fmt.Println("Query strings set:")
	fmt.Println(qp)
	fmt.Println("Request  parametes set:")
	fmt.Println(p)
	fmt.Println("Headers set:")
	fmt.Println(h)

	fmt.Println("Making request to url: ")
	fmt.Println(url)

	c := NewApiClient(url, "post")

	//c.SetRequestParams(p)
	c.SetRawBody(`{"aol_Id":10000011}`)
	c.SetQueryParams(qp)
	c.SetHeaders(h)
	rc, body, err := c.Do()
	fmt.Print("Response code: ")
	fmt.Println(rc)
	fmt.Println("Response body: ")
	fmt.Println(string(body))
	fmt.Print("Error generated: ")
	fmt.Println(err)
	fmt.Println("Testing package generated messages are as below")

}
