FROM golang:alpine

RUN apk update && apk add --no-cache git

WORKDIR /app

RUN go install github.com/codegangsta/gin@latest

# CMD ["go", "run", "main.go"]

# COPY . .

# RUN go get .

# RUN go mod tidy

# RUN go mod vendor


