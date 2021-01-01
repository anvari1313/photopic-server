FROM golang:1.15.6 AS build

WORKDIR /src

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN make build

FROM ubuntu:20.04

WORKDIR /app

COPY --from=build /src/photopic-server /app

