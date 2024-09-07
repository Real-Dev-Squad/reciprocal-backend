# Reciprocal Backend
This is the backend for the Reciprocal project. It is a Go application built using the Standard Library.

## Setup local development environment
1. Install [Go 1.23.1](https://golang.org/dl/)
  - if you have a different version of Go installed, visit the [Managing Go installations](https://go.dev/doc/manage-install) page to learn how to manage multiple versions of Go on your machine.

2. Install [Docker desktop](https://www.docker.com/products/docker-desktop/), make sure you have [compose V2.*](https://docs.docker.com/compose/migrate/), available when you install Docker Desktop.

3. Install Packages
```bash
go mod download
```

4. Run the setup script to create .env file and setup database, make sure you have docker running and `make` installed.
```bash
make setup
```

5. Run the application
```bash
make watch
```

6. You can now access the application at `http://localhost:4000`