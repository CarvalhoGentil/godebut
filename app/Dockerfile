FROM golang:latest AS builder

RUN mkdir /app
COPY main.go /app
WORKDIR /app

RUN go mod init main
RUN go get github.com/gorilla/mux
RUN go get github.com/jinzhu/gorm
RUN go get github.com/jinzhu/gorm/dialects/postgres
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./...


FROM alpine:latest AS production
COPY --from=builder /app .
EXPOSE 8085

# CMD ["go","run","main.go"]
CMD ["./main"]