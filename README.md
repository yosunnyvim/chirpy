# Chirpy

A RESTful API built with Go.

## Technologies

- [Go](https://go.dev/)
- [PostgreSQL](https://www.postgresql.org/)
- [SQLC](https://sqlc.dev/)
- [Goose](https://pressly.github.io/goose/)
- [JWT](https://jwt.io/)
- REST API
- Webhooks (Polka)

## API Endpoints

### Users

| Method | Path       | Description               |
|--------|------------|---------------------------|
| POST   | /api/users | Create a new user         |
| PUT    | /api/users | Update authenticated user |

### Authentication

| Method | Path         | Description                               |
|--------|--------------|-------------------------------------------|
| POST   | /api/login   | Login and receive access + refresh tokens |
| POST   | /api/refresh | Refresh an access token                   |
| POST   | /api/revoke  | Revoke a refresh token                    |

### Chirps

| Method | Path                  | Description           |
|--------|-----------------------|-----------------------|
| GET    | /api/chirps           | Get all chirps        |
| GET    | /api/chirps/{chirpID} | Get a single chirp    |
| POST   | /api/chirps           | Create a chirp        |
| DELETE | /api/chirps/{chirpID} | Delete your own chirp |

**Query parameters for `GET /api/chirps`:**
- `author_id` – filter by author UUID
- `sort` – `asc` or `desc` (default: `asc`)

### Webhooks

| Method | Path                | Description                 |
|--------|---------------------|-----------------------------|
| POST   | /api/polka/webhooks | Handle Polka payment events |

### Admin / Ops

| Method | Path           | Description              |
|--------|----------------|--------------------------|
| GET    | /api/healthz   | Health check             |
| GET    | /admin/metrics | Request counter          |
| POST   | /admin/reset   | Reset metrics (dev only) |
