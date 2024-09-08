dev:
	go run main.go
doc:
	swag init -o ./docs

gen:
	@protoc --go_out . --go-grpc_out . proto/*.proto