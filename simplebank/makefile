postgres:
	sudo docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=tesfay2f -e POSTGRES_PASSWORD=tsionawi@2121 -d postgres:12
createdb:
	sudo docker exec -it postgres12 createdb --username=tesfay2f --owner=tesfay2f simple_bank
dropdb:
	sudo docker exec -it postgres12 dropdb --username=tesfay2f simple_bank
migrateup:
	migrate -path db/migrations -database 'postgresql://tesfay2f:tsionawi@2121@localhost:5432/simple_bank?sslmode=disable' -verbose up
migratedown:
	migrate -path db/migrations -database 'postgresql://tesfay2f:tsionawi@2121@localhost:5432/simple_bank?sslmode=disable' -verbose down
backtov1:
	smigrate -path db/migrations -database 'postgresql://tesfay2f:tsionawi@2121@localhost:5432/simple_bank?sslmode=disable' force 1
sqlc:
	sqlc generate
getpid:
	sudo lsof -i :5432
test:
	go test -v -cover ./...
startdocker:
	sudo docker start postgres12
.PHONY: createdb dropdb migrateup migratedown getpid sqlc test startdocker