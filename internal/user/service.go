package user

import (
	"context"
	"github.com/DmitriyKolesnikM8O/Practice24/pkg/logging"
)

type Service struct {
	storage Storage
	logger  *logging.Logger
}

func (s *Service) Create(ctx context.Context, dto CreateProductDTO) (p Product, err error) {
	//TODO
	return
}
