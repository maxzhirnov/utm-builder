# UTM Builder

A modern web application for generating UTM-tagged URLs with a clean, responsive interface.

## Features

- **Real-time URL Generation**: Generate UTM-tagged URLs instantly
- **Modern UI**: Clean, responsive interface built with Tailwind CSS
- **Configurable Options**: Customize sources and mediums through YAML configuration
- **Custom Inputs**: Support for custom source and medium values
- **Copy to Clipboard**: Easy copying of generated URLs
- **Hot Reloading**: Development server with automatic reloading using Air

## Tech Stack

### Backend
- **Go**: Core application logic
- **Standard Library**: Using `net/http` for HTTP server and routing, `html/template` for templates
- **YAML**: Configuration management using `gopkg.in/yaml.v3`

### Frontend
- **HTMX**: Dynamic UI updates without page reloads
- **Tailwind CSS**: Modern, utility-first CSS framework
- **PostCSS**: CSS processing and optimization

## Development Setup

1. **Install Dependencies**:
   ```bash
   # Install Go dependencies
   go mod download

   # Install Node.js dependencies
   npm install
   ```

2. **Build CSS**:
   ```bash
   npm run build:css
   ```

3. **Run the Application**:
   ```bash
   # Development mode with hot reloading
   make run

   # Production build
   make build
   ```

## Configuration

The application uses a YAML configuration file (`config/utm.yaml`) to manage UTM parameters:

```yaml
sources:
  - google
  - facebook
  - twitter
  - linkedin
  - custom

mediums:
  - cpc
  - cpm
  - email
  - social
  - banner
  - custom
```

To update the configuration:
1. Edit `config/utm.yaml`
2. Restart the server

## Project Structure

```
.
├── assets/           # Static assets
│   └── css/         # CSS files
├── cmd/             # Application entry points
│   └── server/      # Main server
├── config/          # Configuration files
│   └── utm.yaml     # UTM parameters configuration
├── internal/        # Internal packages
│   ├── config/      # Configuration management
│   ├── handlers/    # HTTP handlers
│   ├── templates/   # HTML templates
│   └── utm/         # UTM URL generation
├── Makefile         # Build and run commands
└── README.md        # Project documentation
```

## Available Commands

- `make run`: Start the development server with hot reloading
- `make build`: Build the production binary
- `make clean`: Clean build artifacts
- `make setup`: Install dependencies and build CSS

## Contributing

1. Fork the repository
2. Create a feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request

## License

MIT License
