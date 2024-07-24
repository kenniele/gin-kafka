FROM golang:1.22.0-bullseye AS builder

WORKDIR /app

COPY /backend/go.mod /backend/go.sum ./
RUN go mod download

COPY /backend .

RUN go build -o main .

FROM debian:bullseye

WORKDIR /app

COPY --from=builder /app/main /app/main
COPY /frontend /app/frontend

CMD ["/app/main"]

EXPOSE 8080