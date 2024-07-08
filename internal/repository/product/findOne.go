package product

import (
	"context"
	"fmt"
	"github.com/DmitriyKolesnikM8O/Practice24/internal/repository/product/model"
	"github.com/DmitriyKolesnikM8O/Practice24/internal/repository/product/requests"
)

func (r *Repository) FindOne(ctx context.Context, id string) (model.Product, error) {
	r.logger.Tracef(fmt.Sprintf("SQL Query: %s", formatQuery(requests.RequestFindOne)))
	var prod model.Product
	err := r.client.QueryRow(ctx, requests.RequestFindOne, id).Scan(&prod.ID, &prod.Name, &prod.Price, &prod.Count, &prod.Date)
	if err != nil {
		return model.Product{}, err
	}

	return prod, nil
}
