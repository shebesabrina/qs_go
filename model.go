package main

//Food Struct (Model)
type Food struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Calories string `json:"calories"`
}

// Meal Struct (Model)
type Meal struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
