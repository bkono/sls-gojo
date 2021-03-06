.PHONY: deps clean build gen

build:
	# dep ensure
	env GOOS=linux go build -ldflags="-s -w" -o bin/web ./web/cmd/lambda/...
	go build -o bin/local-web ./web/cmd/local/...

gen:
	go generate ./web/...

deploy: gen build
	sls deploy -v

