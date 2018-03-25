all: build


setup:
	@go get -u golang.org/x/tools/...


build:
	@go generate .
	@go build .
