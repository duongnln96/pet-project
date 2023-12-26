# blog-realworld

## Clean Architecture

- Adaptper (Infrastructure) - how the application talks to the external world. Such as SQL queries, HTTP or gRPC Client, file readers and writers, Pub/Sub message publishers.

- Port - as an input to the application, the only way the external can talk to our application. Such as HTTP or gRPC Server, CLI, Pub/Sub message subcribers.

- Application logic (Use-cases) - With this layer, you can not know which database it uses or what URL it calls. It's like an orchestrator.

- In DDD (Domain-Driven Design), a domain layer that holds just the business logic.

## How to generate code

- Wiregen

```golang
// go:build wireinject
// +build wireinject
```

```bash
make wire
```

- Gen query by sqlc

```bash
make sqlc_gen
```

- Create new migration file

```bash
make new_migration name=<migration_name>
```
