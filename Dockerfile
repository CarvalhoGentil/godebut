FROM golang:latest

RUN mkdir /build
COPY ./main/ /build/main

WORKDIR /build/main

RUN export GO111MODULE=on
RUN go get github.com/gorilla/mux
RUN go mod init main
RUN go build

EXPOSE 8085

ENTRYPOINT [ "/build/main/main" ]