FROM golang:1.19.4-alpine3.17 AS builder

WORKDIR /app

COPY go.mod /app
COPY go.sum /app

RUN go mod download

COPY main.go /app

RUN CGO_ENABLED=0 go build -o /app/ksc-employee-mock

FROM alpine:3.17

WORKDIR /app
COPY --from=builder /app/ksc-employee-mock /app/ksc-employee-mock

EXPOSE 8080
ENTRYPOINT [ "/app/ksc-employee-mock" ]