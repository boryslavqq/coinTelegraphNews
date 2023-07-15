FROM golang:1.17

WORKDIR /dockerapp

COPY . .

RUN go mod download

RUN go build -o /main


CMD ["/main"]