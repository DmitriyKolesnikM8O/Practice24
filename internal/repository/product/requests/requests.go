package requests

var (
	RequestCreate = `
			INSERT INTO product 
			    (name, price, count, date) 
			VALUES 
			    ($1, $2, $3, $4) 
			RETURNING id
	`

	RequestFindAll = `
			SELECT 
				id, name, price, count, date model
			FROM 
				public.product
	`

	RequestFindOne = `
			SELECT 
				id, name, price, count, date 
			FROM 
				public.product 
			WHERE id = $1
	`

	RequestUpdate = `
			UPDATE 
			    public.product 
			SET 
			    name = $2, price = $3, count = $4 
			WHERE id = $1
	`

	RequestFindAllForReport = `
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

	RequestDelete = `
			DELETE FROM 
			    public.product 
			WHERE 
			    id = $1
	`
)
