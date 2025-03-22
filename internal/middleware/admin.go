package middleware

import (
    "net/http"
    "ecommerce-app/pkg/logger"
    "os"
    "crypto/subtle"
)

// AdminAuth middleware ensures that only authenticated admin users can access protected routes
func AdminAuth(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        log := logger.New()

        // Get admin credentials from environment variables
        adminUser := os.Getenv("ADMIN_USERNAME")
        adminPass := os.Getenv("ADMIN_PASSWORD")

        // If admin credentials are not set, use defaults (not recommended for production)
        if adminUser == "" {
            adminUser = "admin"
        }
        if adminPass == "" {
            adminPass = "admin123"
        }

        // Get Basic Auth credentials
        user, pass, ok := r.BasicAuth()

        if !ok || subtle.ConstantTimeCompare([]byte(user), []byte(adminUser)) != 1 || 
           subtle.ConstantTimeCompare([]byte(pass), []byte(adminPass)) != 1 {
            w.Header().Set("WWW-Authenticate", `Basic realm="Admin Area")`)
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            log.Error("Unauthorized access attempt to admin area")
            return
        }

        next.ServeHTTP(w, r)
    }
}