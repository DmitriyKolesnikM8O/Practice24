package product

import (
	"context"
	"fmt"
	"github.com/DmitriyKolesnikM8O/Practice24/internal/repository/product/model"
	"github.com/DmitriyKolesnikM8O/Practice24/internal/repository/product/requests"
	"github.com/jackc/pgconn"
)

func (r *Repository) Create(ctx context.Context, product *model.CreateProduct) (err error) {
	r.logger.Tracef(fmt.Sprintf("SQL Query: %s", formatQuery(requests.RequestCreate)))
	err = r.client.QueryRow(ctx, requests.RequestCreate, product.Name, product.Price, product.Count, product.Date).Scan(&product.Name)
	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			newErr := fmt.Errorf("SQL error: %s", pgErr.Message)
			return newErr
		}
	}

	return nil
}
