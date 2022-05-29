ROOT_DIR:=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

run:
	air

tidy:
	go mod tidy -v

publish:
	./publish.sh

test:
	go test -tags=test_all -cover ./...

test.debug:
	go test -tags=test_all -v -cover ./...

test.cov:
	go test -tags=test_all -v -race -covermode=atomic -coverprofile=coverage.out ./...

test.it:
	go test -tags=integration -v -cover ./...

test.unit:
	go test -tags=unit -v -cover ./...

test.clean:
	go clean -testcache

cov.htm:
	go tool cover -html=coverage.out

cov.func:
	go tool cover -func=coverage.out

docker.build:
	docker build -t covid-19-api .

docker.run:
	docker run --rm -it -p 9090:9090 covid-19-api

mock:
	cd mocks && docker run -d --rm -v $(ROOT_DIR)/mocks/apib:/blueprints -p 3500:3000 burakince/drakov --watch -f '/blueprints/*.apib'

gosec:
	gosec ./...

doc:
	godoc

depend.update:
	go get -u

depend.patch:
	go get -u=patch

gen:
	go generate ./...
