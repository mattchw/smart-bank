createdb:
	docker exec -it postgres13 createdb --username=postgres --no-password --encoding=UTF8 --locale=en_US.UTF-8 --template=template0 smart_bank

dropdb:
	docker exec -it postgres13 dropdb --username=postgres --no-password --if-exists smart_bank

migrateup:
	migrate -path ./db/migration -database "postgresql://postgres:test1234@localhost:5432/smart_bank?sslmode=disable" --verbose up

migratedown:
	migrate -path ./db/migration -database "postgresql://postgres:test1234@localhost:5432/smart_bank?sslmode=disable" --verbose down

sqlc:
	sqlc generate

test:
	go test -v ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/mattchw/smart-bank/db/sqlc Store

.PHONY: createdb dropdb migrateup migratedown sqlc test server