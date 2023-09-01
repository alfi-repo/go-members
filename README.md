## Description

Member app

- Add a new member
- Edit existing member
- Show all members
- Delete a member

## Requirements

- Go 1.20
- MariaDB 10.7+

## Running the app

1. Import database migrations in `/db/mariadb`
2. Configure `.env` from `.env.example`
3. Run application

```bash
go run main.go
```

## Test

```bash
go test ./... -v
```

## OpenApi

openapi specification is available in `/api` directory.

## Static Analysis

- `golangci-lint` (https://github.com/golangci/golangci-lint)

## Dependency

- .env loader: `godotenv` (https://github.com/joho/godotenv)
- Env to struct parser: `env` (https://github.com/caarlos0/env)
- HTTP router: `chi` (https://github.com/go-chi/chi)
- Database client: `mysql` (https://github.com/go-sql-driver/mysql)
- Logger: `zerolog` (https://github.com/rs/zerolog)
- UUID generation: `uuid` (https://github.com/google/uuid)
