package product

import (
	"context"
	"fmt"
	"github.com/DmitriyKolesnikM8O/Practice24/internal/repository/product/requests"
)

func (r *Repository) FindOneOnUsersTable(ctx context.Context, name string) (username string, err error) {
	r.logger.Tracef(fmt.Sprintf("SQL Query: %s", formatQuery(requests.RequestFindOneFromUser)))
	var nameUser string
	err = r.client.QueryRow(ctx, requests.RequestFindOneFromUser, name).Scan(&nameUser)
	if err != nil {
		return "", err
	}

	return nameUser, nil
}
