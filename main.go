package main
import(
  // "fmt"
  // "encoding/json"
  "log"
  "net/http"
  // "math/rand"
  // "strconv"
  "github.com/gorilla/mux"
)
//Food Struct (Model)
type Food struct {
  ID string `json:"id"`
  Name string `json:"name"`
  Calories string `json:"calories"`
}
// Meal Struct (Model)
type Meal struct {
  ID string `json:"id"`
  Isbn string `json:"isbn"`
  Name string `json:"name"`
}

// Init foods var as a slice food struct
var foods []Food

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
//GET create food item
func createFood(w http.ResponseWriter, r *http.Request)  {

}
//GET update food item
func updateFood(w http.ResponseWriter, r *http.Request)  {

}
//GET delete food item
func deleteFood(w http.ResponseWriter, r *http.Request)  {

}

func main()  {
  //Init Router
  router := mux.NewRouter()

  // Mock data - @todo - implement DB
  food = append(foods, Food{ID: "1", Name: "Banana", Calories: "100", Meal: &Meal{Name: "Breakfast"}})
  food = append(foods, Food{ID: "2", Name: "Pie", Calories: "300", Meal: &Meal{Name: "Snack"}})
  food = append(foods, Food{ID: "3", Name: "Hot dog", Calories: "10", Meal: &Meal{Name: "Lunch"}})
  food = append(foods, Food{ID: "4", Name: "Strawberries", Calories: "100", Meal: &Meal{Name: "Dinner"}})

  //Route handlers / endpoints
  router.HandleFunc("/api/v1/foods", getFoods).Methods("GET")
  router.HandleFunc("/api/v1/foods/{id}", getFood).Methods("GET")
  router.HandleFunc("/api/v1/foods", createFood).Methods("POST")
  router.HandleFunc("/api/v1/foods/{id}", updateFood).Methods("PUT")
  router.HandleFunc("/api/v1/foods/{id}", deleteFood).Methods("DELETE")

  log.Fatal(http.ListenAndServe(":8000", router))
}
