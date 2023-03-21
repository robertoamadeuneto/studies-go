package handler

import (
	"emailn/internal"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_HandleError_WhenOccursAnInternalServerError_Returns500(t *testing.T) {
	assert := assert.New(t)
	endpoint := func(writer http.ResponseWriter, request *http.Request) (interface{}, int, error) {
		return nil, 0, internal.InternalServerError
	}
	handlerFunc := HandleResponse(endpoint)
	request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()

	handlerFunc.ServeHTTP(response, request)

	assert.Equal(http.StatusInternalServerError, response.Code)
	assert.Contains(response.Body.String(), internal.InternalServerError.Error())
}

func Test_HandleError_WhenOccursValidatorError_Returns422(t *testing.T) {
	assert := assert.New(t)
	errorMessage := "Validator error"
	endpoint := func(writer http.ResponseWriter, request *http.Request) (interface{}, int, error) {
		return nil, 0, errors.New(errorMessage)
	}
	handlerFunc := HandleResponse(endpoint)
	request, _ := http.NewRequest("POST", "/", nil)
	response := httptest.NewRecorder()

	handlerFunc.ServeHTTP(response, request)

	assert.Equal(http.StatusUnprocessableEntity, response.Code)
	assert.Contains(response.Body.String(), errorMessage)
}

func Test_HandleError_WhenAnErrorDoesNotOccur_ReturnsDesiredObjectAndStatus(t *testing.T) {
	assert := assert.New(t)
	object := map[string]string{"id": "id"}
	status := 201
	endpoint := func(writer http.ResponseWriter, request *http.Request) (interface{}, int, error) {
		return object, 201, nil
	}
	handlerFunc := HandleResponse(endpoint)
	request, _ := http.NewRequest("POST", "/", nil)
	response := httptest.NewRecorder()

	handlerFunc.ServeHTTP(response, request)

	assert.Equal(status, response.Code)
	responseBody := map[string]string{}
	json.Unmarshal(response.Body.Bytes(), &responseBody)
	assert.Equal(object, responseBody)
}
