package dto

import "github.com/shopspring/decimal"

type Medicine struct {
	ID    int             `json:"id"`
	Name  string          `json:"name"`
	Stock int             `json:"stock"`
	Price decimal.Decimal `json:"price"`
}

type ReqAddMedicine struct {
	Name  string          `json:"name" binding:"required"`
	Stock int             `json:"stock" binding:"required"`
	Price decimal.Decimal `json:"price" binding:"required"`
}
