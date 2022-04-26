# base image
FROM golang:1.17.5-alpine3.14 as base
WORKDIR /builder

ENV GO111MODULE=on CGO_ENABLED=0

COPY go.mod /builder/
RUN go mod download

COPY . .
RUN go build -o /builder/main /builder/

# runner image
FROM gcr.io/distroless/static:latest
WORKDIR /app
COPY --from=base /builder/main main

EXPOSE 8080
CMD ["/app/main", "-port", "8080", "-timeout", "5"]