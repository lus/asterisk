package mathjs

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/Lukaesebrot/asterisk/static"
	"github.com/valyala/fasthttp"
)

type mjsRequest struct {
	Expression string `json:"expr"`
}

type mjsResponse struct {
	Result string `json:"result,omitempty"`
	Error  string `json:"error,omitempty"`
}

// Evaluate evaluates the given expression and returns the result
func Evaluate(expression string) (string, error) {
	// Acquire a request object
	request := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(request)

	// Acquire a response object
	response := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(response)

	// Define the request values
	request.SetRequestURI(fmt.Sprintf("%s", static.MathJSURL))
	request.Header.SetMethod("POST")
	request.Header.SetContentType("application/json")
	body, _ := json.Marshal(&mjsRequest{
		Expression: expression,
	})
	request.SetBody(body)

	// Perform the request
	err := fasthttp.Do(request, response)
	if err != nil {
		return "", err
	}

	// Unmarshal the response body and return the result
	jsonResponse := new(mjsResponse)
	err = json.Unmarshal(response.Body(), jsonResponse)
	if err != nil {
		return "", err
	}
	if jsonResponse.Error != "" {
		return "", errors.New(jsonResponse.Error)
	}
	return jsonResponse.Result, nil
}
