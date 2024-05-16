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
