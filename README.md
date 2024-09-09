# Golang Transaction Engine

This project is a scalable backend system for processing financial transactions, demonstrating key concepts in Go programming and financial technology.

## Features

- Account creation and retrieval
- Transaction processing with concurrent safety
- In-memory data storage with thread-safe operations
- RESTful API design
- Graceful shutdown
- Structured logging
- Comprehensive error handling and input validation
- Middleware for authentication, rate limiting, and CORS
- Unit and integration testing

## Project Structure

```
go-transaction-engine/
├── cmd/
│   └── server/
│       └── main.go           # Application entry point
├── internal/
│   ├── api/
│   │   ├── handlers.go       # HTTP request handlers
│   │   └── middleware.go     # HTTP middleware (authentication, rate limiting, CORS)
│   ├── models/
│   │   ├── account.go        # Account model definition
│   │   └── transaction.go    # Transaction model definition
│   ├── store/
│   │   └── memory.go         # In-memory data store implementation
│   └── service/
│       └── transaction.go    # Business logic for transaction processing
├── pkg/
│   ├── logger/
│   │   └── logger.go         # Custom logger implementation
│   └── utils/
│       └── string_utils.go   # Utility functions for string operations
├── tests/
│   └── api_test.go           # API integration tests
├── go.mod
├── go.sum
└── README.md
```

## How It Works

1. **Server Initialization**: The main function in `cmd/server/main.go` initializes the logger, in-memory store, and API handlers. It then starts the HTTP server and sets up graceful shutdown.

2. **API Handlers**: The `internal/api/handlers.go` file defines HTTP handlers for creating accounts, retrieving account information, processing transactions, and fetching transaction history.

3. **Middleware**: The `internal/api/middleware.go` file implements authentication, rate limiting, and CORS middleware.

4. **Data Store**: The `internal/store/memory.go` file implements an in-memory store for accounts and transactions. It uses a read-write mutex to ensure thread-safe operations.

5. **Service Layer**: The `internal/service/transaction.go` file contains the business logic for processing transactions.

6. **Models**: The `internal/models` directory contains struct definitions for `Account` and `Transaction`.

7. **Logging**: A custom logger is implemented in `pkg/logger/logger.go` to provide structured logging throughout the application.

8. **Utilities**: The `pkg/utils/string_utils.go` file contains utility functions, including a random string generator.

## Key Design Decisions

1. **Concurrent Safety**: The project uses mutexes to ensure thread-safe operations on the in-memory store, allowing for concurrent request handling.

2. **Separation of Concerns**: The project structure separates API handling, business logic, and data storage, making it easier to maintain and scale.

3. **Middleware**: Implementation of authentication, rate limiting, and CORS middleware demonstrates security and performance considerations.

4. **Graceful Shutdown**: The server implements a graceful shutdown mechanism, ensuring that ongoing requests are completed before the server stops.

5. **RESTful API**: The API follows RESTful principles, making it intuitive and easy to use.

6. **Error Handling**: Proper error handling is implemented throughout the application, with errors being logged and appropriate HTTP status codes returned.

7. **Testing**: Unit and integration tests are included to ensure reliability and ease of maintenance.

8. **Logging**: A custom logging interface is implemented, allowing for easy adaptation to different logging backends.

9. **Code Documentation**: Comprehensive comments are included throughout the codebase, improving readability and maintainability.

## Testing

The project includes both unit tests and integration tests. The `tests/api_test.go` file contains integration tests for the API endpoints. To run the tests, use the following command:

```
go test ./...
```

## Logging

A custom logging interface is implemented in `pkg/logger/logger.go`. This allows for structured logging and easy switching between different logging backends if needed in the future.

## Future Improvements

1. Implement persistent storage (e.g., database integration)
2. Add more complex financial operations (e.g., scheduled transactions, interest calculations)
3. Implement a caching layer for frequently accessed data
4. Expand test coverage
5. Set up CI/CD pipeline for automated testing and deployment
6. Implement more sophisticated error handling and validation
7. Add metrics and monitoring

## Running the Project

1. Clone the repository
2. Navigate to the project directory
3. Run `go mod download` to download dependencies
4. Run `go run cmd/server/main.go` to start the server

The server will start on `localhost:8080`.

## API Endpoints

- `POST /account`: Create a new account
- `GET /account/{id}`: Retrieve account information
- `POST /transaction`: Process a new transaction
- `GET /transactions`: Retrieve transaction history

## License

This project is licensed under the MIT License.