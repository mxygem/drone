go-get:
	go get -u ./...

go-test-no-cache:
	go test --race --count=1 ./...

go-run:
	go run ./go

go-build:
	go build -o ./bin/drone -race ./go

run:
	./bin/drone
