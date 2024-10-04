# GOTTH Stack Template

This repository serves as a fullstack application template built with Go, Templ, TailwindCSS, HTMX, and other modern technologies. It provides a solid foundation for building scalable web applications with server-side rendering, frontend styling, and easy interactivity.

## Project Directory

```plaintext
GOTTH-stack
├── bin/                          # Binary outputs after build
├── cmd/                          # Main application entry points
│   └── main.go                   # Main Go application file
├── config/                       # Configuration files for different environments
├── Dockerfile                    # Docker configuration file for building and running the app
├── internal/                     # Core internal modules (handlers, routes, templates)
│   ├── handlers/                 # API and web route handlers
│   ├── routes/                   # Route definitions
│   └── templates/                # .templ files for server-side rendering (using Templ)
├── node_modules/                 # Frontend dependencies (installed via npm)
├── pkg/                          # Shared packages
│   ├── htmx/                     # HTMX integration and helper functions
│   ├── logger/                   # Logging utility using slog
│   ├── middleware/               # Custom middlewares for Echo framework
│   ├── services/                 # Business logic and services layer
│   ├── types/                    # Custom types and models
│   ├── utils/                    # Utility functions and helpers
│   └── validator/                # Input validation utilities
├── static/                       # Static assets (e.g., JS, CSS, images)
│   ├── styles.css                # Main input CSS file for Tailwind
│   └── dist/
│       └── output.css            # Compiled CSS output (after Tailwind build)
├── store/                        # Database connection and setup (SQLite, MySQL, PostgreSQL)
├── utils/                        # Additional utility functions
├── .gitignore                    # Files and directories to ignore in git
├── Makefile                      # Makefile for building, running, and testing the app
├── go.mod                        # Go module definition
├── go.sum                        # Go module dependencies lock file
└── README.md                     # Project documentation
```

## Features

- **Go backend**: Fast, scalable, and lightweight using the Echo framework.
- **Templ**: Templating language for server-side rendering with `.templ` files.
- **TailwindCSS**: Utility-first CSS framework for fast and modern styling.
- **HTMX**: Enables partial updates and enhances interactivity without JavaScript frameworks.
- **SQLite/MySQL/PostgreSQL support**: Default database setup with SQLite, easily configurable to use MySQL, PostgreSQL, or NoSQL.
- **Viper for CLI configuration**: Environment-based configuration management.
- **Logging with slog**: Structured logging for better observability.
- **Validation**: Input validation using `validator` for form and API request safety.
- **Cloud Storage Scalability**: Extendable to include Google Cloud, AWS S3, or any other cloud provider for static asset and data storage.
- **Docker Support**: Containerized application for easy deployment on cloud platforms like AWS and Google Cloud.

## How to Use

### 1. Clone the repository

```bash
git clone https://github.com/yourusername/GOTTH-stack.git
cd GOTTH-stack
```
### 2. Remove the .git 

### 3. Follow TODO.md for setup and running the application

---

