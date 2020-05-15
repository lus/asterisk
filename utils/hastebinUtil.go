package utils

import (
	"encoding/json"
	"fmt"

	"github.com/Lukaesebrot/asterisk/static"

	"github.com/valyala/fasthttp"
)

type hastebinResponse struct {
	Key string `json:"key"`
}

// CreateHaste creates a haste and returns the URL
func CreateHaste(content string) (string, error) {
	// Acquire a request object
	request := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(request)

	// Acquire a response object
	response := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(response)

	// Define the request values
	request.SetRequestURI(fmt.Sprintf("%sdocuments", static.HastebinURL))
	request.Header.SetMethod("POST")
	request.Header.SetContentType("text/plain")
	request.SetBodyString(content)

	// Perform the request
	err := fasthttp.Do(request, response)
	if err != nil {
		return "", err
	}

	// Unmarshal the response body and return the formatted URL
	jsonResponse := new(hastebinResponse)
	err = json.Unmarshal(response.Body(), jsonResponse)
	return fmt.Sprintf("%s%s", static.HastebinURL, jsonResponse.Key), err
}
