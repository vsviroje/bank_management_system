connectdb:
	docker run --name postgres17 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:17.0-alpine3.20

initdb:
	docker exec -it postgres17 dropdb --if-exists bank_management_system
	docker exec -it postgres17 createdb --username=root --owner=root bank_management_system
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/bank_management_system?sslmode=disable" -verbose up

destroydb:
#	migrate -path db/migration -database "postgresql://root:password@localhost:5432/bank_management_system?sslmode=disable" -verbose down
#	docker exec -it postgres17 dropdb bank_management_system
	docker stop postgres17
	docker rm postgres17
sqlc:
	sqlc generate
test:
	go test -v -cover ./...
.PHONY:
	connectdb initdb destroydb