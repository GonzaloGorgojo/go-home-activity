<h2>Auth Flow</h2>

User SignUP and get a short lived jwt as response and a longlived jwt created that is stored in db.
User Login and get the short lived token as response and long lived refreshed.
Middleware for protected routes check for short jwt, in case is expired search for long lived, in case this is expired it send you to login,
in case this is not expired it also refresh the token.
Also all steps check if user is suspended or not. -> when login, signup also check for this, and middleware for expired short lived token will try to search for refresh token, and when doing this it will check status of user


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
│   ├── users/
│   │   ├── user_repository.go   # UserRepository interface
│   │   ├── user_repository_impl.go  # DB implementation
│   │   ├── user_service.go   # User service
│   │   ├── user_handler.go   # User HTTP handlers
│   │   ├── user.go         # User model
│   │   ├── user_service_test.go
│   │
│   ├── routes/             # Router setup
│   │   ├── routes.go        # API routes
│   │   └── middleware.go    # Global middleware (Auth, Logging)
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

