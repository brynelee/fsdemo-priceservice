package model

import (
	"time"
)

type ProductPrice struct {
	ProductId   int       `json:"ProductId"`
	ProductName string    `json:"ProductName"`
	Price       float64   `json:"Price"`
	UpdateTime  time.Time `json:"UpdateTime"`
}
