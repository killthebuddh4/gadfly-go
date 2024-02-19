package eval

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

var Http types.Exec = func(scope *types.Trajectory, arguments ...types.Value) (types.Value, error) {
	url, ok := arguments[0].(string)

	if !ok {
		return nil, errors.New("url is not a string")
	}

	requestMethod, ok := arguments[1].(string)

	if !ok {
		return nil, errors.New("method is not a string")
	}

	requestHeaders, ok := arguments[2].(map[string]types.Value)

	if !ok {
		return nil, errors.New("headers is not a map")
	}

	requestBody, ok := arguments[3].(map[string]types.Value)

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
