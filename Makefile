BINARY_NAME=iban-number

OS_NAME := $(shell uname -s | tr A-Z a-z)

build:
ifeq ($(OS_NAME),darwin)
	GOARCH=amd64 GOOS=darwin go build -o ${BINARY_NAME} main.go
endif
ifeq ($(OS_NAME),linux)
	GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME} main.go
endif

run:
	./${BINARY_NAME}
		
clean:
	go clean

test:
	go test ./... -cover