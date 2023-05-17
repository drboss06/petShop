package model

type Cart struct {
	Id     int    `json:"id"`
	Code   string `json:"code"`
	Name   string `json:"name"`
	Price  int    `json:"price"`
	Items  int    `json:"items"`
	Dprice int    `json:"dprice"`
}
