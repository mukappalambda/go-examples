GO ?= go

APP_NAME=$$(basename $$(pwd))

GO_LDFLAGS=-ldflags '-s -w'

clean:
	@${GO} clean
	@rm -f ${APP_NAME}

build: clean
	@go build -o ${APP_NAME} ${GO_LDFLAGS} .

run: build
	@./${APP_NAME}
