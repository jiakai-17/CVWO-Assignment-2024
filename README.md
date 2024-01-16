# CS Gossip

- A simple web forum
- Demo: <https://gossip.kirara.dev/>

## Quick start

1. Ensure `docker` and `docker compose` are installed.
   1. Refer to the Docker Docs on guides to install [Docker](https://docs.docker.com/get-docker/) and [Docker Compose](https://docs.docker.com/compose/install/) for your system.
2. Clone the repository.
   1. `git clone https://github.com/jiakai-17/CVWO-Assignment-2024`
   2. `cd CVWO-Assignment-2024`
3. Edit the environment variables in the `docker-compose.yml` configuration file.
   1. `POSTGRES_PASSWORD: YOUR_PASSWORD`: Edit `YOUR_PASSWORD` to your desired password to access the database.
   2. `POSTGRES_DB: YOUR_DB`: Edit `YOUR_DB` to your desired database name.
   3. `DATABASE_URL: "YOUR_CONNECTION_STRING"`: Edit `YOUR_CONNECTION_STRING` to your desired connection string. The connection string should be in the format `"host=localhost user=postgres dbname=YOUR_DB password=YOUR_PASSWORD port=5432"`.
   4. `JWT_SECRETSTRING: "YOUR_JWT_SECRET_STRING"`: Edit `YOUR_JWT_SECRET_STRING` to your desired secret string used to sign JWT tokens.
4. Build and start the application
    1. `docker-compose up -d --build`
5. Access the application at `http://localhost:3000`.

## Environment variables

### Backend

|Name|Description|Default value|Required|Example|
|---|---|---|---|---|
|`DATABASE_URL`|The connection string used to connect to the database.|None|Yes|`"host=localhost user=postgres dbname=YOUR_DB password=YOUR_PASSWORD port=5432"`|
|`JWT_SECRETSTRING`|The secret used to sign JWT tokens.|`secretstring`|No|`"YOUR_JWT_SECRET_STRING"`|
|`BASE_PATH`|The base path of the API.|`/api/v1`|No|`"/api/v1"`|

### Database

|Name|Description|Default value|Required|Example|
|---|---|---|---|---|
|`POSTGRES_PASSWORD`|The password used to access the database.|None|Yes|`"YOUR_PASSWORD"`|
|`POSTGRES_USER`|The username used to access the database.|`postgres`|No|`"YOUR_USERNAME"`|
|`POSTGRES_DB`|The name of the database to be created and used for the application.|Same as `POSTGRES_USER`|No|`"YOUR_DB"`|

## Development

Refer to the respective README files in the `frontend` and `backend` directories for development instructions.

## Acknowledgements

This project was developed as part of the [CVWO Winter Assignment 2024](https://github.com/CVWO/assignment-instructions/tree/cvwo-2324)

The frontend uses [React](https://react.dev/), [Material UI](https://mui.com/material-ui/), [Tailwind CSS](https://tailwindcss.com/) and [Vite](https://vitejs.dev/).

The backend is written in [Go](https://golang.org/), with [gorilla/mux](https://github.com/gorilla/mux) for routing and [sqlc](https://sqlc.dev/) for database access.

The database is [PostgreSQL](https://www.postgresql.org/).

The application is containerized using [Docker](https://www.docker.com/).
