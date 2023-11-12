FROM golang:1.21.4-alpine

RUN mkdir /app
COPY . /app
WORKDIR /app
EXPOSE 8080
CMD [ "/usr/local/go/bin/go", "run", "." ]
