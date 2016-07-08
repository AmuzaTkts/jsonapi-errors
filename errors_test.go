package errors

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetErrorClass(t *testing.T) {
	var cases = []struct {
		code     int
		expected int
	}{
		{301, 0},
		{302, 0},
		{600, 0},
		{900, 0},

		{401, 400},
		{402, 400},
		{403, 400},

		{500, 500},
		{501, 500},
		{502, 500},
	}

	for _, tc := range cases {
		assert.Equal(t, tc.expected, GetErrorClass(tc.code))
	}
}

func TestNewBag(t *testing.T) {
	bag := NewBag()

	assert.IsType(t, &Bag{}, bag)
	assert.Equal(t, 0, len(bag.Errors))
}

func TestNewBagWithError(t *testing.T) {
	bag := NewBagWithError(500, "Oops")

	assert.IsType(t, &Bag{}, bag)
	assert.Equal(t, 1, len(bag.Errors))
	assert.Equal(t, 500, bag.Status)
}

func TestBagStatusWithMultipleErrorsWithSameStatus(t *testing.T) {
	bag := NewBag()

	bag.AddError(502, "Server Error 1")
	bag.AddError(502, "Server Error 2")

	jsonString := `{"errors":[{"detail":"Server Error 1","status":"502"},{"detail":"Server Error 2","status":"502"}],"status":"502"}`
	jsonBytes, _ := json.Marshal(bag)

	assert.Equal(t, 502, bag.Status)
	assert.Equal(t, jsonString, string(jsonBytes))
}

func TestBagStatusWithMultipleErrorsWithSameClass(t *testing.T) {
	bag := NewBag()

	bag.AddError(501, "Server Error 1")
	bag.AddError(502, "Server Error 2")

	jsonString := `{"errors":[{"detail":"Server Error 1","status":"501"},{"detail":"Server Error 2","status":"502"}],"status":"500"}`
	jsonBytes, _ := json.Marshal(bag)

	assert.Equal(t, 500, bag.Status)
	assert.Equal(t, jsonString, string(jsonBytes))
}

func TestBagStatusWithMultipleErrorsWithDifferentClass(t *testing.T) {
	bag := NewBag()

	bag.AddError(401, "Client Error 1")
	bag.AddError(502, "Server Error 1")

	jsonString := `{"errors":[{"detail":"Client Error 1","status":"401"},{"detail":"Server Error 1","status":"502"}],"status":"400"}`
	jsonBytes, _ := json.Marshal(bag)

	assert.Equal(t, 400, bag.Status)
	assert.Equal(t, jsonString, string(jsonBytes))
}

func TestError(t *testing.T) {
	err := NewError(400, "Test error")

	assert.Error(t, err)
	assert.Equal(t, "400: Test error", err.Error())
}

func TestErrorSetAboutLink(t *testing.T) {
	link := "http://google.com"
	err := NewError(400, "Test error")
	err.SetAboutLink(link)

	jsonString := `{"detail":"Test error","links":{"about":"http://google.com"},"status":"400"}`
	jsonBytes, _ := json.Marshal(err)

	assert.Equal(t, link, err.Links.About)
	assert.Equal(t, jsonString, string(jsonBytes))
}

func TestSetPointer(t *testing.T) {
	pointer := "/data/attributes/password"
	err := NewError(400, "Test error")
	err.SetPointer(pointer)

	jsonString := `{"detail":"Test error","source":{"pointer":"/data/attributes/password"},"status":"400"}`
	jsonBytes, _ := json.Marshal(err)

	assert.Equal(t, pointer, err.Source.Pointer)
	assert.Equal(t, jsonString, string(jsonBytes))
}

func TestSetParameter(t *testing.T) {
	parameter := "order"
	err := NewError(400, "Test error")
	err.SetParameter(parameter)

	jsonString := `{"detail":"Test error","source":{"parameter":"order"},"status":"400"}`
	jsonBytes, _ := json.Marshal(err)

	assert.Equal(t, parameter, err.Source.Parameter)
	assert.Equal(t, jsonString, string(jsonBytes))
}
