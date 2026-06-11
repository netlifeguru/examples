# NetLife Guru Go Examples

Runnable examples for the open-source Go packages maintained by [NetLife Guru](https://netlife.guru).

This repository complements the official documentation with small, focused programs that can be read, modified, and executed locally. Each example demonstrates one feature or usage pattern and is organized by package.

## Packages

| Package                                                     | Description                                                                                           | Examples                             | Documentation                                        |
|-------------------------------------------------------------|-------------------------------------------------------------------------------------------------------|--------------------------------------|------------------------------------------------------|
| [`db`](https://github.com/netlifeguru/db)                   | Shared database layer for queries, execution, transactions, dialect-specific SQL, and result mapping. | [`./db`](./db)                       | [Documentation](https://netlife.guru/docs/go/db)     |
| [`db-mysql`](https://github.com/netlifeguru/db-mysql)       | MySQL driver for the NetLife Guru database layer.                                                     | [`./db/mysql`](./db/mysql)           | [Documentation](https://netlife.guru/docs/go/db)     |
| [`db-postgres`](https://github.com/netlifeguru/db-postgres) | PostgreSQL driver for the NetLife Guru database layer.                                                | [`./db/postgresql`](./db/postgresql) | [Documentation](https://netlife.guru/docs/go/db)     |
| [`db-scylla`](https://github.com/netlifeguru/db-scylla)     | ScyllaDB driver for the NetLife Guru database layer.                                                  | [`./db/scylla`](./db/scylla)         | [Documentation](https://netlife.guru/docs/go/db)     |
| [`form`](https://github.com/netlifeguru/form)               | Type-safe form and request validation for Go applications.                                            | [`./form`](./form)                   | [Documentation](https://netlife.guru/docs/go/form)   |
| [`logger`](https://github.com/netlifeguru/logger)           | Structured logging built on top of Go's `log/slog`.                                                   | [`./logger`](./logger)               | [Documentation](https://netlife.guru/docs/go/logger) |
| [`mapper`](https://github.com/netlifeguru/mapper)           | Lightweight database row mapper for structs, maps, and custom handlers.                               | [`./mapper`](./mapper)               | [Documentation](https://netlife.guru/docs/go/mapper) |
| [`router`](https://github.com/netlifeguru/router)           | Fast HTTP router focused on low-allocation request handling and clean application structure.          | [`./router`](./router)               | [Documentation](https://netlife.guru/docs/go/router) |

## Requirements

* Go 1.25 or newer
* Git
* A supported database for database examples
* Docker or locally installed database services for examples that require MySQL, PostgreSQL, or ScyllaDB

## Getting Started

Clone the repository:

```bash
git clone https://github.com/netlifeguru/examples.git
cd examples
```

Download the dependencies:

```bash
go mod download
```

Run an example by passing its directory to `go run`:

```bash
go run ./router/getting-started
go run ./logger/structured
go run ./form/getting-started
```

Each example is an independent `package main` program. Open its directory to inspect the complete source code and any example-specific setup instructions.

## Repository Structure

```text
.
├── db/
│   ├── databases/
│   ├── mysql/
│   ├── postgresql/
│   └── scylla/
├── form/
├── logger/
├── mapper/
├── router/
├── go.mod
└── README.md
```

The top-level directories correspond to NetLife Guru packages. Nested directories contain focused examples for individual features, APIs, and practical use cases.

## Database Examples

Database examples may require a running database and environment-specific connection settings.

Schema files are available in [`db/databases`](./db/databases):

```text
db/databases/
├── db_mysql.sql
├── db_postgresql.sql
└── db_scylla.cql
```

Review the selected example before running it and configure the required connection values. Never commit real credentials or production connection strings.

Typical commands:

```bash
go run ./db/mysql/01_getting_started
go run ./db/postgresql/01_getting_started
go run ./db/scylla/01_getting_started
```

## Explore the Examples

### DB

The database examples cover connection setup, configuration, connection pools, typed query helpers, dialect-specific SQL, insert/update/delete operations, transactions, batches, lightweight transactions, and SQL model files.

```bash
go run ./db/mysql/05_select_list
go run ./db/postgresql/30_transactions
go run ./db/scylla/30_batch
```

### Form

The form examples cover validation rules, optional values, conditional validation, HTTP request validation, custom responses, schema composition, and practical application workflows.

```bash
go run ./form/getting-started
go run ./form/practical/registration
go run ./form/httpform/httpform-response
```

### Logger

The logger examples demonstrate initialization, structured fields, chained loggers, context-aware logging, file output, terminal output, source information, and integration with HTTP servers.

```bash
go run ./logger/getting-started
go run ./logger/structured
go run ./logger/context_logging
```

### Mapper

The mapper examples demonstrate scanning database rows into structs and maps, nullable values, joins, field tags, naming conversion, custom row converters, caching, partial results, and error handling.

```bash
go run ./mapper/usage/01_getting_started
go run ./mapper/usage/03_join
go run ./mapper/edge_cases/01_null_values
```

### Router

The router examples cover route registration, handlers, middleware, groups, mounting, static files, health checks, logging, recovery, rate limiting, profiling, multiple servers, and custom error handling.

```bash
go run ./router/default
go run ./router/middleware
go run ./router/custom_error_handler
```

Most HTTP examples listen on a local port such as `:8080`. The exact address is visible in the example source.

## Build and Verify

Build every example in the repository:

```bash
go build ./...
```

Run all tests:

```bash
go test ./...
```

Run static analysis:

```bash
go vet ./...
```

Clean and verify module dependencies:

```bash
go mod tidy
git diff --exit-code -- go.mod go.sum
```

## Documentation

* [NetLife Guru](https://netlife.guru)
* [Documentation](https://netlife.guru/docs)
* [Go packages](https://github.com/netlifeguru)

The package documentation explains the APIs and recommended usage patterns. This repository provides complete runnable programs that accompany those guides.

## Contributing

Contributions, corrections, and new examples are welcome.

When adding an example:

1. Place it under the directory of the package it demonstrates.
2. Keep the example focused on one feature or use case.
3. Use a standalone `package main` program whenever possible.
4. Include comments only where they clarify an important decision.
5. Do not include secrets, production credentials, or private endpoints.
6. Make sure `go build ./...`, `go test ./...`, and `go vet ./...` succeed.

Before opening a pull request, format the code and tidy the module:

```bash
gofmt -w .
go mod tidy
go test ./...
go vet ./...
```

## Support

For documentation and project information, visit [netlife.guru](https://netlife.guru).

For questions and feedback, contact [info@netlife.guru](mailto:info@netlife.guru).

## License

This repository is licensed under the MIT License. See [`LICENSE`](./LICENSE) for details.
