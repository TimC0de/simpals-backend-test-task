FROM golang:1.21.4-alpine

ARG PKG

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /worker ${PKG}/cmd

CMD [ "/worker" ]
