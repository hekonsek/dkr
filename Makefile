all: test build

test:
	go test github.com/hekonsek/dkr

build:
	go build -o out/dkr main/*.go
	go build -o out/dkr-proxy proxy/proxy.go

docker-build: build
	docker build out -t hekonsek/dkr

docker-push: docker-build
	docker push hekonsek/dkr

install:
	docker rm dkr
	docker create --name dkr hekonsek/dkr
	sudo docker cp dkr:/bin/dkr /usr/bin/
	sudo docker cp dkr:/bin/dkr-proxy /usr/bin/

commands:
	docker build commands/packer -t hekonsek/dkr-packer
	docker push hekonsek/dkr-packer