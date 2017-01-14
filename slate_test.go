package main

import (
	"testing"
)

func TestSwaggerSchemaInitialization(t *testing.T) {
	schema, err := load("http://petstore.swagger.io/v2/swagger.json")
	if err != nil {
		t.Error(err)
	}
	if len(schema.Paths) != 14 {
		t.Error()
	}
	if len(schema.Paths["/pet"]) != 2 {
		t.Error()
	}
}
