GO ?= go

APP_NAME=$$(basename $$(pwd))

GO_LDFLAGS=-ldflags '-s -w'

clean:
	@${GO} clean
	@rm -f ${APP_NAME}

build: clean
	@go build -o ${APP_NAME} ${GO_LDFLAGS} .

build-image: clean
	@docker build -t ${APP_NAME}:latest .

up-container: build-image
	@docker run -dt --name ${APP_NAME} -p 8080:8080 ${APP_NAME}:latest

down-container:
	@docker rm -f ${APP_NAME}

run: build
	@./${APP_NAME}
