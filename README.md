# starter-go

A small Go starter project for an HTTP API with a simple storage abstraction.

The app starts an HTTP server, wires it to an in-memory user store, and exposes a few example routes. It is intentionally minimal, so it can be used as a base for experimenting with API handlers, storage implementations, middleware, and tests.

## Requirements

- Go 1.25.4 or newer

## Getting Started

Run the server:

```sh
go run .
```

By default, the server listens on `:3000`.

To choose a different address:

```sh
go run . -listenadrr=:8080
```

## API Routes

- `GET /user` returns a sample user from the configured store.
- `GET /user/id` returns a sample user and demonstrates calling a utility helper.
- `/foo`, `/bar`, and `/baz` are placeholder routes for additional handlers.

Example response from `/user`:

```json
{ "id": 1, "name": "Foo" }
```

## Project Structure

```text
.
├── api/       HTTP server, routes, handlers, middleware, and auth notes
├── storage/   Storage interface plus memory and example Mongo implementations
├── types/     Shared domain types
├── util/      Utility helpers
├── main.go    Application entrypoint
└── go.mod     Go module definition
```

## Testing

Run the test suite:

```sh
go test ./...`
```
