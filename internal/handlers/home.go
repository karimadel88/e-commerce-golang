package handlers

import (
	"ecommerce-app/internal/models"
	"ecommerce-app/internal/service"
	"encoding/json"
	"math"
	"net/http"
)

// HomeHandler returns a handler function that fetches and returns products with pagination
func HomeHandler(productService service.ProductService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Ensure the request is for the root path
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		// Default pagination values
		page := 1
		pageSize := 6

		// Get products with pagination
		products, err := productService.GetProductsPaginated(page, pageSize)
		if err != nil {
			http.Error(w, "Failed to fetch products", http.StatusInternalServerError)
			return
		}

		// Get total count for pagination metadata
		total, err := productService.CountProducts()
		if err != nil {
			http.Error(w, "Failed to count products", http.StatusInternalServerError)
			return
		}

		// Create response with pagination metadata
		response := struct {
			Products []models.Product `json:"products"`
			Pagination struct {
				Total      int64 `json:"total"`
				Page       int   `json:"page"`
				PageSize   int   `json:"pageSize"`
				TotalPages int   `json:"totalPages"`
			} `json:"pagination"`
		}{
			Products: products,
		}

		response.Pagination.Total = total
		response.Pagination.Page = page
		response.Pagination.PageSize = pageSize
		response.Pagination.TotalPages = int(math.Ceil(float64(total) / float64(pageSize)))

		// Set response headers for JSON
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		// Encode response as JSON and send
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}
	}
}
