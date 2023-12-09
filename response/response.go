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
	if operationType == INSERTION {
		message := make(map[string]string)
		message["success_message"] = fmt.Sprintf("Insertion in [%s] is successful!", functionName)
		c.JSON(statusCode, &APIResponse{
			Error:      []APIError{},
			Data:       message,
			StatusCode: statusCode,
		})
	} else if operationType == DELETION {
		message := make(map[string]string)
		message["success_message"] = fmt.Sprintf("Deletion in [%s] is successful!", functionName)
		c.JSON(statusCode, &APIResponse{
			Error:      []APIError{},
			Data:       message,
			StatusCode: statusCode,
		})
	} else if operationType == UPDATE {
		message := make(map[string]string)
		message["success_message"] = fmt.Sprintf("Update in [%s] is successful!", functionName)
		c.JSON(statusCode, &APIResponse{
			Error:      []APIError{},
			Data:       message,
			StatusCode: statusCode,
		})
	} else {
		message := make(map[string]string)
		message["success_message"] = fmt.Sprintf("Operation in [%s] is Succesfull!", functionName)
		c.JSON(statusCode, &APIResponse{
			Error:      []APIError{},
			Data:       message,
			StatusCode: statusCode,
		})
	}
}
