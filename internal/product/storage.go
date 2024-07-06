package product

import (
	"context"
)

// любая БД должна реализовывать эти методы
type Repository interface {
	Create(ctx context.Context, product *CreateProduct) (err error)
	FindAll(ctx context.Context) (p []Product, err error)
	FindOne(ctx context.Context, id string) (Product, error)
	Update(ctx context.Context, product UpdateProduct) error
	FindAllForReport(ctx context.Context) (rep []Product, res MonthSales, err error)
	Delete(ctx context.Context, id string) error
}
