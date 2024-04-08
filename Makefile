PACKAGE=app-metrica
API=./internal/services/gateway/
SWAGGER_ARG=-g api.go -d ${API}

GIT_HASH=`git rev-parse --short HEAD`
BUILD_DATE=`date +%FT%T%z`

lint:
	golangci-lint run

swag:
	swag init --parseDependency ${SWAGGER_ARG}
	swag fmt ${SWAGGER_ARG}