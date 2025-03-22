package router

import (
	"ecommerce-app/internal/handlers"
	"ecommerce-app/internal/middleware"
	"ecommerce-app/internal/service"
	"net/http"
)

// SetupUserRoutes configures user-related routes
func SetupUserRoutes(userService service.UserService, authService service.AuthService) {
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