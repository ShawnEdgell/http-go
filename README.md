# Go Hello World HTTP Server (Dockerized)

A simple "Hello, World!" HTTP server written in Go, containerized with Docker, and deployed. This project demonstrates:

- Basic HTTP server setup using the standard `net/http` package.
- A `/` route and a `/hello` route.
- Structured JSON logging using the standard `log/slog` package.
- Containerization with Docker for consistent deployment.
- Serving via a reverse proxy (Caddy) with automatic HTTPS.

## Key Technologies

- Go (version 1.24.3, using `net/http` and `log/slog`)
- Docker & Docker Compose
- Caddy (as the reverse proxy on the host VPS)

## Deployed Application

The application is currently hosted on a Virtual Private Server (VPS) and is accessible at:

➡️ **[https://go.skatebit.app](https://go.skatebit.app)**

### Endpoints (Live)

- `GET https://go.skatebit.app/`
  - Displays: `Welcome to the Go HTTP Server!`
- `GET https://go.skatebit.app/hello`
  - Responds with: `Hello, World from Go!`

## Running for Local Development (Using Docker)

To run this application locally for development or testing using Docker:

### Prerequisites (Local Development)

- [Docker Desktop](https://www.docker.com/products/docker-desktop/) (for Windows/macOS) or Docker Engine (for Linux) installed.
- Git (for cloning the repository).

### Steps

1.  **Clone the repository (if you haven't already):**

    ```bash
    git clone [https://github.com/ShawnEdgell/http-go.git](https://github.com/ShawnEdgell/http-go.git) # Replace with your actual repo URL
    cd http-go
    ```

2.  **Build the Docker image:**
    From the project's root directory (where the `Dockerfile` is):

    ```bash
    docker build -t http-go-app-local .
    ```

3.  **Run the Docker container:**
    This command maps port 8080 on your local machine to port 8080 in the container.

    ```bash
    docker run -d -p 8080:8080 --name my-go-test-container http-go-app-local
    ```

    - `-d`: Run in detached mode (background).
    - `-p 8080:8080`: Map host port to container port.
    - `--name my-go-test-container`: Give the container a name.

4.  **Access locally:**
    Open your browser and go to:

    - `http://localhost:8080`
    - `http://localhost:8080/hello`

5.  **To stop and remove the local test container:**
    ```bash
    docker stop my-go-test-container
    docker rm my-go-test-container
    ```

## Deployment Overview (on VPS)

This application is deployed on a Virtual Private Server (VPS) using:

- **Docker:** The Go application is containerized using the `Dockerfile` within this project.
- **Docker Compose:** This application is run as a service defined in a centralized `docker-compose.yml` file (typically located in the `~/projects/` directory on the VPS), which manages all deployed application containers.
- **Caddy:** Acts as a reverse proxy on the VPS, routing traffic from `https://go.skatebit.app` to the running Docker container and automatically handling SSL/TLS certificates via Let's Encrypt.

## Project Structure (This Repository)

- `main.go`: Contains all the server logic, including request handlers, middleware, and server setup.
- `go.mod`: Manages project Go module definition.
- `Dockerfile`: Instructions to build the Docker image for this specific application.
- `.gitignore`: Specifies intentionally untracked files that Git should ignore.
- `README.md`: This file.

---
