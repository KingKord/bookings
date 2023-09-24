# base go image
FROM golang:1.20.5-alpine
RUN go version
ENV GOPATH=/


COPY ./ ./


RUN go mod download
RUN go build -o bookingApp ./cmd/web

CMD [ "./bookingApp", "-dbhost=postgres",  "-dbname=postgres", "-dbuser=postgres", "-dbpass=password", "-cache=false", "-production=false"]
