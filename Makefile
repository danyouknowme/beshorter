postgres:
	docker run --name postgresql --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it postgresql createdb --username=root --owner=root beshorter

dropdb:
	docker exec -it postgresql dropdb beshorter

migrateup:
	migrate -path database/migration -database "postgresql://root:secret@localhost:5432/beshorter?sslmode=disable" -verbose up

migratedown:
	migrate -path database/migration -database "postgresql://root:secret@localhost:5432/beshorter?sslmode=disable" -verbose down

test:
	go test -v ./...

server:
	go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown test server
