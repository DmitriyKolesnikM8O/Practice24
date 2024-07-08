package product

import (
	"context"
	"fmt"
	"github.com/DmitriyKolesnikM8O/Practice24/internal/repository/product/model"
	"github.com/DmitriyKolesnikM8O/Practice24/internal/repository/product/requests"
)

func (r *Repository) FindAllForReport(ctx context.Context) (rep []model.Product, res model.MonthSales, err error) {
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
