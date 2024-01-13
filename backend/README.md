# Backend

## Environment variables

- `DATABASE_URL`: **[Required]** The URL of the database to connect to. Example: 
  `"host=localhost user=postgres dbname=DATABASE password=PASSWORD port=5432"`
- `JWT_SECRETSTRING`: The secret string used to sign JWT tokens. Defaults to `secretstring`.
- `BASE_PATH`: The base path of the API. Defaults to `/api/v1`.

## Running the server (development)

1. Ensure that you have a PostgreSQL database running.
2. Ensure that you have provided the necessary environment variables (e.g: `DATABASE_URL`).
3. Run the following command:
    ```bash
    go run backend/cmd/server
    ```
4. The server should now be running on port `9090`.
