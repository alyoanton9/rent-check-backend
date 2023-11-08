FROM golang:1.21-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN go build -v -o /rent-checklist-app cmd/main.go

EXPOSE 8080

CMD ["/rent-checklist-app"]
