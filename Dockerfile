FROM golang:latest

WORKDIR /
COPY main.go /

RUN go mod init main
RUN go get github.com/gorilla/mux

EXPOSE 8085

CMD ["go","run","main.go"]