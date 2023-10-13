run:
	go run main.go

tidy:
	go mod tidy

build:
	go build -o bin/main main.go

migrateup:
	migrate -database "mysql://development:development@tcp(localhost:3306)/cozy_warehouse" -path database/migrations up

migratedown:
	migrate -database "mysql://development:development@tcp(localhost:3306)/cozy_warehouse" -path database/migrations down