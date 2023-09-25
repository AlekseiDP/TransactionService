postgres:
	docker run --name postgres -e POSTGRES_PASSWORD=Qwerty123! -e POSTGRES_USER=pgsu  -p 5433:5432 -d postgres

createdb:
	docker exec -it postgres createdb --username=pgsu --owner=pgsu TransactionService

dropdb:
	docker exec -it postgres dropdb TransactionService

migrate:
	go run migrate/migrate.go

.PHONY: postgres createdb dropdb migrate