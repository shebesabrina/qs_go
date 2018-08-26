package main

// import (
//     "database/sql"
//     "errors"
// )

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

//
// func (p *product) getProduct(db *sql.DB) error {
//   return errors.New("Not implemented")
// }
//
// func (p *product) updateProduct(db *sql.DB) error {
//   return errors.New("Not implemented")
// }
//
// func (p *product) deleteProduct(db *sql.DB) error {
//   return errors.New("Not implemented")
// }
//
// func (p *product) createProduct(db *sql.DB) error {
//   return errors.New("Not implemented")
// }
//
// func getProducts(db *sql.DB, start, count int) ([]product, error) {
//   return nil, errors.New("Not implemented")
// }
