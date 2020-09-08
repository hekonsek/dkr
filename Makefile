all: test build

test:
	go test github.com/hekonsek/dkr github.com/hekonsek/dkr/dkr

build:
	go build -o out/dkr main/*.go

docker-build: build
	docker build out -t hekonsek/dkr

docker-push: docker-build
	docker push hekonsek/dkr

install: docker-build
	docker rm dkr
	docker create --name dkr hekonsek/dkr
	sudo docker cp dkr:/bin/dkr /usr/bin/

commands:
	docker build commands/packer -t hekonsek/dkr-packer
	docker push hekonsek/dkr-packer
	docker build commands/docker-last-id -t hekonsek/dkr-docker-last-id
	docker push hekonsek/dkr-docker-last-id

images:
	docker build --target docker images -t hekonsek/dkr-docker
	docker push hekonsek/dkr-docker
	docker build --target go images -t hekonsek/dkr-go
	docker push hekonsek/dkr-go
	docker build --target docker-go images -t hekonsek/dkr-docker-go
	docker push hekonsek/dkr-docker-go