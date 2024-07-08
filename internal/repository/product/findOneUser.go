package product

import (
	"context"
	"fmt"
	"github.com/DmitriyKolesnikM8O/Practice24/internal/repository/auth"
	"github.com/DmitriyKolesnikM8O/Practice24/internal/repository/product/requests"
)

func (r *Repository) FindOneUser(ctx context.Context, username string) (auth.User, error) {
	r.logger.Tracef(fmt.Sprintf("SQL Query: %s", formatQuery(requests.RequestFindOneUser)))
	var OneUser auth.User

	err := r.client.QueryRow(ctx, requests.RequestFindOneUser, username).Scan(&OneUser.ID, &OneUser.Username, &OneUser.Password)
	if err != nil {
		return auth.User{}, err
	}

	return OneUser, nil
}
