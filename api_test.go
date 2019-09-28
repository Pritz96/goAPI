package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

var timeResponse = "The time is " + time.Now().Format(time.RFC1123)

func TestTimeHandler(t *testing.T) {
	tests := []struct {
		method           string
		expectedStatus   int
		expectedResponse string
	}{
		{
			method:           "GET",
			expectedStatus:   http.StatusOK,
			expectedResponse: timeResponse,
		},
		{
			method:           "POST",
			expectedStatus:   http.StatusMethodNotAllowed,
			expectedResponse: "HTTP request method not allowed",
		},
		{
			method:           "DELETE",
			expectedStatus:   http.StatusMethodNotAllowed,
			expectedResponse: "HTTP request method not allowed",
		},
	}
	for _, test := range tests {
		req, err := http.NewRequest(test.method, "/test", nil)
		if err != nil {
			t.Fatal(err)
		}
		resp := httptest.NewRecorder()
		handler := http.HandlerFunc(timeHandler)
		handler.ServeHTTP(resp, req)

		status := resp.Code
		if status != test.expectedStatus {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, test.expectedStatus)
		}

		expected := test.expectedResponse
		if resp.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				resp.Body.String(), expected)
		}
	}
}

func TestHelloHandler(t *testing.T) {
	tests := []struct {
		method           string
		expectedStatus   int
		expectedResponse string
	}{
		{
			method:           "GET",
			expectedStatus:   http.StatusOK,
			expectedResponse: "Hello",
		},
		{
			method:           "POST",
			expectedStatus:   http.StatusMethodNotAllowed,
			expectedResponse: "HTTP request method not allowed",
		},
		{
			method:           "DELETE",
			expectedStatus:   http.StatusMethodNotAllowed,
			expectedResponse: "HTTP request method not allowed",
		},
	}
	for _, test := range tests {
		req, err := http.NewRequest(test.method, "/test", nil)
		if err != nil {
			t.Fatal(err)
		}
		resp := httptest.NewRecorder()
		handler := http.HandlerFunc(helloHandler)
		handler.ServeHTTP(resp, req)

		status := resp.Code
		if status != test.expectedStatus {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, test.expectedStatus)
		}

		expected := test.expectedResponse
		if resp.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				resp.Body.String(), expected)
		}
	}
}
