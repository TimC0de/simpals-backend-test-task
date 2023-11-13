FROM golang:1.21.4-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /api_service

ARG API_ACCESS_PORT
EXPOSE ${API_ACCESS_PORT}

CMD [ "/api_service" ]
