MIGRATION_DIR := database/migrations

mysql-start:
	sudo service mysql start
	sudo mysql -u development -p

mysql:
	sudo mysql -u development -p

run:
	go run main.go

tidy:
	go mod tidy

build:
	go build -o bin/main main.go

migrate-create:
	@if [ -z $(NAME) ]; then echo "Usage: make create-migration NAME=<migration_name>"; exit 1; fi
	migrate create -ext sql -dir $(MIGRATION_DIR) $(NAME)

migrate-up:
	migrate -path $(MIGRATION_DIR) -database "mysql://development:development@tcp(localhost:3306)/cozy_warehouse" up

migrate-down:
	migrate -path $(MIGRATION_DIR) -database "mysql://development:development@tcp(localhost:3306)/cozy_warehouse" down

migrate-fix:
	@if [ -z $(VERSION) ]; then echo "Usage: make migrate-fix VERSION=<version>"; exit 1; fi
	migrate -path $(MIGRATION_DIR) -database "mysql://development:development@tcp(localhost:3306)/cozy_warehouse" force $(VERSION)