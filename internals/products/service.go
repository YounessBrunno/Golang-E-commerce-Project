package products

import "context"

type ProductService interface {
   ListProducts(ctx context.Context) ([]string, error)
}

type productSvc struct {
   // ProductRepo ProductRepository
}

func NewProductService() ProductService {
   return &productSvc{
	  // ProductRepo: repo,
   }
}

func (s *productSvc) ListProducts(ctx context.Context) ([]string, error) {
   return []string{"Product 1", "Product 2", "Product 3"}, nil
}