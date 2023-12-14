MIGRATION_DIR := database/migrations

db-start:
	docker run --name db-dev -e MYSQL_ROOT_PASSWORD=my-secret-pw -d mysql:tag

mysql-start:
	sudo service mysql start
	sudo mysql -u development -p

mysql:
	sudo mysql -u development -p

build:
	docker build -t be-app:v1 .

run:
	docker run -it be-app:v1

migrate-create:
	@if [ -z $(NAME) ]; then echo "Usage: make create-migration NAME=<migration_name>"; exit 1; fi
	migrate create -ext sql -dir $(MIGRATION_DIR) $(NAME)

migrate-up:
	migrate -path $(MIGRATION_DIR) -database "postgres://postgres:mysecretpassword@localhost:5432/cozy_warehouse?sslmode=disable" up

migrate-down:
	migrate -path $(MIGRATION_DIR) -database "postgres://postgres:mysecretpassword@localhost:5432/cozy_warehouse?sslmode=disable" down

migrate-fix:
	@if [ -z $(VERSION) ]; then echo "Usage: make migrate-fix VERSION=<version>"; exit 1; fi
	migrate -path $(MIGRATION_DIR) -database "postgres://postgres:mysecretpassword@localhost:5432/cozy_warehouse?sslmode=disable" force $(VERSION)