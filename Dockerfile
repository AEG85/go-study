FROM golang:1.26.1-alpine

RUN apk add --no-cache make

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o /app/exe main.go

CMD [ "/app/exe" ]