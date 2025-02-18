Project folder structure:

```
│── cmd/                    # Entry points for different executables
│   └── server/             # HTTP server
│       └── main.go         # Main application entry point
│
├── internal/               # Private application code
│   ├── database/           # Database connection & initialization
│   │   ├── db.go           # Database setup
│   │   ├── migrations/     # SQL migrations (optional)
│   │   └── seed/           # Seed data scripts (optional)
│   │
│   ├── repositories/       # Data access layer
│   │   ├── user_repository.go   # UserRepository interface
│   │   ├── user_repository_impl.go  # DB implementation
│   │   ├── product_repository.go    # Example for products
│   │   └── ...
│   │
│   ├── services/           # Business logic layer
│   │   ├── user_service.go   # User service
│   │   ├── product_service.go # Example for products
│   │   └── ...
│   │
│   ├── handlers/           # HTTP handlers (controllers)
│   │   ├── user_handler.go   # User HTTP handlers
│   │   ├── product_handler.go # Example for products
│   │   └── ...
│   │
│   ├── routes/             # Router setup
│   │   ├── routes.go        # API routes
│   │   └── middleware.go    # Global middleware (Auth, Logging)
│   │
│   ├── models/             # Data models
│   │   ├── user.go         # User model
│   │   ├── product.go      # Example for products
│   │   └── ...
│   │
│   ├── config/             # Configuration files (env, settings)
│   │   ├── config.go       # Load environment variables
│   │   └── app.env         # Environment variables
│   │
│   ├── utils/              # Utility functions/helpers
│   │   ├── logger.go       # Custom logger
│   │   ├── response.go     # Common API response helpers
│   │   └── ...
│   │
│   ├── tests/              # Unit and integration tests
│   │   ├── user_service_test.go
│   │   ├── product_service_test.go
│   │   └── ...
│   │
│   └── middleware/         # Middleware functions (auth, logging)
│       ├── auth.go         # Authentication middleware
│       ├── logging.go      # Request logging middleware
│       ├── cors.go         # CORS handling
│       └── ...
│
├── pkg/                    # Reusable packages (if needed)
│   ├── auth/               # JWT or OAuth logic
│   ├── cache/              # Caching logic (Redis, in-memory)
│   ├── email/              # Email sending logic
│   └── ...
│
├── vendor/                 # Third-party dependencies (Go modules)
│
├── .env                    # Environment variables file
├── .gitignore              # Git ignore file
├── go.mod                  # Go module dependencies
├── go.sum                  # Go checksum file
└── README.md               # Project documentation
```
