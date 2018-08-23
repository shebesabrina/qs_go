package main
import(
  "fmt"
  "encoding/json"
  "log"
  "net/http"
  "math/rand"
  "strconv"
  "github.com/gorilla/mux"
)
func main()  {
  //Init Router
  router := Mux.NewRouter()

  //Route handlers / endpoints
  router.HandleFunc("/api/v1/foods", getFoods).Methods("GET")
  router.HandleFunc("/api/v1/foods/{id}", getFood).Methods("GET")
  router.HandleFunc("/api/v1/foods", createFood).Methods("POST")
  router.HandleFunc("/api/v1/foods/{id}", updateFood).Methods("PUT")
  router.HandleFunc("/api/v1/foods/{id}", deleteFood).Methods("DELETE")
}
