package product

import (
	"context"
	"fmt"
	"github.com/DmitriyKolesnikM8O/Practice24/internal/product"
	"github.com/DmitriyKolesnikM8O/Practice24/pkg/client/postgres"
	"github.com/DmitriyKolesnikM8O/Practice24/pkg/logging"
	"strings"
)

type repository struct {
	client postgresql.Client
	logger *logging.Logger
}

func formatQuery(q string) string {
	return strings.ReplaceAll(strings.ReplaceAll(q, "\t", ""), "\n", " ")
}

// реализация всех методов интерфейса
func (r *repository) Create(ctx context.Context, product *product.CreateProduct) (err error) {
	q := `
			INSERT INTO product 
			    (name, price, count, date) 
			VALUES 
			    ($1, $2, $3, $4) 
			RETURNING id
	`
	r.logger.Tracef(fmt.Sprintf("SQL Query: %s", formatQuery(q)))
	if err := r.client.QueryRow(ctx, q, product.Name, product.Price, product.Count, product.Date); err != nil {

		return nil
	}
	return nil
}

func (r *repository) FindAll(ctx context.Context) (p []product.Product, err error) {
	q := `
			SELECT 
				id, name, price, count, date 
			FROM 
				public.product
	`

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

func (r *repository) FindOne(ctx context.Context, id string) (product.Product, error) {
	q := `
			SELECT 
				id, name, price, count, date 
			FROM 
				public.product 
			WHERE id = $1
	`

	r.logger.Tracef(fmt.Sprintf("SQL Query: %s", formatQuery(q)))

	var prod product.Product
	err := r.client.QueryRow(ctx, q, id).Scan(&prod.ID, &prod.Name, &prod.Price, &prod.Count, &prod.Date)
	if err != nil {
		return product.Product{}, err
	}

	return prod, nil
}

func (r *repository) Update(ctx context.Context, product product.UpdateProduct) error {
	q := `
			UPDATE 
			    public.product 
			SET 
			    name = $2, price = $3, count = $4 
			WHERE id = $1
	`

	r.logger.Tracef(fmt.Sprintf("SQL Query: %s", formatQuery(q)))
	_, err := r.client.Exec(ctx, q, product.ID, product.Name, product.Price, product.Count)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) FindAllForReport(ctx context.Context) (rep []product.Product, res product.MonthSales, err error) {
	q := `
			SELECT 
				id, name, price, count, date
			FROM 
				public.product 
			WHERE 
				date >= CURRENT_DATE - INTERVAL '1 month' 
			GROUP BY 
				id, name, price, count, date
			ORDER BY 
				date
	`

	var resSales product.MonthSales
	r.logger.Tracef(fmt.Sprintf("SQL Query: %s", formatQuery(q)))
	rows, err := r.client.Query(ctx, q)
	if err != nil {
		return nil, resSales, err
	}
	//массив для всех данных
	products := make([]product.Product, 0)

	//идем по выдаче
	for rows.Next() {
		var prod product.Product

		//записываем в переменные структуры
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
	q := `
			DELETE FROM 
			    public.product 
			WHERE 
			    id = $1
	`

	r.logger.Tracef(fmt.Sprintf("SQL Query: %s", formatQuery(q)))

	_, err := r.client.Exec(ctx, q, id)
	if err != nil {
		return err
	}

	return nil

}

func NewRepository(client postgresql.Client, logger *logging.Logger) product.Repository {
	return &repository{
		client: client,
		logger: logger,
	}
}
