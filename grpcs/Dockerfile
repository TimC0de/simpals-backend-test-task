FROM golang:1.21.4-alpine

RUN mkdir /app
COPY . /app
WORKDIR /app
CMD [ "/usr/local/go/bin/go", "run", "." ]
