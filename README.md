# Ecommerce GO API

## Getting Started

### API

Enter api dir:

```sh
cd api
```

Clone the existing `.env.example` into `.env` and setup environment variables:

```sh
cp .env.example .env
```

Create database container:

```sh
docker compose up -d
```

Install migration tool [migrate](https://github.com/go-migration/migrate):

```sh
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

Run migrations e.g.:

```sh
migrate -path ./db/migrations -database "postgresql://local_user:local_password@localhost:5432/local_db?sslmode=disable" up
```

Note that the database URI must match the `.env` PostgreSQL variables.
