FROM golang:latest

WORKDIR /go/src
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN apt update

# RUN go build ./cmd/app/main.go

# CMD [ "tail", "-f", "/dev/null" ]
CMD [ "go", "run", "./cmd/app/main.go" ]