FROM golang:1.19.0-bullseye
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go mod download
RUN go build -o service .
CMD ["/app/service"]
