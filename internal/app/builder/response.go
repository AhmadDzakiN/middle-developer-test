package builder

import (
	"reflect"
)

type SuccessResp struct {
	StatusCode int         `json:"status_code"`
	Status     string      `json:"status"`
	Data       interface{} `json:"data,omitempty"`
}

type ErrorResp struct {
	StatusCode int    `json:"status_code"`
	Status     string `json:"status"`
	Message    string `json:"message,omitempty"`
}

// BuildResponse build api response
func BuildResponse(respType string, data interface{}, statusCode int) interface{} {
	if respType == "success" {
		resp := SuccessResp{}
		resp.Status = "success"
		resp.StatusCode = statusCode

		if reflect.ValueOf(data).Kind() == reflect.Slice {
			resp.Data = data
		} else {
			temp := []interface{}{}
			if reflect.ValueOf(data).IsValid() {
				temp = []interface{}{
					data,
				}

			}

			resp.Data = temp
		}

		return resp
	} else {
		resp := ErrorResp{}
		resp.Status = "error"
		resp.StatusCode = statusCode
		resp.Message = data.(string)

		return resp
	}
}
