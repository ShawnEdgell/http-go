# Go Hello World HTTP Server

A simple "Hello, World!" HTTP server written in Go. This project demonstrates:

- Basic HTTP server setup using the standard `net/http` package.
- A `/hello` route and a root `/` route.
- Structured JSON logging using the standard `log/slog` package.

## Prerequisites

- [Go](https://go.dev/dl/) (version 1.21+ recommended for `slog`) installed.

## Running the Server

1.  Clone or navigate to the project directory:

    ```bash
    # Example:
    # git clone <your-repo-url>
    # cd go-hello-server
    ```

2.  Run the server:
    ```bash
    go run main.go
    ```
    The server will start on `http://localhost:8080`. You will see JSON logs in your terminal.

## Endpoints

- `GET /`: Displays a welcome message.
  - Example: `curl http://localhost:8080/`
- `GET /hello`: Responds with "Hello, World from Go!".
  - Example: `curl http://localhost:8080/hello`

## Project Structure

- `main.go`: Contains all the server logic, including request handlers, middleware, and server setup.
- `go.mod`: Manages project dependencies.

---
