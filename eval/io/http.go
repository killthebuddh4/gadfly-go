package io

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/killthebuddh4/gadflai/types"
)

func Http(trajectory *types.Trajectory, eval types.Exec) (types.Value, error) {
	types.ExpandTraj(trajectory)

	urlV, err := eval(trajectory.Children[0])

	if err != nil {
		return nil, err
	}

	url, ok := urlV.(string)

	if !ok {
		return nil, errors.New("url is not a string")
	}

	requestMethodV, err := eval(trajectory.Children[1])

	if err != nil {
		return nil, err
	}

	requestMethod, ok := requestMethodV.(string)

	if !ok {
		return nil, errors.New("method is not a string")
	}

	requestHeadersV, err := eval(trajectory.Children[2])

	if err != nil {
		return nil, err
	}

	requestHeaders, ok := requestHeadersV.(map[string]types.Value)

	if !ok {
		return nil, errors.New("headers is not a map")
	}

	requestBodyV, err := eval(trajectory.Children[3])

	if err != nil {
		return nil, err
	}

	requestBody, ok := requestBodyV.(map[string]types.Value)

	if !ok {
		return nil, errors.New("body is not a map")
	}

	jsonRequestBody, err := json.Marshal(requestBody)

	if err != nil {
		return nil, err
	}

	var req *http.Request
	switch strings.ToUpper(requestMethod) {
	case "GET":
		req, err = http.NewRequest(requestMethod, url, nil)
	case "POST", "PUT", "PATCH", "DELETE":
		req, err = http.NewRequest(requestMethod, url, bytes.NewBuffer(jsonRequestBody))
	default:
		return nil, fmt.Errorf("unsupported method: %s", requestMethod)
	}

	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	for field, header := range requestHeaders {
		value, ok := header.(string)

		if !ok {
			return nil, errors.New("request header value is not a string")
		}

		req.Header.Add(field, value)
	}

	client := &http.Client{}

	response, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	var raw interface{}

	err = json.Unmarshal(responseBody, &raw)

	if err != nil {
		return nil, err
	}

	value := make(map[string]types.Value)

	for k, v := range raw.(map[string]interface{}) {
		value[k] = v
	}

	return value, nil
}
