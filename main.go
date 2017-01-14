package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	schema, err := Load("http://petstore.swagger.io/v2/swagger.json")
	if err != nil {
		return
	}
	fmt.Printf("%#v\r\n", schema)
}

func Load(url string) (interface{}, error) {
	var schema interface{}
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
	return schema, nil
}
