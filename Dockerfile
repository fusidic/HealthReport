FROM golang:alpine as builder

RUN apk update && apk upgrade && \
    apk add --no-cache git

RUN mkdir /app
WORKDIR /app

ENV GO111MODULE=on GOPROXY=https://goproxy.cn

COPY . .

RUN go mod tidy
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o HealthReport main.go mail.go parse.go request.go client.go date.go file.go initrequest.go

# Run container
FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /info
WORKDIR /
COPY --from=builder /app/HealthReport .
COPY ./report.sh .

CMD ["./report.sh"]