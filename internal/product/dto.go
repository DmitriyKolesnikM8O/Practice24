package product

type CreateProductDTO struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
	Count int    `json:"count"`
}
