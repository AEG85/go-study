FROM golang:1.26.1-alpine

RUN apk add --no-cache make

WORKDIR /app

COPY . .

RUN go mod tidy

CMD [ "make", "run-http-app" ]