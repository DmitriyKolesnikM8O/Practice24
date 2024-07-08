package product

import (
	"context"
	"fmt"
	"github.com/DmitriyKolesnikM8O/Practice24/internal/repository/product/model"
	"github.com/DmitriyKolesnikM8O/Practice24/internal/repository/product/requests"
)

func (r *Repository) FindAll(ctx context.Context) (p []model.Product, err error) {
	r.logger.Tracef(fmt.Sprintf("SQL Query: %s", formatQuery(requests.RequestFindAll)))
	rows, err := r.client.Query(ctx, requests.RequestFindAll)
	if err != nil {

		return nil, err
	}

	products := make([]model.Product, 0)
	for rows.Next() {
		var prod model.Product
		err := rows.Scan(&prod.ID, &prod.Name, &prod.Price, &prod.Count, &prod.Date)
		if err != nil {
			return nil, err
		}

		products = append(products, prod)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}
