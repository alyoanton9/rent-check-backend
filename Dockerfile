FROM golang:latest as build

WORKDIR /app

RUN apt-get update \
    && apt-get install -y \
    ca-certificates

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN go build -v -o /rent-checklist-app cmd/main.go

EXPOSE 80
EXPOSE 443

# volume to store TLS certificates
VOLUME ["/var/www/.cache"]

CMD ["/rent-checklist-app"]
