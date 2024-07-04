package product

import (
	"context"
	"fmt"
	"github.com/DmitriyKolesnikM8O/Practice24/internal/product"
	"github.com/DmitriyKolesnikM8O/Practice24/pkg/client/postgres"
	"github.com/DmitriyKolesnikM8O/Practice24/pkg/logging"
	"github.com/jackc/pgconn"
	"strings"
)

type repository struct {
	client postgresql.Client
	logger *logging.Logger
}

func formatQuery(q string) string {
	return strings.ReplaceAll(strings.ReplaceAll(q, "\t", ""), "\n", " ")
}

func (r *repository) Create(ctx context.Context, product *product.Product) error {
	q := `
			INSERT INTO product 
			    (name, price, count) 
			VALUES 
			    ($1, $2, $3) 
			RETURNING id
	`

	if err := r.client.QueryRow(ctx, q, product.Name, product.Price, product.Count).Scan(&product.ID); err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s, Details: %s, Where: %s, Code: %s",
				pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code))
			r.logger.Error(newErr)
			return newErr
		}
		return err
	}
	return nil
}

// TODO Maybe change to CreateReport OR change FindOne to CreateReport?????
func (r *repository) FindAll(ctx context.Context) (p []product.Product, err error) {
	q := `SELECT id, name, price, count FROM public.product`

	r.logger.Tracef(fmt.Sprintf("SQL Query: %s", formatQuery(q)))
	rows, err := r.client.Query(ctx, q)
	if err != nil {

		return nil, err
	}

	//массив для всех данных
	products := make([]product.Product, 0)

	//идем по выдаче
	for rows.Next() {
		var prod product.Product

		//записываем в переменные структуры
		err := rows.Scan(&prod.ID, &prod.Name, &prod.Price, &prod.Count)

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

func (r *repository) FindOne(ctx context.Context, id string) (product.Product, error) {
	q := `SELECT id, name, price, count FROM public.product WHERE id = $1`

	r.logger.Tracef(fmt.Sprintf("SQL Query: %s", formatQuery(q)))

	var prod product.Product
	err := r.client.QueryRow(ctx, q, id).Scan(&prod.ID, &prod.Name, &prod.Price, &prod.Count)
	if err != nil {
		return product.Product{}, err
	}

	return prod, nil
}

func (r *repository) Update(ctx context.Context, product product.Product) error {
	//TODO implement me
	panic("implement me")
}

func (r *repository) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func NewRepository(client postgresql.Client, logger *logging.Logger) product.Repository {
	return &repository{
		client: client,
		logger: logger,
	}
}
