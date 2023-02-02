postgres:
	docker run --name booking-app -p 2345:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15-alpine

createdb:
	docker exec -it booking-app createdb --username=root --owner=root booking_app

dropdb:
	docker exec -it booking-app dropdb booking_app

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:2345/booking_app?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:2345/booking_app?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server