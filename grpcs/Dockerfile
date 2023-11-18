FROM golang:1.21.4-alpine

ARG PKG

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /grpc_service ${PKG}/cmd

ARG GRPC_SERVICE_ACCESS_PORT
EXPOSE ${GRPC_SERVICE_ACCESS_PORT}

CMD [ "/grpc_service" ]
