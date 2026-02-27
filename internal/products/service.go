package products

import (
	"context"
	"log"

	repo "github.com/jack/ecom/internal/adapters/postgresql/sqlc"
)

type Service interface {
	ListProducts(ctx context.Context) ([]repo.Product, error)
	FindProductById(ctx context.Context, id int64) (repo.Product, error)
}

type svc struct {
	// repository
	repo repo.Querier
}

func NewService(repo repo.Querier) Service {
	return &svc{
		repo: repo,
	}
}

func (s *svc) ListProducts(ctx context.Context) ([]repo.Product, error) {
	if products, err := s.repo.ListProducts(ctx); err != nil {
		log.Fatalf("err:", err)
	} else {
		return products, nil
	}
	return nil, nil
}

func (s *svc) FindProductById(ctx context.Context, id int64) (repo.Product, error) {
	if product, err := s.repo.FindProductByID(ctx, id); err != nil {
		log.Fatalf("err:", err)
	} else {
		return product, nil
	}
	return repo.Product{}, nil
}
