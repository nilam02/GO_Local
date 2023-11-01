package http

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

// MakeRequest performs an HTTP request with the provided URL, method, headers, and body, and returns the response body and any error encountered.
func MakeRequest(url, method string, headers map[string]string, body []byte) ([]byte, error) {
	// Create an HTTP client
	client := &http.Client{}

	// Create a new HTTP request with the specified method and URL
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err // Return an error if creating the request fails
	}

	// Set the headers for the request
	for key, value := range headers {
		req.Header.Set(key, value) // Set the provided key-value pairs as request headers
	}

	// If a request body is provided, add it to the request
	if body != nil {
		req.Body = ioutil.NopCloser(bytes.NewReader(body)) // Wrap the body in a NopCloser and set it as the request body
	}

	// Send the HTTP request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err // Return an error if the request fails
	}
	defer resp.Body.Close() // Ensure the response body is closed to prevent resource leaks

	// Read the response body
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err // Return an error if reading the response body fails
	}

	// Return the response body and nil error to indicate success
	return respBody, nil
}