dev:
	cls
	go run main.go

test-nc:
	cls
	go clean -testcache
	go test ./test/api/... -p 1

test-v:
	cls
	go clean -testcache
	go test ./test/api/... -v -p 1

test-prod:
	go clean -testcache
	go test ./test/api/... -p 1

build:
	go build -o ./dist/app

start:
	dist/app/backend.exe