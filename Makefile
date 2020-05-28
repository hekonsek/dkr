all: test build

test:
	go test github.com/hekonsek/dkr

build:
	go build -o out/dkr main/*.go

docker-build: build
	docker build out -t hekonsek/dkr

docker-push: docker-build
	docker push hekonsek/dkr