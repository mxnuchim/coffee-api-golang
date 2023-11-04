# Go API with Docker, Postgres, and Chi Router

![Go](https://img.shields.io/badge/Go-1.16-blue)
![Docker](https://img.shields.io/badge/Docker-20.10.7-blue)
![Postgres](https://img.shields.io/badge/Postgres-13.3-blue)
![Chi Router](https://img.shields.io/badge/Chi%20Router-5.0.5-blue)

This is a RESTful API built with the Go programming language, Docker for containerization, Postgres for the database, and the Chi router for routing requests. It provides a solid foundation to kickstart your Go API development.

## Features

- **Dockerized Environment:** Easily deploy your application and database in Docker containers.

- **Postgres Database:** Utilize the powerful Postgres database to store and manage your data.

- **Chi Router:** Efficiently handle routing and middleware for your API using the Chi router.

## Getting Started

### Prerequisites

Before you begin, ensure you have the following prerequisites installed on your system:

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)
- [Go](https://golang.org/dl/)
- [PostgreSQL](https://www.postgresql.org/)

### Installation

1. Clone this repository:

   ```bash
   git clone https://github.com/your-username/your-repo.git
   ```

2. Build and run the Docker containers using MakeFile command:

   ```bash
   make start_container
   ```

3. Your Go API will be accessible at:

   ```bash
   http://localhost:8080/api/v1/coffees
   ```

## API Endpoints

Here are the available endpoints for the API:

- `GET /api/v1/coffees`: Get a list of all coffees.
- `GET /api/v1/coffees/coffee/{id}`: Get a specific coffee by ID.
- `POST /api/v1/coffees/new`: Create a new coffee.
- `PATCH /api/v1/coffees/coffee/{id}`: Update a specific coffee by ID.
- `DELETE /api/v1/coffees/coffee/{id}`: Delete a specific coffee by ID.

You can test these endpoints using tools like [curl](https://curl.haxx.se/) or API clients like [Postman](https://www.postman.com/) or [Insomnia](https://insomnia.rest/).

### Database

This project includes a Postgres database. You can configure the database connection in the `Makefile` file.

### Acknowledgments

- [Go](https://golang.org/)
- [Docker](https://www.docker.com/)
- [PostgreSQL](https://www.postgresql.org/)
- [Chi Router](https://github.com/go-chi/chi)
