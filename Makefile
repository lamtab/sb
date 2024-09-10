postgres:
	docker run --name postgres -p 5431:5432 -e POSTGRES_PASSWORD=root -d postgres

createdb:
	docker exec -it postgres createdb --username=postgres --owner=postgres sb

dropdb:
	docker exec -it postgres dropdb sb

migrateup:
	migrate -path db/migration -database "postgresql://postgres:root@localhost:5431/sb?sslmode=disable" -verbose up  

migratedown:
	migrate -path db/migration -database "postgresql://postgres:root@localhost:5431/sb?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown sqlc