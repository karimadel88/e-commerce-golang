package router

import (
	"ecommerce-app/internal/db"
	"ecommerce-app/internal/handlers"
	"ecommerce-app/internal/middleware"
	"ecommerce-app/internal/repository"
	"ecommerce-app/internal/service"
	"net/http"
)

// SetupRoutes configures all application routes
func SetupRoutes(authService service.AuthService, userService service.UserService, 
	productService service.ProductService, orderService service.OrderService) {
	// Initialize cart repository and service
	cartRepo := repository.NewCartRepository(db.GetDB())
	cartService := service.NewCartService(cartRepo)
	
	// Initialize handlers
	authHandler := handlers.NewAuthHandler(authService)
	adminHandler := handlers.NewAdminHandler(productService, orderService, userService)
	
	// Setup route groups
	setupAuthRoutes(authHandler)
	setupAdminRoutes(adminHandler, authService)
	setupUserRoutes(userService, authService, cartService)
	
	// Basic handler (to test)
	http.HandleFunc("/", handlers.HomeHandler(productService))

	// Swagger documentation route
}

// setupAuthRoutes configures authentication-related routes
func setupAuthRoutes(authHandler *handlers.AuthHandler) {
	// Authentication routes
	http.HandleFunc("/auth/register", authHandler.Register)
	http.HandleFunc("/auth/login", authHandler.Login)
	http.HandleFunc("/auth/reset-password-request", authHandler.RequestPasswordReset)
	http.HandleFunc("/auth/reset-password", authHandler.ResetPassword)
}

// setupAdminRoutes configures admin-related routes
func setupAdminRoutes(adminHandler *handlers.AdminHandler, authService service.AuthService) {
	// Admin routes with authentication
	http.HandleFunc("/admin/dashboard", middleware.AdminAuth(adminHandler.GetDashboardStats))
	http.HandleFunc("/admin/products", middleware.AdminAuth(adminHandler.ListProducts))
	http.HandleFunc("/admin/products/create", middleware.AdminAuth(adminHandler.CreateProduct))
	http.HandleFunc("/admin/orders", middleware.AdminAuth(adminHandler.ListOrders))
	http.HandleFunc("/admin/orders/update-status", middleware.AdminAuth(adminHandler.UpdateOrderStatus))
}

// setupUserRoutes configures user-related routes
func setupUserRoutes(userService service.UserService, authService service.AuthService, cartService service.CartService) {
	// Initialize cart handler
	cartHandler := handlers.NewCartHandler(cartService)
	
	// Cart routes
	http.HandleFunc("/user/cart", middleware.UserAuth(authService)(cartHandler.GetCart))
	http.HandleFunc("/user/cart/add", middleware.UserAuth(authService)(cartHandler.AddToCart))
	http.HandleFunc("/user/cart/remove", middleware.UserAuth(authService)(cartHandler.RemoveFromCart))
	// User routes with authentication
	http.HandleFunc("/user/profile", middleware.UserAuth(authService)(func(w http.ResponseWriter, r *http.Request) {
		userID, ok := middleware.GetUserID(r)
		if !ok {
			http.Error(w, "User not found", http.StatusInternalServerError)
			return
		}
		
		user, err := userService.GetUserByID(userID)
		if err != nil {
			http.Error(w, "Failed to get user profile", http.StatusInternalServerError)
			return
		}
		
		handlers.ResponseWithJSON(w, map[string]interface{}{
			"success": true,
			"data": map[string]interface{}{
				"id": user.ID,
				"email": user.Email,
			},
		}, http.StatusOK)
	}))
}