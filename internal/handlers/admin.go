package handlers

import (
	"ecommerce-app/internal/models"
	"ecommerce-app/internal/service"
	"ecommerce-app/pkg/logger"
	"encoding/json"
	"math"
	"net/http"
	"strconv"
)

type AdminHandler struct {
    productService service.ProductService
    orderService   service.OrderService
    userService    service.UserService
    log            *logger.Logger
}

func NewAdminHandler(productService service.ProductService, orderService service.OrderService, userService service.UserService) *AdminHandler {
    return &AdminHandler{
        productService: productService,
        orderService:   orderService,
        userService:    userService,
        log:            logger.New(),
    }
}

// GetDashboardStats returns overview statistics for the admin dashboard
func (h *AdminHandler) GetDashboardStats(w http.ResponseWriter, r *http.Request) {
    var stats struct {
        TotalProducts int64 `json:"total_products"`
        TotalOrders   int64 `json:"total_orders"`
        TotalUsers    int64 `json:"total_users"`
    }

    var err error
    stats.TotalProducts, err = h.productService.CountProducts()
    if err != nil {
        h.log.Error("Failed to count products: " + err.Error())
        http.Error(w, "Failed to get dashboard stats", http.StatusInternalServerError)
        return
    }
    
    stats.TotalOrders, err = h.orderService.CountOrders()
    if err != nil {
        h.log.Error("Failed to count orders: " + err.Error())
        http.Error(w, "Failed to get dashboard stats", http.StatusInternalServerError)
        return
    }
    
    stats.TotalUsers, err = h.userService.CountUsers()
    if err != nil {
        h.log.Error("Failed to count users: " + err.Error())
        http.Error(w, "Failed to get dashboard stats", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(stats)
}

// ListProducts returns products for admin management with pagination
func (h *AdminHandler) ListProducts(w http.ResponseWriter, r *http.Request) {
    // Parse pagination parameters
    page := 1
    pageSize := 10
    
    // Get page from query string
    pageStr := r.URL.Query().Get("page")
    if pageStr != "" {
        if pageVal, err := strconv.Atoi(pageStr); err == nil && pageVal > 0 {
            page = pageVal
        }
    }
    
    // Get page size from query string
    pageSizeStr := r.URL.Query().Get("pageSize")
    if pageSizeStr != "" {
        if pageSizeVal, err := strconv.Atoi(pageSizeStr); err == nil && pageSizeVal > 0 {
            pageSize = pageSizeVal
        }
    }
    
    // Get products with pagination
    products, err := h.productService.GetProductsPaginated(page, pageSize)
    if err != nil {
        h.log.Error("Failed to fetch products: " + err.Error())
        http.Error(w, "Failed to fetch products", http.StatusInternalServerError)
        return
    }
    
    // Get total count for pagination metadata
    total, err := h.productService.CountProducts()
    if err != nil {
        h.log.Error("Failed to count products: " + err.Error())
        http.Error(w, "Failed to fetch products", http.StatusInternalServerError)
        return
    }
    
    // Create response with pagination metadata
    response := struct {
        Products []models.Product `json:"products"`
        Pagination struct {
            Total       int64 `json:"total"`
            Page        int   `json:"page"`
            PageSize    int   `json:"pageSize"`
            TotalPages  int   `json:"totalPages"`
        } `json:"pagination"`
    }{
        Products: products,
    }
    
    response.Pagination.Total = total
    response.Pagination.Page = page
    response.Pagination.PageSize = pageSize
    response.Pagination.TotalPages = int(math.Ceil(float64(total) / float64(pageSize)))
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

// CreateProduct handles new product creation
func (h *AdminHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
    var product models.Product
    if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
        h.log.Error("Invalid product data: " + err.Error())
        http.Error(w, "Invalid product data", http.StatusBadRequest)
        return
    }

    if err := h.productService.CreateProduct(&product); err != nil {
        h.log.Error("Failed to create product: " + err.Error())
        http.Error(w, "Failed to create product", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(product)
}

// ListOrders returns orders for admin management with pagination
func (h *AdminHandler) ListOrders(w http.ResponseWriter, r *http.Request) {
    // Parse pagination parameters
    page := 1
    pageSize := 10
    
    // Get page from query string
    pageStr := r.URL.Query().Get("page")
    if pageStr != "" {
        if pageVal, err := strconv.Atoi(pageStr); err == nil && pageVal > 0 {
            page = pageVal
        }
    }
    
    // Get page size from query string
    pageSizeStr := r.URL.Query().Get("pageSize")
    if pageSizeStr != "" {
        if pageSizeVal, err := strconv.Atoi(pageSizeStr); err == nil && pageSizeVal > 0 {
            pageSize = pageSizeVal
        }
    }
    
    // Get orders with pagination
    orders, err := h.orderService.GetOrdersWithUserPaginated(page, pageSize)
    if err != nil {
        h.log.Error("Failed to fetch orders: " + err.Error())
        http.Error(w, "Failed to fetch orders", http.StatusInternalServerError)
        return
    }
    
    // Get total count for pagination metadata
    total, err := h.orderService.CountOrders()
    if err != nil {
        h.log.Error("Failed to count orders: " + err.Error())
        http.Error(w, "Failed to fetch orders", http.StatusInternalServerError)
        return
    }
    
    // Create response with pagination metadata
    response := struct {
        Orders []models.Order `json:"orders"`
        Pagination struct {
            Total       int64 `json:"total"`
            Page        int   `json:"page"`
            PageSize    int   `json:"pageSize"`
            TotalPages  int   `json:"totalPages"`
        } `json:"pagination"`
    }{
        Orders: orders,
    }
    
    response.Pagination.Total = total
    response.Pagination.Page = page
    response.Pagination.PageSize = pageSize
    response.Pagination.TotalPages = int(math.Ceil(float64(total) / float64(pageSize)))
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

// UpdateOrderStatus handles order status updates
func (h *AdminHandler) UpdateOrderStatus(w http.ResponseWriter, r *http.Request) {
    var update struct {
        OrderID uint   `json:"order_id"`
        Status  string `json:"status"`
    }

    if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
        h.log.Error("Invalid update data: " + err.Error())
        http.Error(w, "Invalid update data", http.StatusBadRequest)
        return
    }

    if err := h.orderService.UpdateOrderStatus(update.OrderID, update.Status); err != nil {
        h.log.Error("Failed to update order status: " + err.Error())
        http.Error(w, "Failed to update order status", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
}