FROM golang:1.22

WORKDIR /usr/src/app

COPY app/* ./
RUN go mod download && go mod verify

CMD ["go", "run", "main.go"]
