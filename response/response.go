package response

import (
	"fmt"
	"reflect"

	"github.com/gin-gonic/gin"
)

const (
	INSERTION = 1
	DELETION  = 2
	UPDATE    = 3
	OPERATION = 4
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
type APIOperationResponse struct {
	Error      []APIError `json:"error"`
	Message    string     `json:"message"`
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
func SuccessOperation(c *gin.Context, statusCode int, functionName string, operationType int) {
	operationResponse := &APIOperationResponse{
		Error:      []APIError{},
		StatusCode: statusCode,
	}
	if operationType == INSERTION {
		operationResponse.Message = fmt.Sprintf("Insertion in [%s] is successful!", functionName)
		c.JSON(statusCode, operationResponse)
	} else if operationType == DELETION {
		operationResponse.Message = fmt.Sprintf("Deletion in [%s] is successful!", functionName)
		c.JSON(statusCode, operationResponse)
	} else if operationType == UPDATE {
		operationResponse.Message = fmt.Sprintf("Update in [%s] is successful!", functionName)
		c.JSON(statusCode, operationResponse)
	} else {
		operationResponse.Message = fmt.Sprintf("Operation in [%s] is Succesfull!", functionName)
		c.JSON(statusCode, operationResponse)
	}
}
