package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResponseSucces(t *testing.T) {
	data := "ini data"
	code := 200
	isError := false
	result := &Response{
		Code:    code,
		Status:  getStatus(code),
		IsError: isError,
		Data:    data,
	}

	assert.Equal(t, result, New(data, code, isError), "expected response")
}
func TestResponseError(t *testing.T) {
	data := "ini error"
	code := 400
	isError := true
	result := &Response{
		Code:        code,
		Status:      getStatus(code),
		IsError:     isError,
		Description: data,
	}

	assert.Equal(t, result, New(data, code, isError), "expected response")
}

// Sub test
func TestGetStatus(t *testing.T) {
	t.Run("status OK", func(t *testing.T) {
		result := getStatus(200)
		assert.Equal(t, "OK", result, "description must be OK")
	})

	t.Run("status Created", func(t *testing.T) {
		result := getStatus(201)
		assert.Equal(t, "Created", result, "description must be Created")
	})

	t.Run("status Bad Request", func(t *testing.T) {
		result := getStatus(400)
		assert.Equal(t, "Bad Request", result, "description must be Bad Request")
	})

	t.Run("status Unauthorized", func(t *testing.T) {
		result := getStatus(401)
		assert.Equal(t, "Unauthorized", result, "description must be Unauthorized")
	})

	t.Run("status Not Found", func(t *testing.T) {
		result := getStatus(404)
		assert.Equal(t, "Not Found", result, "description must be Not Found")
	})

	t.Run("status Bad Gateway", func(t *testing.T) {
		result := getStatus(501)
		assert.Equal(t, "Bad Gateway", result, "description must be Bad Gateway")
	})

	t.Run("status Internal Server Error", func(t *testing.T) {
		result := getStatus(500)
		assert.Equal(t, "Internal Server Error", result, "description must be Internal Server Error")
	})

	t.Run("status Not Modified", func(t *testing.T) {
		result := getStatus(304)
		assert.Equal(t, "Not Modified", result, "description must be Not Modified")
	})

	t.Run("status Code Undefined", func(t *testing.T) {
		result := getStatus(666)
		assert.Equal(t, "Undefined", result, "description must be Undefined")
	})
}
