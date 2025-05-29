package entity

import "github.com/shopspring/decimal"

type Medicine struct {
	ID    int
	Name  string
	Stock int
	Price decimal.Decimal
}

type ReqAddMedicine struct {
	Name  string
	Stock int
	Price decimal.Decimal
}
