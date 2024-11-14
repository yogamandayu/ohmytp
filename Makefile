rest:
	@go run main.go r
migrate:
	@go run main.go dbm
db:
	@go run main.go dbg
hooks:
	@go run main.go hooks
worker-notification:
	@go run main.go wkn