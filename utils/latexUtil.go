package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/Lukaesebrot/asterisk/static"
	"github.com/valyala/fasthttp"
)

type rtexRequest struct {
	Code    string `json:"code"`
	Format  string `json:"format"`
	Quality int    `json:"quality"`
	Density int    `json:"density"`
}

type rtexResponse struct {
	Status      string `json:"status"`
	Log         string `json:"log"`
	Filename    string `json:"filename,omitempty"`
	Description string `json:"description,omitempty"`
}

// RenderLaTeX renders a LaTeX expression and returns the image URL
func RenderLaTeX(expression string) (string, error) {
	// Format the expression
	expression = strings.Replace(static.LaTeXTemplate, "#CONTENT#", expression, 1)

	// Acquire a request object
	request := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(request)

	// Acquire a response object
	response := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(response)

	// Define the request values
	request.SetRequestURI(fmt.Sprintf("%s", static.RTeXURL))
	request.Header.SetMethod("POST")
	request.Header.SetContentType("application/json")
	body, _ := json.Marshal(&rtexRequest{
		Code:    expression,
		Format:  "png",
		Quality: 85,
		Density: 200,
	})
	request.SetBody(body)

	// Perform the request
	err := fasthttp.Do(request, response)
	if err != nil {
		return "", err
	}

	// Unmarshal the response body and return the formatted URL
	jsonResponse := new(rtexResponse)
	err = json.Unmarshal(response.Body(), jsonResponse)
	if err != nil {
		return "", err
	}
	if jsonResponse.Status == "error" {
		return "", errors.New(jsonResponse.Description)
	}
	return fmt.Sprintf("%s/%s", static.RTeXURL, jsonResponse.Filename), nil
}
