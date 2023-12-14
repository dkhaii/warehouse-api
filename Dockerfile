FROM golang:1.21-alpine3.19

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o warehouse-api .

CMD ["./warehouse-api"]