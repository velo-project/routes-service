# Routes Service

This is a Go project that provides a service for managing tracks and routes.

## Getting Started

To get started with this project, you'll need to have Go installed on your machine. You'll also need to have a running instance of PostgreSQL.

### Prerequisites

- Go
- PostgreSQL
- Docker (optional)

### Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/theMyntt/velo-routes-service.git
   ```
2. Install the dependencies:
   ```sh
   go mod tidy
   ```
3. Set up the environment variables:
   ```sh
   cp .env.example .env
   ```
   The following environment variables are used:
   - `POSTGRES_CONNECTION_STRING`: The connection string for the PostgreSQL database.
   - `RSA_PUBLIC_KEY`: The path to the RSA public key for verifying JWT tokens.

4. Run the database migrations:
   ```sh
   make migrate-up
   ```
5. Run the server:
   ```sh
   go run cmd/api/main.go
   ```

### Docker

To build and run the project using Docker, you can use the following commands:

1. Build the Docker image:
   ```sh
   docker build -t routes-service .
   ```
2. Run the Docker container:
   ```sh
   docker run --env-file .env -p 8080:8080 routes-service
   ```

## Folder Structure

The project is organized into the following folders:

- `cmd/api`: Contains the main application entry point.
- `db/migrations`: Contains the database migrations.
- `internal`: Contains the core application logic.
  - `adapters`: Contains the adapters for connecting to external services.
    - `database`: Contains the adapters for modifying PostgreSQL Data
    - `grpc`: Contains the adapters for call procedures in other APIs
    - `http`: Contains the rest controllers
  - `core`: Contains the core domain logic.
    - `domain`: Contains the domain models.
    - `ports`: Contains the interfaces for the repositories and services.
    - `services`: Contains the application services.
- `plan`: Contains the project plan.
- `proto`: Contains the protobuf files for the gRPC services.

## API Endpoints

The following API endpoints are available:

- `POST /tracks`: Creates a new track.
- `GET /tracks`: Finds all routes for a given user.
- `DELETE /tracks/:trackId`: Deletes a track.

## Dependencies

The project uses the following dependencies:

- [gin-gonic/gin](https://github.com/gin-gonic/gin)
- [golang-jwt/jwt/v5](https://github.com/golang-jwt/jwt)
- [google.golang.org/grpc](https://grpc.io/)
- [google.golang.org/protobuf](https://developers.google.com/protocol-buffers)
- [lib/pq](https://github.com/lib/pq)
- [joho/godotenv](https://github.com/joho/godotenv)

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.
