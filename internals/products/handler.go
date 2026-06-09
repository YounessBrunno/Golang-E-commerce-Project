package products


type Handler struct {
   productService ProductService
}

func NewHandler(productService ProductService) *Handler {
   return &Handler{
	 productService: productService,
   }
}