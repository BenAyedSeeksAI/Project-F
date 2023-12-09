package response

import (
	"reflect"

	"github.com/gin-gonic/gin"
)

type APIError struct {
	ErrorMessage string `json:"error_message"`
	Description  string `json:"description"`
}
type APIResponse struct {
	Error      []APIError `json:"error"`
	Data       any        `json:"data"`
	StatusCode int        `json:"statusCode"`
}

func Success(c *gin.Context, statusCode int, data any) {

	var nonNilData any
	value := reflect.ValueOf(data)
	if data == nil || (value.Kind() == reflect.Pointer && value.IsNil()) {
		nonNilData = map[string]any{}
	} else {
		nonNilData = data
	}

	c.JSON(statusCode, &APIResponse{
		Error:      []APIError{},
		Data:       nonNilData,
		StatusCode: statusCode,
	})

}
