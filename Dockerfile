FROM golang:latest

WORKDIR /go/src/app

COPY . .
RUN go build -o main src/ClubInit.go

CMD ["bash"]