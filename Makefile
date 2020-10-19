all: test build

test:
	go test github.com/hekonsek/dkr github.com/hekonsek/dkr/dkr

build:
	go build -o out/dkr main/*.go

docker: build
	docker build out -t hekonsek/dkr:`versioon current`

docker-release: docker
	docker push hekonsek/dkr:`versioon current`
	docker tag hekonsek/dkr:`versioon current` hekonsek/dkr:latest
	docker push hekonsek/dkr:latest

git-version:
	@git tag v`versioon current`
	git push --tags

bump:
	versioon bump

release: bump docker-release git-version

install: docker
	docker rm dkr
	docker create --name dkr hekonsek/dkr:`grep 'version =' main/version.go | cut -d '"' -f 2`
	sudo docker cp dkr:/bin/dkr /usr/bin/

install-latest:
	docker rm dkr
	docker create --name dkr hekonsek/dkr
	sudo docker cp dkr:/bin/dkr /usr/bin/

images:
	docker build --target docker images -t hekonsek/dkr-docker
	docker push hekonsek/dkr-docker
	docker build --target go images -t hekonsek/dkr-go
	docker push hekonsek/dkr-go
	docker build --target docker-go images -t hekonsek/dkr-docker-go
	docker push hekonsek/dkr-docker-go