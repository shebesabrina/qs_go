package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
  "github.com/gorilla/mux"
)

func TestGetAllFoodsSuccessfully(t *testing.T) {
		testFood := Food{
			ID:       "1",
			Name:     "Banana",
			Calories: "100",
		}

    request, err := http.NewRequest("GET", "api/v1/foods", testFood)
		if err != nil {
			t.Fatal(err)
		}
		response := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc("/api/v1/foods", UpdateFood).Methods("GET")
		router.ServeHTTP(response, request)

		if status := response.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
}

// func TestPutSuccess(t *testing.T) {
// 	testFood := Food{
// 		ID:       "1",
// 		Name:     "Banana",
// 		Calories: "100",
// 	}
// 	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
// 	// pass 'nil' as the third parameter.
// 	request, err := createRequest("PUT", "/api/v1/foods/1", testFood)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
// 	response := httptest.NewRecorder()
// 	router := mux.NewRouter()
// 	router.HandleFunc("/api/v1/foods/{id}", UpdateFood).Methods("PUT")
// 	router.ServeHTTP(response, request)
//
// 	// Check the status code is what we expect.
// 	if status := response.Code; status != http.StatusOK {
// 		t.Errorf("handler returned wrong status code: got %v want %v",
// 			status, http.StatusOK)
// 	}
//
// 	expected, err := parseResponse(response)
//
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	if testFood != expected {
// 		t.Errorf("handler returned unexpected body: got %v want %v",
// 			response.Body.String(), expected)
// 	}
// }
//
// func TestPutNoIdMatch(t *testing.T) {
// 	testFood := Food{
// 		ID:       "2",
// 		Name:     "Banana",
// 		Calories: "100",
// 	}
// 	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
// 	// pass 'nil' as the third parameter.
// 	request, err := createRequest("PUT", "/api/v1/foods/1", testFood)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
// 	response := httptest.NewRecorder()
// 	router := mux.NewRouter()
// 	router.HandleFunc("/api/v1/foods/{id}", UpdateFood).Methods("PUT")
// 	router.ServeHTTP(response, request)
//
// 	// Check the status code is what we expect.
// 	if status := response.Code; status != http.StatusBadRequest {
// 		t.Errorf("handler returned wrong status code: got %v want %v",
// 			status, http.StatusBadRequest)
// 	}
// }
//
// func TestPutNoBody(t *testing.T) {
// 	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
// 	// pass 'nil' as the third parameter.
//   buff := bytes.NewBufferString("")
//
// 	request, err := http.NewRequest("PUT", "/api/v1/foods/1", buff)
//   if err != nil {
//     t.Fatal(err)
//   }
// 	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
// 	response := httptest.NewRecorder()
// 	router := mux.NewRouter()
// 	router.HandleFunc("/api/v1/foods/{id}", UpdateFood).Methods("PUT")
// 	router.ServeHTTP(response, request)
//
// 	// Check the status code is what we expect.
// 	if status := response.Code; status != http.StatusBadRequest {
// 		t.Errorf("handler returned wrong status code: got %v want %v",
// 			status, http.StatusBadRequest)
// 	}
// }
//
// func TestPutNoId(t *testing.T) {
// 	testFood := Food{
// 		ID:       "1",
// 		Name:     "Banana",
// 		Calories: "100",
// 	}
// 	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
// 	// pass 'nil' as the third parameter.
// 	request, err := createRequest("PUT", "/api/v1/foods/10000", testFood)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
// 	response := httptest.NewRecorder()
// 	router := mux.NewRouter()
// 	router.HandleFunc("/api/v1/foods/{id}", UpdateFood).Methods("PUT")
// 	router.ServeHTTP(response, request)
//
// 	// Check the status code is what we expect.
// 	if status := response.Code; status != http.StatusNotFound {
// 		t.Errorf("handler returned wrong status code: got %v want %v",
// 			status, http.StatusNotFound)
// 	}
// }

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

func parseResponse(response *httptest.ResponseRecorder) (Food, error) {
	// Check the response body is what we expect.
  	responseFood := Food{}

	 err := json.NewDecoder(response.Body).Decode(&responseFood)

	if err != nil {
		return responseFood, err
	}

	return responseFood, nil
}
