.PHONY: migrate index processway target sanity

migrate:
	go run ./cmd/database/main.go

index:
	go run ./cmd/index/main.go

processway:
	go run ./cmd/processway/main.go

target:
	go run ./cmd/target/main.go

sanity:
	go run ./cmd/sanity/main.go

