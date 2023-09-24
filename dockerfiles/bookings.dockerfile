FROM alpine:latest

RUN mkdir /app

COPY /dockerfiles/bookingApp /app

CMD [ "/app/bookingApp", "-dbhost=postgres",  "-dbname=bookings", "-dbuser=postgres", "-dbpass=password", "-cache=false", "-production=false"]
