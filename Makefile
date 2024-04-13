build:
	@go build -o bin/hotel-service.exe

#run: build
#	@./bin/api.exe

test:
	@go test -v ./...

g-docs:
	@swag init && ./docs/fix.sh