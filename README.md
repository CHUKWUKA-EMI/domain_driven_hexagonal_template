# Domain-Driven Design using Hexagonal Architecture

You are free to clone and use this repository as a starter template for your next project.

To start using this template, make sure you have Go installed on your local machine, then follow these steps:

1. clone the repository by running the command:

   > git clone `https://github.com/CHUKWUKA-EMI/domain_driven_hexagonal_template.git`

2. Install project dependencies:

   > go mod download

3. Run project:

   > go run cmd/main.go

4. Edit the project as you want

## Technologies

- Go (Programming language) `v1.24.1`
- Docker

## Project Structure

```
├── Dockerfile
├── README.md
├── cloudbuild.yaml
├── cmd
│   └── main.go
├── go.mod
├── go.sum
└── internal
    ├── adapter
    │   ├── http
    │   │   ├── handler
    │   │   │   └── handler.go
    │   │   ├── middleware
    │   │   │   └── middleware.go
    │   │   └── route
    │   │       └── route.go
    │   └── persistence
    │       └── mongo_db_repo.go
    ├── application
    │   ├── user_dto.go
    │   └── user_service.go
    ├── domain
    │   ├── repository.go
    │   ├── service.go
    │   ├── user.go
    │   ├── user_onbaording_step.go
    │   └── user_role.go
    └── infrastructure
        ├── auth
        │   └── firebase_auth.go
        ├── config
        │   └── config.go
        ├── constants
        │   └── constants.go
        ├── db
        │   └── db.go
        ├── http
        │   └── client.go
        ├── logger
        │   └── logger.go
        └── secret_manager
            └── secret.go
```
