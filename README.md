# Reciprocal Backend
This is the backend for the Reciprocal project. It is a Go application built using the Standard Library.

## Setup local development environment
1. Install [Go 1.23.1](https://golang.org/dl/)
  - if you have a different version of Go installed, visit the [Managing Go installations](https://go.dev/doc/manage-install) page to learn how to manage multiple versions of Go on your machine.

2. Install [Docker desktop](https://www.docker.com/products/docker-desktop/), make sure you have [compose V2.*](https://docs.docker.com/compose/migrate/), available when you install Docker Desktop.

3. Install [air](https://github.com/air-verse/air?tab=readme-ov-file#installation) to enable hot reloading of the application.

4. Install Packages
```bash
go mod download
```

5. Run the setup script to create .env file and setup database, make sure you have docker running and `make` installed.
```bash
make setup
```

6. Create the tables used to keep track of migrations
```bash
make migration-init
```
- if you do not have make installed, you can run the following command:
```bash
go run cmd/bun/main.go db init
```

7. Run the application
```bash
make watch
```

7. You can now access the application at `http://localhost:4000`

## Migrations
### Create the tables used to keep track of migrations
1. To create the tables used to keep track of migrations, run the following command:
```bash
make migration-init
```
#### If you do not have make installed, you can run the following command:
```bash
go run cmd/bun/main.go db init
```

### Create a new migration
1. To create a new migration, run the following command:
```bash
make create-migration
```
2. You will be prompted to enter the name of the migration.
```bash
Enter migration name: <migration_name>
```
#### If you do not have make installed, you can run the following command:
```bash
go run cmd/bun/main.go db create_sql <migration_name>
```

### Run migrations
1. To run migrations, run the following command:
```bash
make migrate
```
#### If you do not have make installed, you can run the following command:
```bash
go run cmd/bun/main.go db migrate
```

### Rollback the last migration
1. To rollback migrations, run the following command:
```bash
make rollback
```
#### If you do not have make installed, you can run the following command:
```bash
go run cmd/bun/main.go db rollback
```

### Check migration status
1. To check the status of migrations, run the following command:
```bash
make status
```
#### If you do not have make installed, you can run the following command:
```bash
go run cmd/bun/main.go db status
```
