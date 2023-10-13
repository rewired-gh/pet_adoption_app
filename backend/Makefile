CONTAINER_NAME=mysql
DB_NAME=pet_adoption

setcontainer:
	docker run --name $(CONTAINER_NAME) -p 3306:3306 -e MYSQL_ROOT_PASSWORD=mysql -d mysql:8

createdb:
	docker exec -it $(CONTAINER_NAME) mysql -uroot -pmysql -e "create database $(DB_NAME);"

dropdb:
	docker exec -it $(CONTAINER_NAME) mysql -uroot -pmysql -e "drop database $(DB_NAME);"

migrateup:
	migrate --path db/migration --database "mysql://root:mysql@tcp(localhost:3306)/$(DB_NAME)?parseTime=true" --verbose up

migratedown:
	migrate --path db/migration --database "mysql://root:mysql@tcp(localhost:3306)/$(DB_NAME)?parseTime=true" --verbose down

sqlc:
	sqlc generate

server:
	go run main.go

PHONY: setcontainer createdb dropdb migrateup migratedown sqlc