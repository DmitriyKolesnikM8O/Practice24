package product

import (
	"github.com/jackc/pgtype"
)

// New table
type Product struct {
	ID    int         `json:"id"`
	Name  string      `json:"name"`
	Price int         `json:"price"`
	Count int         `json:"count"`
	Date  pgtype.Date `json:"date"`
}

type Report struct {
	Name        string      `json:"name"`
	TotalPrice  int         `json:"total_price"`
	TotalCount  int         `json:"total_count"`
	GeneralSale int         `json:"general_sale"`
	Date        pgtype.Date `json:"date"`
}

type MonthSales struct {
	Counts int `json:"month_counts"`
	Sales  int `json:"month_sales"`
}
