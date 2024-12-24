# Project Setup

## Docker Containers
So far, we only have a relational postgres db
```bash
docker-compose up -d 
```
## Migraciones

### Creación de Migraciones
we are gonna be using the golang-migrate module to handle script migration in out relation postgres db:

If we wanna create a new migration script, we can follow this example:
```bash
migrate create -seq -ext sql -dir <dir> <file_name>
```

#### Ejemplo
```bash
migrate create -seq -ext sql -dir ./cmd/migrate/migrations create_users
```

### Executing db script migration process
To execute this migration file, we can run the following command (example):
```bash
migrate -path=./cmd/migrate/migrations -database="postgres://postgres:debtspassword@localhost/debts?sslmode=disable" up
```

### Swagger API Rest documentation
We are gonna be using gin-swagger, we can install it, executing:
```bash
go get -u github.com/swaggo/gin-swagger
```
After adding all the required annotation, we can execute the following command to see our swagger updated:
```bash
swag init
```
