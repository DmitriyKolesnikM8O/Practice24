package product

import (
	repository2 "github.com/DmitriyKolesnikM8O/Practice24/internal/repository"
	postgresql "github.com/DmitriyKolesnikM8O/Practice24/pkg/client/postgres"
	"github.com/DmitriyKolesnikM8O/Practice24/pkg/logging"
	"strings"
)

type Repository struct {
	client postgresql.Client
	logger *logging.Logger
}

func NewRepository(client postgresql.Client, logger *logging.Logger) repository2.Repository {
	return &Repository{
		client: client,
		logger: logger,
	}
}

func formatQuery(q string) string {
	return strings.ReplaceAll(strings.ReplaceAll(q, "\t", ""), "\n", " ")
}
