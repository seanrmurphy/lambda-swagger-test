GOBUILD=go build
FUNCTION=lambda-swagger-test

build:
	GOOS=linux GOARCH=amd64 ${GOBUILD}
	chmod +x ${CURDIR}/${FUNCTION}
	zip -j ${CURDIR}/${FUNCTION}.zip ${CURDIR}/${FUNCTION}

.PHONY: build
