# E-commerce Application

A modern e-commerce application built with Go, featuring a clean architecture, RESTful APIs, and comprehensive e-commerce functionality.

## Project Structure

```
├── cmd/                  # Application entry points
│   └── server/          # HTTP server implementation
├── docs/                # Documentation
│   └── swagger.yaml     # API documentation
├── internal/            # Private application code
│   ├── config/         # Configuration management
│   ├── db/             # Database setup and migrations
│   ├── handlers/       # HTTP request handlers
│   ├── middleware/     # HTTP middleware components
│   ├── models/         # Data models
│   ├── repository/     # Data access layer
│   ├── router/         # HTTP routing setup
│   └── service/        # Business logic layer
├── pkg/                 # Shared utilities
│   └── logger/         # Logging utility
├── static/              # Static assets
│   └── css/            # CSS stylesheets
├── templates/           # HTML templates
├── .env                 # Environment configuration
├── go.mod              # Go module file
└── README.md           # Project documentation
```

## Features

### Authentication & Authorization
- User registration and login
- Password reset functionality
- JWT-based authentication
- Role-based access control (Admin/User)

### Product Management
- Product listing with pagination
- Product search and filtering
- Product details view
- Admin product management (CRUD operations)

### Shopping Cart
- Add/remove products
- Update quantities
- Cart persistence

### Order Management
- Order creation and processing
- Order history
- Admin order management
- Order status updates

### Admin Dashboard
- Overview statistics
- Product management
- Order management
- User management

### Technical Features
- Clean and organized project structure
- RESTful API design
- Custom logging utility
- Database migrations
- Swagger API documentation
- Secure password handling
- Environment-based configuration

## Getting Started

1. Clone the repository
2. Copy `.env.example` to `.env` and configure your environment variables
3. Install dependencies: `go mod tidy`
4. Run the server: `go run cmd/server/main.go`
5. Visit `http://localhost:8080` in your browser

## API Documentation

API documentation is available via Swagger at `/swagger/index.html` when running the server.

## Development

This project follows the standard Go project layout and best practices for organizing Go code. It implements a clean architecture pattern with clear separation of concerns:

- Handlers: HTTP request handling
- Services: Business logic
- Repositories: Data access
- Models: Data structures

## License

MIT License