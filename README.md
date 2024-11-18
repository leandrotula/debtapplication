# Project Setup

## Docker Containers

## Migraciones

### Creación de Migraciones
Usaremos la implementación de golang-migrate para manejar las migraciones:

Para crear una nueva migración, ejecuta:
```bash
migrate create -seq -ext sql -dir <dir> <file_name>
```

#### Ejemplo
```bash
migrate create -seq -ext sql -dir ./cmd/migrate/migrations create_users
```

### Ejecución del Proceso de Migración
To execute this migration file, we can run the following command (example):
```bash
migrate -path=./cmd/migrate/migrations -database="postgres://postgres:debtspassword@localhost/debts?sslmode=disable" up
```
