# UTM Builder

A modern web application for generating UTM-tagged URLs, built with Go, HTMX, and Tailwind CSS.

## Features

- 🚀 Simple and intuitive interface for generating UTM-tagged URLs
- ⚡ Real-time URL generation with HTMX
- 🎨 Modern UI with Tailwind CSS
- 📋 One-click URL copying with visual feedback
- 🔄 Hot reloading for development
- 📱 Responsive design

## Tech Stack

### Backend
- **Go** - Fast and efficient server-side language
- **net/http** - Go's standard library HTTP server with built-in routing
- **html/template** - Go's standard library template engine

### Frontend
- **HTMX** - For dynamic, AJAX-like functionality without writing JavaScript
- **Tailwind CSS** - Utility-first CSS framework for rapid UI development
- **Modern JavaScript** - For clipboard functionality

## Development Setup

1. Install Go (1.24.2 or later)
2. Install Node.js and npm
3. Clone the repository
4. Install dependencies:
   ```bash
   make deps
   ```
5. Run the development server:
   ```bash
   make run
   ```

## Available Commands

- `make run` - Start the development server with hot reload
- `make run-prod` - Run the production server
- `make build` - Build the application
- `make clean` - Clean build artifacts
- `make setup` - Set up the development environment
- `make deps` - Install all dependencies

## Project Structure

```
.
├── assets/
│   └── css/           # CSS files (Tailwind)
├── cmd/
│   └── server/        # Main application entry point
├── internal/
│   ├── handlers/      # HTTP handlers
│   ├── templates/     # HTML templates
│   ├── utm/          # UTM parameter handling
│   └── dictionary/   # Predefined UTM values
├── .air.toml         # Hot reload configuration
├── Makefile          # Build and run commands
├── tailwind.config.js # Tailwind configuration
└── package.json      # Frontend dependencies
```

## Why This Stack?

- **Go**: Provides excellent performance and a simple, maintainable codebase
- **HTMX**: Enables dynamic functionality without complex JavaScript frameworks
- **Tailwind CSS**: Allows rapid UI development with utility classes
- **Hot Reloading**: Speeds up development with instant feedback

## Contributing

1. Fork the repository
2. Create a feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request

## License

MIT
