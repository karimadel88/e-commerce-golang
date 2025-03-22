package router

import (
	"ecommerce-app/internal/handlers"
	"ecommerce-app/internal/middleware"
	"ecommerce-app/internal/service"
	"net/http"
)

// SetupAdminRoutes configures admin-related routes
func SetupAdminRoutes(adminHandler *handlers.AdminHandler, authService service.AuthService) {
	// Admin routes with authentication
	http.HandleFunc("/admin/dashboard", middleware.AdminAuth(adminHandler.GetDashboardStats))
	http.HandleFunc("/admin/products", middleware.AdminAuth(adminHandler.ListProducts))
	http.HandleFunc("/admin/products/create", middleware.AdminAuth(adminHandler.CreateProduct))
	http.HandleFunc("/admin/orders", middleware.AdminAuth(adminHandler.ListOrders))
	http.HandleFunc("/admin/orders/update-status", middleware.AdminAuth(adminHandler.UpdateOrderStatus))
}