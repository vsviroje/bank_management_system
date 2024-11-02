connectdb:
	docker run --name postgres17 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:17.0-alpine3.20

initdb:
	docker exec -it postgres17 dropdb --if-exists bank_management_system
	docker exec -it postgres17 createdb --username=root --owner=root bank_management_system

migrateup:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/bank_management_system?sslmode=disable" -verbose up

sqlc:
	sqlc generate

mock:
	mockgen -package mock -destination db/mock/store.go github.com/Golang/bank_management_system/db/sqlc Store

test:
	go test -v -cover ./...

server:
	go run main.go

migratedown:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/bank_management_system?sslmode=disable" -verbose down

destroydb:
#	docker exec -it postgres17 dropdb bank_management_system
	docker stop postgres17
	docker rm postgres17

.PHONY:
	connectdb initdb migrateup sqlc mockgen test server migratedown destroydb 