network:
	docker network create kruty_craft

pg:
	docker run -d \
		--name kruty_craft \
		--network kruty_craft \
		-p 5432:5432 \
  	-e POSTGRES_USER=root \
		-e POSTGRES_PASSWORD=secret \
  	postgres

pgrm:
	docker stop kruty_craft
	docker rm kruty_craft

createdb:
	docker exec -it kruty_craft createdb --username=root --owner=root kruty_craft

dropdb:
	docker exec -it kruty_craft dropdb kruty_craft

schema:
	migrate create -ext sql -dir external/database/migrations -seq init_schema

migrateup:
	migrate -path external/database/migrations -database "postgresql://root:secret@localhost:5432/kruty_craft?sslmode=disable" -verbose up

migratedown:
	migrate -path external/database/migrations -database "postgresql://root:secret@localhost:5432/kruty_craft?sslmode=disable" -verbose down

sqlc:
	sqlc generate

templ:
	templ generate

serve:
	go run cmd/main.go

