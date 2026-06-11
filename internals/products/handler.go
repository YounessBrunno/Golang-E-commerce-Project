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
   products, err := h.productService.ListProducts(r.Context())

   if  err != nil {
       json.Write(w, http.StatusInternalServerError, map[string]string{"error": "Failed to list products"})

   }
   
   json.Write(w, http.StatusOK, products)
   
}