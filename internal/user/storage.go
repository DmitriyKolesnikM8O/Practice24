package user

import "context"

// Create(ctx context.Context, user User) (string, error)
type Storage interface {
	Create(ctx context.Context, product Product) (string, error)
	FindOne(ctx context.Context, id string) (Product, error)
	Update(ctx context.Context, product Product) error
	Delete(ctx context.Context, id string) error
}
