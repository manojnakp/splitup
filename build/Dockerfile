# syntax=docker/dockerfile:1

FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY tools tools
COPY internal internal
COPY cmd cmd
COPY pkg pkg

RUN go build -o /server ./cmd

EXPOSE 8080

CMD [ "/server" ]
