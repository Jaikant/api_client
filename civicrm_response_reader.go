package apiclient

import (
	"encoding/json"
	"errors"
	"log"
	"reflect"
	"strings"
)

type CrmResponse struct {
	IsError    int                      `json:"is_error"`
	ErrorCode  interface{}              `json:"error_code"`
	Version    int                      `json:"version"`
	Count      int64                    `json:"count"`
	Id         interface{}              `json:"id"`
	ErrorMsg   string                   `json:"error_message"`
	ErrorMsg2  string                   `json:"1"`
	Reason     string                   `json:"reason"`
	Level      string                   `json:"level"`
	Result     int64                    `json:"result"`
	Values     []map[string]interface{} `json:"values"`
	RCode      string                   `json:"code"`
	StrErrCode string
	IDs        []string `json:"ids"`
	FieldName  string   `json:"field"`
	RegId      string   `json:"registration_id"`
	ApiKey     string   `json:"apiKey"`
}

// type Value struct {

// }
func ParseCrmResponse(data []byte) (CrmResponse, error, int) {
	var crmResponse CrmResponse
	var statusCode int
	err := json.Unmarshal(data, &crmResponse)
	if err != nil {
		log.Print("Error parsing Crm data: ")
		log.Println(err)
		log.Println(string(data))
		statusCode = 500
		err = errors.New("Invalid json response returned form crm server.")
	}
	if crmResponse.IsError == 1 {
		err = errors.New(crmResponse.ErrorMsg)
		if crmResponse.ErrorMsg != "" {
			if crmResponse.ErrorMsg == "ERROR: No CMS user associated with given api-key" {
				crmResponse.StrErrCode = "invalid_api_key"
			}
			err = errors.New(crmResponse.ErrorMsg)
		} else if crmResponse.ErrorMsg2 != "" {
			err = errors.New(crmResponse.ErrorMsg2)
		} else {
			err = errors.New("Some unknown error")
		}
		if crmResponse.ErrorCode != nil {
			et := reflect.TypeOf(crmResponse.ErrorCode).Kind()
			if et == reflect.String {
				if crmResponse.ErrorCode.(string) == "duplicate" {
					crmResponse.StrErrCode = "duplicate_entry"
					err = errors.New("duplicate_entry")
				}
			}
			if et == reflect.Float64 {
				if crmResponse.ErrorCode.(float64) == 2001 {
					crmResponse.StrErrCode = "not_existing_value"
					err = errors.New("not_existing_value")
				}
			}

		}

		statusCode = 400
	}
	return crmResponse, err, statusCode

}

func CiviCrmCountryApiUrl(url, country string, urlMap map[string]string) string {
	cc := strings.Trim(country, " ")
	cc = strings.ToLower(cc)
	url = urlMap[cc]
	if url == "" {
		url = urlMap["default"]
	}
	crmUrl := strings.Trim(url, "/")

	crmUrl = crmUrl + "/" + cc + "/sites/all/modules/civicrm/extern/rest.php"
	return crmUrl
}
