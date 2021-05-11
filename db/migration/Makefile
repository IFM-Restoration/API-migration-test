postgres:
	docker run --name postgres13 -p 5432:5432 -e POSTGRES_USER=amayer -e POSTGRES_PASSWORD=ifmrestoration -d postgres:13-alpine

createdb:
	docker exec -it postgres13 createdb --username=amayer --owner=amayer api_migration_test

dropdb:
	docker exec -it postgres13 dropdb api_migration_test

migrateup: 
	migrate -path db/migration -database "postgresql://amayer:ifmrestoration@localhost:5432/api_migration_test?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://amayer:ifmrestoration@localhost:5432/api_migration_test?sslmode=disable" -verbose down

.PHONY: postgres createdb dropdb migrateup migratedown

