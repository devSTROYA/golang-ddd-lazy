# Go Clean Architecture (Echo + DDD + Wire)

A small backend service that demonstrates:

- Echo as HTTP transport layer
- DDD-style domain modeling
- clear separation between domain, application, presentation, and infrastructure
- dependency injection using Google Wire

This project currently implements user registration/authentication and todo management with in-memory repositories.

## Tech Stack

- Go `1.25`
- [Echo v4](https://echo.labstack.com/) for HTTP server and routing
- [Google Wire](https://github.com/google/wire) for compile-time dependency injection
- [JWT](https://github.com/golang-jwt/jwt) for auth
- In-memory local repositories (no external database yet)

## Architecture Overview

The codebase follows a layered approach:

- `domain`: core business rules, value objects, entities, domain errors, repository interfaces
- `application`: use-cases (commands and queries) orchestrating domain logic
- `presentation`: HTTP controllers, routers, request/response DTO mapping
- `infrastructure`: concrete implementations (env config, local repositories, local unit of work)
- `common`: shared concerns (error response types, Echo filter/middleware helpers)

Dependency direction is inward:

- `presentation -> application -> domain`
- `infrastructure -> domain` (implements repository interfaces)

`main` composes all dependencies through Wire and starts Echo.

## Project Structure

```text
.
├── main.go
├── wire.go / wire_gen.go
├── factory.go
├── domain/
│   ├── user/
│   └── todo/
├── application/
│   ├── user/
│   │   ├── commands/register_user/
│   │   └── queries/get_current_user/
│   └── todo/
│       ├── commands/add_todo/
│       ├── commands/complete_todo/
│       └── queries/get_todos_user/
├── presentation/http/
│   ├── user/
│   └── todo/
├── infrastructure/
│   ├── config/
│   └── local/
└── common/
    ├── platform-echo/
    └── types/
```

## Dependency Injection (Google Wire)

- `wire.go` defines providers and injector graph for `InitializeApp()`
- `wire_gen.go` is generated code used at runtime

Regenerate Wire output when constructor/provider signatures change:

```bash
go generate ./...
```

You can also run:

```bash
go run github.com/google/wire/cmd/wire
```

## Environment Configuration

Configuration is loaded by `infrastructure/config/env.go`.

Supported `GO_ENV` values:

- `LOCAL` -> `.env.local`
- `DEV` (default) -> `.env.development.local`
- `PROD` -> `.env.production.local`
- any other value -> fallback `.env`

Important environment variables:

- `PORT` (example: `8080`; app listens as `:PORT`)
- `JWT_SECRET` (required for signing/validating tokens)
- `JWT_EXPIRATION_TIME` in hours (required for practical token expiry behavior)
- `DATABASE_URL` (currently optional and not used by local repositories)

### Minimal example (`.env.development.local`)

```env
PORT=8080
JWT_SECRET=super-secret-key
JWT_EXPIRATION_TIME=24
```

## Run the App

Install dependencies first:

```bash
go mod tidy
```

Run without hot reload:

```bash
go run .
```

Server starts on `http://localhost:8080` by default.

## Development Mode (Air Hot Reload)

This project can be run with [Air](https://github.com/air-verse/air) for live reload during development.

Install Air (if needed):

```bash
go install github.com/air-verse/air@latest
```

Run with Air:

```bash
air
```

Notes:

- Air watches file changes and recompiles/restarts the app automatically.
- This repository does not currently include a `.air.toml`; Air will use its default config unless you add one.

## API Endpoints

### Health

- `GET /` -> `{"message":"Hello World!"}`

### Auth

- `POST /auth/register`
  - body:
    ```json
    {
      "name": "John Doe",
      "email": "john@example.com",
      "password": "your-password"
    }
    ```
  - response `201`:
    ```json
    {
      "accessToken": "<jwt-token>"
    }
    ```

- `GET /auth/info` (requires `Authorization: Bearer <token>`)

### Todos (all routes require auth)

- `GET /todos`
- `POST /todos`
  - body:
    ```json
    {
      "title": "Finish clean architecture writeup",
      "description": "Add README and examples"
    }
    ```
- `PATCH /todos/:todoId`

## Example cURL Flow

Register:

```bash
curl -X POST http://localhost:8080/auth/register \
  -H "Content-Type: application/json" \
  -d '{"name":"John Doe","email":"john@example.com","password":"strong-password"}'
```

Get current user:

```bash
curl http://localhost:8080/auth/info \
  -H "Authorization: Bearer <accessToken>"
```

Add todo:

```bash
curl -X POST http://localhost:8080/todos \
  -H "Authorization: Bearer <accessToken>" \
  -H "Content-Type: application/json" \
  -d '{"title":"Learn DDD","description":"Focus on boundaries"}'
```

List todos:

```bash
curl http://localhost:8080/todos \
  -H "Authorization: Bearer <accessToken>"
```

Complete todo:

```bash
curl -X PATCH http://localhost:8080/todos/<todoId> \
  -H "Authorization: Bearer <accessToken>"
```

## Error Handling

Global HTTP errors are normalized by `common/platform-echo/filters/http_filter.go`.

Error payload format:

```json
{
  "traceId": "<request-id>",
  "error": {
    "code": "SOME_ERROR_CODE",
    "details": []
  }
}
```

Common domain-driven error codes include:

- `INVALID_EMAIL_FORMAT`
- `USER_NAME_TOO_SHORT`
- `PASSWORD_TOO_SHORT`
- `EMAIL_ALREADY_IN_USE`
- `TODO_TITLE_TOO_SHORT`
- `TODO_ALREADY_COMPLETED`
- `USER_DOES_NOT_EXIST`
- `TODO_DOES_NOT_EXIST`

## Current Limitations

- Storage is in-memory (`infrastructure/local/*`), so data is reset on restart
- No persistence/database adapter wired yet
- No test files are currently present in this repository

## Next Improvements

- Add repository implementations for a real database (PostgreSQL/MySQL)
- Add unit tests for domain/application layers
- Add integration tests for HTTP endpoints
- Add API docs (OpenAPI/Swagger)

