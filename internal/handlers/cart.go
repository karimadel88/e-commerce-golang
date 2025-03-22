package handlers

import (
	"ecommerce-app/internal/middleware"
	"ecommerce-app/internal/service"
	"encoding/json"
	"net/http"
	"strconv"
)

// CartHandler handles cart-related HTTP requests
type CartHandler struct {
	cartService service.CartService
}

// NewCartHandler creates a new instance of CartHandler
func NewCartHandler(cartService service.CartService) *CartHandler {
	return &CartHandler{
		cartService: cartService,
	}
}

// GetCart handles retrieving a user's cart
func (h *CartHandler) GetCart(w http.ResponseWriter, r *http.Request) {
	userID, ok := middleware.GetUserID(r)
	if !ok {
		ResponseWithJSON(w, map[string]interface{}{"error": "Unauthorized"}, http.StatusUnauthorized)
		return
	}

	cartItems, err := h.cartService.GetCart(userID)
	if err != nil {
		ResponseWithJSON(w, map[string]interface{}{"error": "Failed to get cart"}, http.StatusInternalServerError)
		return
	}

	ResponseWithJSON(w, map[string]interface{}{"cart": cartItems}, http.StatusOK)
}

// AddToCart handles adding a product to the user's cart
func (h *CartHandler) AddToCart(w http.ResponseWriter, r *http.Request) {
	userID, ok := middleware.GetUserID(r)
	if !ok {
		ResponseWithJSON(w, map[string]interface{}{"error": "Unauthorized"}, http.StatusUnauthorized)
		return
	}

	var cartItem struct {
		ProductID uint `json:"product_id"`
		Quantity  int  `json:"quantity"`
	}
	
	if err := json.NewDecoder(r.Body).Decode(&cartItem); err != nil {
		ResponseWithJSON(w, map[string]interface{}{"error": "Invalid request"}, http.StatusBadRequest)
		return
	}

	err := h.cartService.AddToCart(userID, cartItem.ProductID, cartItem.Quantity)
	if err != nil {
		ResponseWithJSON(w, map[string]interface{}{"error": "Failed to add to cart"}, http.StatusInternalServerError)
		return
	}

	ResponseWithJSON(w, map[string]interface{}{"message": "Item added to cart"}, http.StatusCreated)
}

// RemoveFromCart handles removing a product from the user's cart
func (h *CartHandler) RemoveFromCart(w http.ResponseWriter, r *http.Request) {
	userID, ok := middleware.GetUserID(r)
	if !ok {
		ResponseWithJSON(w, map[string]interface{}{"error": "Unauthorized"}, http.StatusUnauthorized)
		return
	}

	productID := r.URL.Query().Get("product_id")
	if productID == "" {
		ResponseWithJSON(w, map[string]interface{}{"error": "Product ID required"}, http.StatusBadRequest)
		return
	}

	// Convert productID string to uint
	prodID, err := strconv.ParseUint(productID, 10, 32)
	if err != nil {
		ResponseWithJSON(w, map[string]interface{}{"error": "Invalid product ID"}, http.StatusBadRequest)
		return
	}

	err = h.cartService.RemoveFromCart(userID, uint(prodID))
	if err != nil {
		ResponseWithJSON(w, map[string]interface{}{"error": "Failed to remove from cart"}, http.StatusInternalServerError)
		return
	}

	ResponseWithJSON(w, map[string]interface{}{"message": "Item removed from cart"}, http.StatusOK)
}