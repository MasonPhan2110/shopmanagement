postgres:
	docker run --name postgres14 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14-alpine
createdb:
	docker exec -it postgres14  createdb --username=root --owner=root shop_management
dropdb:
	docker exec -it postgres14 dropdb shop_management
migrateup:
	migrate -path db/migration -database "postgresql://root:secret@127.0.0.1:5432/shop_management?sslmode=disable"  -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/shop_management?sslmode=disable"  -verbose down
sqlc:
	sqlc generate
.PHONY: postgres createdb dropdb migrateup migratedown sqlc