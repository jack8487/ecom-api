package products

import (
	"context"

	repo "github.com/jack/ecom/internal/adapters/postgresql/sqlc"
)

type Service interface {
	ListProducts(ctx context.Context) ([]repo.Product, error)
	FindProductById(ctx context.Context, id int64) (repo.Product, error)
	CreateProduct(ctx context.Context, name string, priceInCenter int32, quantity int32) (repo.Product, error)
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
		// log.Fatalf("err: %v", err)
		return nil, err
	} else {
		return products, nil
	}
}

func (s *svc) FindProductById(ctx context.Context, id int64) (repo.Product, error) {
	if product, err := s.repo.FindProductByID(ctx, id); err != nil {
		// log.Fatalf("err: %v", err)
		return repo.Product{}, err
	} else {
		return product, nil
	}
}

func (s *svc) CreateProduct(ctx context.Context, name string, priceInCenter int32, quantity int32) (repo.Product, error) {
	if product, err := s.repo.CreateProduct(ctx, repo.CreateProductParams{
		Name:          name,
		PriceInCenter: priceInCenter,
		Quantity:      quantity,
	}); err != nil {
		// log.Fatalf("err: %v", err)
		return repo.Product{}, err
	} else {
		return product, nil
	}
}
