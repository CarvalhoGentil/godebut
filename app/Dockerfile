FROM golang:latest AS builder

RUN mkdir /app
COPY main.go /app
COPY /environ/environ.go /usr/local/go/src/environ/
WORKDIR /app

RUN go mod init main
RUN go get github.com/gorilla/mux
RUN go get github.com/jinzhu/gorm
RUN go get github.com/jinzhu/gorm/dialects/postgres
RUN go get ./...
RUN CGO_ENABLED=0 GOOS=linux go build -o main .


FROM alpine:latest AS production
COPY --from=builder /app .
COPY /pages ./pages
EXPOSE 8085

CMD ["./main"]