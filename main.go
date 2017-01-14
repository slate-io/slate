package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/rightlag/go-swagger-object-model"
)

func main() {}

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
