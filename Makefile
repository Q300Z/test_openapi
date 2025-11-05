run:
	go run cmd/api/main.go

build:
	make swag && go build -o bin/api cmd/api/main.go

test:
	go test -v ./...

swag:
	${HOME}/go/bin/swag init -g cmd/api/main.go -o internal/swagger/docs

generate-sdk: swag
	openapi-generator-cli generate -i internal/swagger/docs/swagger.json -g go -o sdk/go
	openapi-generator-cli generate -i internal/swagger/docs/swagger.json -g typescript-axios -o sdk/ts --additional-properties=npmRepository=https://npm.pkg.github.com/,npmName=@Q300Z/test_openapi-ts-sdk,validationAttributes=true,withInterfaces=true,axiosVersion=^1.13