# E-commerce Application

A modern e-commerce application built with Go, featuring a clean architecture and standard project layout.

## Project Structure

```
├── cmd/                 # Application entry points
│   └── server/         # HTTP server implementation
├── pkg/                 # Shared utilities
│   └── logger/         # Logging utility
├── static/             # Static assets
│   └── css/           # CSS stylesheets
├── templates/          # HTML templates
│   └── *.html         # Template files
├── go.mod             # Go module file
└── README.md          # Project documentation
```

## Features

- Clean and organized project structure
- Custom logging utility
- Responsive web interface
- Modern CSS styling

## Getting Started

1. Clone the repository
2. Install dependencies: `go mod tidy`
3. Run the server: `go run cmd/server/main.go`
4. Visit `http://localhost:8080` in your browser

## Development

This project follows the standard Go project layout and best practices for organizing Go code.

## License

MIT License