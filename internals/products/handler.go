package products

import (
	"github.com/YounessBrunno/Golang-E-commerce-Project/internals/json"
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
   products := []string{"Product 1", "Product 2", "Product 3"}
   
   json.Write(w, http.StatusOK, products)
   
}