# gin-task-api

A simple CRUD API server for managing tasks implemented using Gin.

## Getting Started

### Prerequisites

- Go version 1.21.0 or higher is required for local development.
- Docker is required for running the application in a containerized environment.

### Running Locally

#### Non-Docker Development

To run the application locally without Docker, execute the following command:

```bash
go run main.go
```

#### Docker Development

For Docker development, follow these steps:

1. Build the Docker image:

    ```bash
    make build
    ```

2. Start the server:

    ```bash
    make run-server
    ```

    You'll see console output similar to the following:

    ```bash
    docker run --rm -p 8080:8080 --name gogolook-interview-api-server gin-server
    [GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

    [GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
    - using env:   export GIN_MODE=release
    - using code:  gin.SetMode(gin.ReleaseMode)

    [GIN-debug] GET    /tasks                    --> gin-task-api/handlers.(*TaskHandler).GetAllTasks-fm (3 handlers)
    [GIN-debug] POST   /tasks                    --> gin-task-api/handlers.(*TaskHandler).CreateTask-fm (3 handlers)
    [GIN-debug] PUT    /tasks/:id                --> gin-task-api/handlers.(*TaskHandler).UpdateTask-fm (3 handlers)
    [GIN-debug] DELETE /tasks/:id                --> gin-task-api/handlers.(*TaskHandler).DeleteTask-fm (3 handlers)
    [GIN-debug] GET    /swagger/*any             --> github.com/swaggo/gin-swagger.CustomWrapHandler.func1 (3 handlers)
    [GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
    Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
    [GIN-debug] Listening and serving HTTP on :8080
    ```

    The API server is now accessible at `http://localhost:8080`.

### Stopping the Server

To stop the server, either press Ctrl-C in the same console where it's running, or execute:

```bash
make stop-server
```

## Documentation

API documentation is generated using [gin-swagger](https://github.com/swaggo/gin-swagger). The generated swagger files are located within the `docs/` folder. You can access the documentation page at `http://localhost:8080/swagger/index.html` for local development.

If you update the API comments, regenerate the files by executing:

```bash
swag init
```

## Database

The application uses SQLite as its database with in-memory mode, meaning the data won't persist after the server is stopped.

## Test

To run the test:

```bash
go test -v ./...
```
