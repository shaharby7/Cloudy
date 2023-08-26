# syntax=docker/dockerfile:1
FROM golang:1.19-alpine AS build
WORKDIR /app
COPY ./go.mod ./
RUN go mod download
COPY ./ ./
RUN go build -o /program ./cmd/fakeprovider/fakeprovider.go

#######
FROM build
WORKDIR /
COPY --from=build /program /program
ENTRYPOINT ["/program"]