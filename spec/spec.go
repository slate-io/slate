package spec

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/rightlag/go-swagger-object-model"
)

func LoadSchema(u *url.URL) (*models.Schema, error) {
	var schema models.Schema
	response, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	specification, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(specification, &schema)
	if err != nil {
		return nil, err
	}
	return &schema, nil
}
