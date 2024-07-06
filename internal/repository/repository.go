package repository

import (
	"context"
	"github.com/DmitriyKolesnikM8O/Practice24/internal/repository/product/model"
)

type Repository interface {
	Create(ctx context.Context, product *model.CreateProduct) (err error)
	FindAll(ctx context.Context) (p []model.Product, err error)
	FindOne(ctx context.Context, id string) (model.Product, error)
	Update(ctx context.Context, product model.UpdateProduct) error
	FindAllForReport(ctx context.Context) (rep []model.Product, res model.MonthSales, err error)
	Delete(ctx context.Context, id string) error
}
