package products

import ("context"
        repo "github.com/YounessBrunno/Golang-E-commerce-Project/internals/adapters/postgresql/sqlc" 
       )                                                         
        

type ProductService interface {
   ListProducts(ctx context.Context) ([]repo.Product, error)
}

type productSvc struct {
   repo repo.Querier
}
func NewProductService(repo repo.Querier) ProductService {
   return &productSvc{
	  repo: repo,
   }
}

func (s *productSvc) ListProducts(ctx context.Context) ([]repo.Product, error) {
   return s.repo.ListProducts(ctx)
}