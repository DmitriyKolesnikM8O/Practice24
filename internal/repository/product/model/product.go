package model

import (
	"github.com/jackc/pgtype"
)

type Product struct {
	ID    int         `json:"id"`
	Name  string      `json:"name"`
	Price int         `json:"price"`
	Count int         `json:"count"`
	Date  pgtype.Date `json:"date"`
}

type MonthSales struct {
	Counts int `json:"month_counts"`
	Sales  int `json:"month_sales"`
}

type CreateProduct struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
	Count int    `json:"count"`
	Date  string `json:"date"`
}

type UpdateProduct struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Count int    `json:"count"`
}

type CombinedResponse struct {
	FirstType  *[]Product  `json:"all_sales_products,omitempty"`
	SecondType *MonthSales `json:"general_statistics,omitempty"`
}
