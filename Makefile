postgres:
	docker run --name postgres-tors -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:latest

mysql:
	docker run --name mysql-tors -p 3306:3306 -e MYSQL_ROOT_PASSWORD=secret -d mysql:latest

createdb: 
	docker exec -it postgres-tors createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres-tors dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

migrate:
	migrate create -ext sql -dir db/migration -seq init_schema

.PHONEY: postgres createdb dropdb migrateup migratedown sqlc migrate