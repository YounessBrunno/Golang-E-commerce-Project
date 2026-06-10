package products

import (
	"encoding/json"
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

   w.Header().Set("Content-Type", "application/json")
   w.WriteHeader(http.StatusCreated)
   
   
   if err := json.NewEncoder(w).Encode(products); err != nil {
      http.Error(w, "Failed to encode products", http.StatusInternalServerError)
   }
}