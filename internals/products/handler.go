package products

import (
   "net/http"
)

type Handler struct {
   productService ProductService
}

func NewHandler(productService ProductService) *Handler {
   return &Handler{
	 productService: productService,
   }
}

func (h *Handler) ListProducts(w http.ResponseWriter, r *http.Request) {

}