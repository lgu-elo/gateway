.SILENT: daemon lint test swag build

daemon:
	APP_ENV=development CompileDaemon \
		--build="go build -o ./bin/gateway ./cmd/gateway/main.go" \
		--command="./bin/gateway" \
		-graceful-kill \
		-log-prefix=false \
		-polling \
		-polling-interval=350

build:
	GOARCH=amd64 \
	GOOS=linux \
	CGO_ENABLED=0 \
	APP_ENV=development \
	go build -o ./bin/gateway ./cmd/gateway/main.go

lint:
	golangci-lint run -c ./.golangcilint.yaml

test:
	mkdir -p ./tmp
	go clean -testcache
	go test -v ./... -coverprofile=./tmp/coverage.out

swag:
	swag fmt
	swag init -g ./internal/server/server.go -parseInternal true