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
  Isbn string `json:"isbn"`
  Name string `json:"name"`
  Calories string `json:"calories"`
}
// Meal Struct (Model)
type Meal struct {
  ID string `json:"id"`
  Isbn string `json:"isbn"`
  Name string `json:"name"`
}

//GET all foods
func getFoods(w http.ResponseWriter, r *http.Request)  {

}
//GET single food
func getFood(w http.ResponseWriter, r *http.Request)  {

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

  //Route handlers / endpoints
  router.HandleFunc("/api/v1/foods", getFoods).Methods("GET")
  router.HandleFunc("/api/v1/foods/{id}", getFood).Methods("GET")
  router.HandleFunc("/api/v1/foods", createFood).Methods("POST")
  router.HandleFunc("/api/v1/foods/{id}", updateFood).Methods("PUT")
  router.HandleFunc("/api/v1/foods/{id}", deleteFood).Methods("DELETE")

  log.Fatal(http.ListenAndServe(":8000", router))
}
