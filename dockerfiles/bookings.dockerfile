# base go image
FROM golang:1.20.5-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN go mod download

# CGO=0 means that we do not use C libraries
RUN CGO_ENABLED=0 go build -o bookingApp ./cmd/web

# give brokerApp an executable flag
RUN chmod +x /app/bookingApp

FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/bookingApp /app

CMD [ "/app/bookingApp", "-dbhost=postgres",  "-dbname=postgres", "-dbuser=postgres", "-dbpass=password", "-cache=false", "-production=false"]
