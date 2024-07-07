package product

import (
	"context"
	"fmt"
	repository2 "github.com/DmitriyKolesnikM8O/Practice24/internal/repository"
	"github.com/DmitriyKolesnikM8O/Practice24/internal/repository/auth"
	"github.com/DmitriyKolesnikM8O/Practice24/internal/repository/product/model"
	"github.com/DmitriyKolesnikM8O/Practice24/internal/repository/product/requests"

	"github.com/DmitriyKolesnikM8O/Practice24/pkg/client/postgres"
	"github.com/DmitriyKolesnikM8O/Practice24/pkg/logging"
	"github.com/jackc/pgconn"
	"strings"
)

type repository struct {
	client postgresql.Client
	logger *logging.Logger
}

func NewRepository(client postgresql.Client, logger *logging.Logger) repository2.Repository {
	return &repository{
		client: client,
		logger: logger,
	}
}

func formatQuery(q string) string {
	return strings.ReplaceAll(strings.ReplaceAll(q, "\t", ""), "\n", " ")
}

func (r *repository) Create(ctx context.Context, product *model.CreateProduct) (err error) {
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

func (r *repository) FindAll(ctx context.Context) (p []model.Product, err error) {
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

func (r *repository) FindOne(ctx context.Context, id string) (model.Product, error) {
	r.logger.Tracef(fmt.Sprintf("SQL Query: %s", formatQuery(requests.RequestFindOne)))
	var prod model.Product
	err := r.client.QueryRow(ctx, requests.RequestFindOne, id).Scan(&prod.ID, &prod.Name, &prod.Price, &prod.Count, &prod.Date)
	if err != nil {
		return model.Product{}, err
	}

	return prod, nil
}

func (r *repository) Update(ctx context.Context, product model.UpdateProduct) error {
	r.logger.Tracef(fmt.Sprintf("SQL Query: %s", formatQuery(requests.RequestUpdate)))
	_, err := r.client.Exec(ctx, requests.RequestUpdate, product.ID, product.Name, product.Price, product.Count)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) FindAllForReport(ctx context.Context) (rep []model.Product, res model.MonthSales, err error) {
	var resSales model.MonthSales
	r.logger.Tracef(fmt.Sprintf("SQL Query: %s", formatQuery(requests.RequestFindAllForReport)))
	rows, err := r.client.Query(ctx, requests.RequestFindAllForReport)
	if err != nil {
		return nil, resSales, err
	}

	products := make([]model.Product, 0)
	for rows.Next() {
		var prod model.Product
		err := rows.Scan(&prod.ID, &prod.Name, &prod.Price, &prod.Count, &prod.Date)
		resSales.Sales += prod.Count * prod.Price
		resSales.Counts += prod.Count
		if err != nil {
			return nil, resSales, err
		}

		products = append(products, prod)
	}
	if err := rows.Err(); err != nil {
		return nil, resSales, err
	}

	return products, resSales, nil
}

func (r *repository) Delete(ctx context.Context, id string) error {
	r.logger.Tracef(fmt.Sprintf("SQL Query: %s", formatQuery(requests.RequestDelete)))
	_, err := r.client.Exec(ctx, requests.RequestDelete, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) CreateUser(ctx context.Context, user *auth.User) error {
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

func (r *repository) FindOneOnUsersTable(ctx context.Context, name string) (username string, err error) {
	r.logger.Tracef(fmt.Sprintf("SQL Query: %s", formatQuery(requests.RequestFindOneFromUser)))
	var nameUser string
	err = r.client.QueryRow(ctx, requests.RequestFindOneFromUser, name).Scan(&nameUser)
	if err != nil {
		return "", err
	}

	return nameUser, nil
}

func (r *repository) FindOneUser(ctx context.Context, username string) (auth.User, error) {
	r.logger.Tracef(fmt.Sprintf("SQL Query: %s", formatQuery(requests.RequestFindOneUser)))
	var OneUser auth.User

	err := r.client.QueryRow(ctx, requests.RequestFindOneUser, username).Scan(&OneUser.ID, &OneUser.Username, &OneUser.Password)
	if err != nil {
		return auth.User{}, err
	}

	return OneUser, nil
}
