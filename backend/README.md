# CS Gossip - Backend

## Development

1. `go mod download`
2. Ensure that you have a PostgreSQL database running.
3. Supply the necessary environment variables (e.g: `DATABASE_URL`, `JWT_SECRETSTRING`).
4. `go run backend`
5. Access the server at `http://127.0.0.1:9090`.

## Environment Variables

- `DATABASE_URL`: **[Required]** The URL of the database to connect to. Example:
  `"host=localhost user=postgres dbname=DATABASE password=PASSWORD port=5432"`
- `JWT_SECRETSTRING`: The secret string used to sign JWT tokens. Defaults to `secretstring`.
- `BASE_PATH`: The base path of the API. Defaults to `/api/v1`.

## API Documentation

API documentation is available on SwaggerHub at <https://app.swaggerhub.com/apis-docs/jk17/CS-Gossip-Backend-API/1.0>

## Directory Structure

```
.
├───cmd
│   └───server           // Starts the server
├───docs                 // Swagger documentation
├───internal
│   ├───database         // Handles database access
│   ├───handlers
│   │   ├───comments     // Handle comment-related requests (CRUD)
│   │   ├───threads      // Handle thread-related requests (CRUD, searching, etc)
│   │   └───user         // Handle user-related requests (login, register, etc)
│   ├───models           // Models for Threads, Comments and Users
│   ├───router           // Handles routing to the correct handler
│   └───utils            // Utility functions (e.g: JWT signing, password hashing, etc)
└───sql                  // SQL queries for sqlc
```
