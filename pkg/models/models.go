package models

type Produce struct {
	ID           int     `json:"id"`
	Produce_Code string  `json:"prod_code"`
	Name         string  `json:"name"`
	Unit_Price   float64 `json:"price"`
}