package product

import (
	"context"
	"fmt"
	"github.com/DmitriyKolesnikM8O/Practice24/internal/repository/auth"
	"github.com/DmitriyKolesnikM8O/Practice24/internal/repository/product/requests"
	"github.com/jackc/pgconn"
)

func (r *Repository) CreateUser(ctx context.Context, user *auth.User) error {
	r.logger.Tracef(fmt.Sprintf("SQL Query: %s", requests.RequestCreateUser))
	err := r.client.QueryRow(ctx, requests.RequestCreateUser, user.Username, user.Password).Scan(user.ID)
	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			newErr := fmt.Errorf("SQL error: %s", pgErr.Message)
			return newErr
		}
	}

	return nil
}
