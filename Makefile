sqlc:
	@sqlc generate
rest:
	@go run main.go r
migrate:
	@go run main.go m
hooks:
	@go run main.go git:pre-commit