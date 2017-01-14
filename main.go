package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Schema struct {
	Swagger             string      `json:"swagger"`
	Info                interface{} `json:"info"`
	Host                string      `json:"host,omitempty"`
	BasePath            string      `json:"basePath,omitempty"`
	Schemes             []string    `json:"schemes,omitempty"`
	Consumes            []string    `json:"consumes,omitempty"`
	Produces            []string    `json:"produces,omitempty"`
	Paths               interface{} `json:"paths"`
	Definitions         interface{} `json:"definitions,omitempty"`
	Parameters          interface{} `json:"parameters,omitempty"`
	Responses           interface{} `json:"responses,omitempty"`
	SecurityDefinitions interface{} `json:"securityDefinitions,omitempty"`
	Security            interface{} `json:"security,omitempty"`
	Tags                interface{} `json:"tags,omitempty"`
	ExternalDocs        interface{} `json:"externalDocs,omitempty"`
}

func main() {
	schema, err := NewSchema("http://petstore.swagger.io/v2/swagger.json")
	if err != nil {
		return
	}
	fmt.Printf("%#v\r\n", schema)
}

func NewSchema(url string) (*Schema, error) {
	schema, err := load(url)
	if err != nil {
		return nil, err
	}
	return schema, nil
}

func load(url string) (*Schema, error) {
	var schema Schema
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
