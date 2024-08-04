package response

import (
	"fmt"
	"reflect"
)

func GenerateCodeResponse(httpCode int, featureCode string, layerCode string) string {
	codeResponse := fmt.Sprintf("%d", httpCode) + "-" + featureCode + "-" + layerCode
	return codeResponse
}

type Response struct {
	Code    string `json:"code,omitempty"`
	Status  string `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
}

/*
func to create map response json, and http response code
err: error message
featureCode: code features from constanta. according to the feature that are being made
*/
func WebResponseError(err error, featureCode string) (response map[string]any, httpCode int) {

	errCode, layerCode, errMessage := CheckHandlerErrorCode(err)
	code := GenerateCodeResponse(errCode, featureCode, layerCode)

	if reflect.TypeOf(err).String() == "validation.ValidationError" {
		valErr := err.(ValidationError)
		response = map[string]any{
			"meta": Response{
				Code:    code,
				Message: "failed",
			},
			"messages": valErr.Errors,
		}
	} else {
		response = map[string]any{
			"meta": Response{
				Code:    code,
				Message: "failed",
			},
			"messages": []string{errMessage.Error()},
		}
	}

	return response, errCode
}

/*
func to create map response json, and http response code
messages: message response
featureCode: code features from constanta. according to the feature that are being made
data: if data response is available, put here. otherwise put nil.
*/
func WebResponseSuccess(messages string, featureCode string, data any) (response map[string]any, httpCode int) {
	httpCode = CheckHandlerSuccessCode(messages)
	code := GenerateCodeResponse(httpCode, featureCode, RESPONSE_SUCCESS_CODE)

	response = map[string]any{
		"meta": Response{
			Code:   code,
			Status: "success",
		},
		"messages": []string{messages},
		"data":     data,
	}

	return response, httpCode
}
