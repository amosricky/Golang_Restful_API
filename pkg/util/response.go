package util

import (
	"Golang_Restful_API/pkg/setting"
	"github.com/sirupsen/logrus"
	"net/http"
)

type ResponseBody struct {
	Code    int         `json:"error_code"`
	Message string      `json:"error_message"`
	Result  interface{} `json:"result"`
}

func NewResponseBody(extendError ExtendError) ResponseBody {
	return ResponseBody{Code: extendError.GetCode(), Message: extendError.GetMessage()}
}

func (r *ResponseBody) SetExtendError(extendError ExtendError) {
	r.Code = extendError.GetCode()
	r.Message = extendError.GetMessage()

	if setting.ServerSetting.RunMode == "debug" {
		logrus.Debugf("Code:%d Message:%s", r.Code, r.Message)
	}
}

func (r *ResponseBody) StatusCode() int {

	statusCode := 500
	if http.StatusContinue <= r.Code && r.Code <= http.StatusNetworkAuthenticationRequired {
		statusCode = r.Code
	}

	return statusCode
}