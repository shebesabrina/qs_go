package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"math/rand"
	// "github.com/rs/cors"
	"github.com/gorilla/handlers"
	"os"
)

func getPort() string {
  p := os.Getenv("PORT")
  if p != "" {
    return ":" + p
  }
  return ":3000"
}

// Init foods var as a slice food struct
var foods []Food

func init() {
	// Mock data - @todo - implement DB
	foods = append(foods, Food{ID: "1", Name: "Banana", Calories: "100"})
	foods = append(foods, Food{ID: "2", Name: "Pie", Calories: "300"})
	foods = append(foods, Food{ID: "3", Name: "Hot dog", Calories: "10"})
	foods = append(foods, Food{ID: "4", Name: "Strawberries", Calories: "100"})
}


//GET all foods
func getFoods(w http.ResponseWriter, r *http.Request)  {
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(foods)
}
//GET single food
func getFood(w http.ResponseWriter, r *http.Request)  {
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r) //get params
  //Loop through books and find with id
  for _, item := range foods {
    if item.ID == params["id"] {
      json.NewEncoder(w).Encode(item)
      return
    }
  }
  json.NewEncoder(w).Encode(&Food{})
}
//POST create food item
func createFood(w http.ResponseWriter, r *http.Request)  {
  w.Header().Set("Content-Type", "application/json")
  var food Food
  _ = json.NewDecoder(r.Body).Decode(&food)
  food.ID = strconv.Itoa(rand.Intn(10000)) //Mock ID - not safe for production
  foods = append(foods, food)
  json.NewEncoder(w).Encode(food)
}

// PUT update food item
func UpdateFood(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	// fmt.Printf("params: %+v\n", params)
	for index, item := range foods {
		// fmt.Printf("item: %+v\n", item)
		if item.ID == params["id"] {
			var inputFood Food
			err := json.NewDecoder(r.Body).Decode(&inputFood)
      fmt.Printf("decode: %+v\n", inputFood)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(err.Error()))
				return
			}
      if inputFood.ID != item.ID {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("ids do not match"))
				return
			}
			foods[index] = inputFood
			json.NewEncoder(w).Encode(inputFood)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Food not found"))
}


//GET delete food item
func deleteFood(w http.ResponseWriter, r *http.Request)  {
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r)
  for index, item := range foods {
    if item.ID == params["id"] {
      foods = append(foods[:index], foods[index+1:]...)
      break
    }
  }
  json.NewEncoder(w).Encode(foods)
}


func main() {
	//Init Router
	router := mux.NewRouter()

	port := getPort()

	allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With"})
  allowedOrigins := handlers.AllowedOrigins([]string{"*"})
  allowedMethods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})


	//Route handlers / endpoints
	router.HandleFunc("/api/v1/foods", getFoods).Methods("GET")
	router.HandleFunc("/api/v1/foods/{id}", getFood).Methods("GET")
	router.HandleFunc("/api/v1/foods", createFood).Methods("POST")
	router.HandleFunc("/api/v1/foods/{id}", UpdateFood).Methods("PUT")
	router.HandleFunc("/api/v1/foods/{id}", deleteFood).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(
		allowedHeaders, allowedOrigins, allowedMethods)(router)))
	// log.Fatal(http.ListenAndServe(":8000", router))
}
