package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/rightlag/go-swagger-object-model"
)

func main() {
	schema, err := NewSchema("http://petstore.swagger.io/v2/swagger.json")
	if err != nil {
		return
	}
	// fmt.Printf("%#v\r\n", schema.Paths["/pet"].Put.Responses["400"])
	fmt.Printf("%#v\r\n", schema)
}

func NewSchema(url string) (*models.Schema, error) {
	schema, err := load(url)
	if err != nil {
		return nil, err
	}
	return schema, nil
}

func load(url string) (*models.Schema, error) {
	var schema models.Schema
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	jsonString, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(jsonString, &schema)
	if err != nil {
		return nil, err
	}
	return &schema, nil
}
