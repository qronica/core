FROM golang:1.19-alpine

WORKDIR /app

#Copy and run files for dependencies
COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

ENV CGO_ENABLED 0

RUN go build -o qronica .

EXPOSE 8090

CMD ["./qronica", "serve", "--http", "0.0.0.0:8090"]