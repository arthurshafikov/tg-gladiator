FROM golang:1.23.1-alpine AS build

RUN set -ex; \
    apk update; \
    apk add --no-cache git; \
    apk add build-base

ENV CODE_DIR /go/src/
ENV TEST_DIR ./tests

WORKDIR ${CODE_DIR}

COPY go.mod .
COPY go.sum .
RUN go mod download

ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.7.2/wait /wait
RUN chmod +x /wait

CMD /wait && \
    go test -v ${TEST_DIR}
