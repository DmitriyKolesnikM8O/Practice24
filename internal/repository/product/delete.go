package product

import (
	"context"
	"fmt"
	"github.com/DmitriyKolesnikM8O/Practice24/internal/repository/product/requests"
)

func (r *Repository) Delete(ctx context.Context, id string) error {
	r.logger.Tracef(fmt.Sprintf("SQL Query: %s", formatQuery(requests.RequestDelete)))
	_, err := r.client.Exec(ctx, requests.RequestDelete, id)
	if err != nil {
		return err
	}

	return nil
}
