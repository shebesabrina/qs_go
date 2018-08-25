package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
  "bytes"
)

func TestPutSuccess(t *testing.T) {
	testFood := Food{
		ID:       "1",
		Name:     "Banana",
		Calories: "100",
	}
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	request, err := createRequest("PUT", "/api/v1/foods/{id}", testFood)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	response := httptest.NewRecorder()
	handler := http.HandlerFunc(UpdateFood)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(response, request)

	// Check the status code is what we expect.
	if status := response.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected, err := createExpected(testFood)

	if err != nil {
		t.Fatal(err)
	}

	if response.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			response.Body.String(), expected)
	}
}

func createRequest(verb, url string, input Food) (*http.Request, error) {
	buff := &bytes.Buffer{}
	err := json.NewEncoder(buff).Encode(input)

	if err != nil {
		return nil, err
	}

	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	return http.NewRequest(verb, url, buff)
}

func createExpected(expectedResult Food) (string, error) {
	// Check the response body is what we expect.
	bytes, err := json.Marshal(expectedResult)

	if err != nil {
		return "", err
	}

	return string(bytes), err
}
