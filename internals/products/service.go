package products

import "context"

type ProductService interface {
   ListProducts(ctx context.Context) ([]string, error)
}

type ProductSvc struct {
   // ProductRepo ProductRepository
}

func NewProductService() ProductService {
   return &ProductSvc{
	  // ProductRepo: repo,
   }
}

func (s *ProductSvc) ListProducts(ctx context.Context) ([]string, error) {
   return []string{"Product 1", "Product 2", "Product 3"}, nil
}