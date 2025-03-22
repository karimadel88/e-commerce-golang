package db

import (
	"ecommerce-app/internal/models"
	"ecommerce-app/internal/service"
	"ecommerce-app/pkg/logger"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// Product categories and adjectives for generating random product names
var (
	categories = []string{"Laptop", "Smartphone", "Headphones", "Camera", "Tablet", "Smartwatch", "Speaker", "Monitor", "Keyboard", "Mouse"}
	adjectives = []string{"Pro", "Ultra", "Premium", "Elite", "Max", "Lite", "Plus", "Extreme", "Smart", "Wireless"}
	brands     = []string{"TechX", "Gadgetify", "ElectroPro", "DigiMax", "SmartTech", "FutureTech", "InnoGear", "NextGen", "PrimeTech", "VisionTech"}
)

// SeedProducts creates 100 sample products in the database
func SeedProducts(productService service.ProductService) error {
	log := logger.New()
	log.Info("Starting to seed 100 products...")

	// Seed random number generator
	rand.Seed(time.Now().UnixNano())

	// Create 100 products
	for i := 1; i <= 1000000; i++ {
		// Generate random product data
		product := generateRandomProduct()

		// Create product in database
		err := productService.CreateProduct(product)
		if err != nil {
			log.Error(fmt.Sprintf("Failed to create product %d: %s", i, err.Error()))
			return err
		}

		if i%10 == 0 {
			log.Info(fmt.Sprintf("Created %d products", i))
		}
	}

	log.Info("Successfully seeded 100 products")
	return nil
}

// generateRandomProduct creates a random product with realistic data
func generateRandomProduct() *models.Product {
	// Pick random elements for product name
	brand := brands[rand.Intn(len(brands))]
	category := categories[rand.Intn(len(categories))]
	adjective := adjectives[rand.Intn(len(adjectives))]

	// Generate model number
	modelNum := fmt.Sprintf("%s-%d", strings.ToUpper(category[:3]), 1000+rand.Intn(9000))

	// Create product name
	name := fmt.Sprintf("%s %s %s %s", brand, adjective, category, modelNum)

	// Generate description
	description := fmt.Sprintf("The %s %s %s is a high-quality %s featuring the latest technology. "+
		"This %s model offers exceptional performance and reliability for all your needs. "+
		"Model: %s", brand, adjective, category, strings.ToLower(category), strings.ToLower(category), modelNum)

	// Generate price between $50 and $2000
	price := 50.0 + rand.Float64()*1950.0
	// Round to 2 decimal places
	price = float64(int(price*100)) / 100

	// Generate stock between 5 and 200
	stock := 5 + rand.Intn(196)

	return &models.Product{
		Name:        name,
		Description: description,
		Price:       price,
		Stock:       stock,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}