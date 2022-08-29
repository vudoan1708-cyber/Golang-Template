# syntax=docker/dockerfile:1

# Please keep up to date with the new-version of Golang docker for builder
FROM golang:1.19 as base

# Create another stage called "dev" that is based off of our "base" stage (so we have golang available to us)
FROM base as dev

# Air for live reload if no npm is used
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

WORKDIR /app

COPY go.mod .
COPY go.sum .
COPY vendor/ ./vendor

RUN go mod download

RUN go mod vendor

RUN go mod tidy

COPY . .

RUN go build -o bin/main cmd/main.go

CMD ["air"]
