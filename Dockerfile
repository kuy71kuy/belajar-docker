FROM golang:1.21-alpine

COPY . /bdocker

WORKDIR /bdocker

RUN go mod tidy

RUN go build -o bdocker .

CMD ["/bdocker/bdocker"]
