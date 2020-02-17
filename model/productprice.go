package model

import (
	"time"
)

type ProductPrice struct {
	ProductId   int
	ProductName string
	Price       float64
	UpdateTime  time.Time
}
