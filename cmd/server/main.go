package main

import (
	"ecommerce-app/internal/config"
	"ecommerce-app/internal/db"
	"ecommerce-app/internal/repository"
	"ecommerce-app/internal/router"
	"ecommerce-app/internal/service"
	"ecommerce-app/pkg/logger"
	"net/http"
)

func main() {
    log := logger.New()

    // Load configuration
    cfg, err := config.Load()
    if err != nil {
        log.Error("Failed to load config: " + err.Error())
        return
    }

    // Connect to database
    dbConn, err := db.Connect(cfg.DatabaseURL)
    if err != nil {
        log.Error("Failed to connect to database: " + err.Error())
        return
    }
    defer func() {
        sqlDB, _ := dbConn.DB()
        sqlDB.Close()
    }()
    
    // Set global database connection
    db.SetDB(dbConn)

    // Run migrations
    if err := db.Migrate(dbConn); err != nil {
        log.Error("Failed to migrate database: " + err.Error())
        return
    }

    // Test database connection and migrations
    if err := db.TestConnection(cfg.DatabaseURL); err != nil {
        log.Error("Database connection test failed: " + err.Error())
        return
    }

    // Initialize repositories
    productRepo := repository.NewProductRepository(dbConn)
    orderRepo := repository.NewOrderRepository(dbConn)
    userRepo := repository.NewUserRepository(dbConn)
    
    // Initialize services
    productService := service.NewProductService(productRepo)
    orderService := service.NewOrderService(orderRepo)
    userService := service.NewUserService(userRepo)
    authService := service.NewAuthService(userService)
    
    // Seed test user for development/testing
    // testEmail := "test@example.com"
    // testPassword := "password123"
    // log.Info("Seeding test user...")
    // _, err = authService.SeedTestUser(testEmail, testPassword)
    // if err != nil {
    //     log.Error("Failed to seed test user: " + err.Error())
    //     // Continue execution even if seeding fails
    // } else {
    //     log.Info("Test user available with email: " + testEmail)
    // }
   // Seed database with sample products
    // log.Info("Seeding database with sample products...")
    // if err := db.SeedProducts(productService); err != nil {
    //     log.Error("Failed to seed products: " + err.Error())
    //     // Continue execution even if seeding fails
    // }

    // log.Info("Done seeding database")
    
    // Setup routes using the router package
    router.SetupRoutes(authService, userService, productService, orderService)

    // Start server
    log.Info("Server starting on port " + cfg.Port)
    err = http.ListenAndServe(":"+cfg.Port, nil)
    if err != nil {
        log.Error("Error starting server: " + err.Error())
    }
}