.PHONY: run

run:
	go run cmd/main.go 

# -- Database commands:
dbstart: runcontainer create migrateup

# run docker container
# you can use any version of postgres image eg. 10.6-alpine
runcontainer:
	docker run --name TEST_DB -p 5432:5432 -e POSTGRES_USER=pgadmin -e POSTGRES_PASSWORD=Secret123 -d postgres:10.6-alpine

create:
	docker exec -it TEST_DB createdb --username=pgadmin product_test
drop:
	docker exec -it TEST_DB dropdb --username=pgadmin product_test

migrateup:
	migrate -path internal/app/db/migration -database "postgresql://pgadmin:Secret123@localhost:5432/product_test?sslmode=disable" -verbose up

migratedown:
	migrate -path internal/app/db/migration -database "postgresql://pgadmin:Secret123@localhost:5432/product_test?sslmode=disable" -verbose down

sqlc:
	sqlc generate
psql:
	docker exec -it product-postgres-1 psql -U pgadmin product
