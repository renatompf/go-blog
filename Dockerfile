FROM golang:1.21.5-alpine AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /go-blog

EXPOSE 8080

CMD [ "/go-blog" ]
