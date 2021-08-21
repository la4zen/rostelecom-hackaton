FROM golang:latest

RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go mod download
RUN go build ./cmd/app/main.go

CMD ["/app/main"]