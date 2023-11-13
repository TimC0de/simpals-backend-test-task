FROM golang:1.21.4-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /grpc_service

CMD [ "/grpc_service" ]
