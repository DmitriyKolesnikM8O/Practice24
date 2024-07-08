package product

import (
	"context"
	"fmt"
	"github.com/DmitriyKolesnikM8O/Practice24/internal/repository/product/model"
	"github.com/DmitriyKolesnikM8O/Practice24/internal/repository/product/requests"
)

func (r *Repository) Update(ctx context.Context, product model.UpdateProduct) error {
	r.logger.Tracef(fmt.Sprintf("SQL Query: %s", formatQuery(requests.RequestUpdate)))
	_, err := r.client.Exec(ctx, requests.RequestUpdate, product.ID, product.Name, product.Price, product.Count)
	if err != nil {
		return err
	}

	return nil
}
