# Build Stage
# First pull Golang image
FROM golang:1.17-alpine as build-env
 
WORKDIR $GOPATH/src/ums_app
 
# Copy application data into image
COPY . .

EXPOSE 3010

RUN go get

RUN apk update && apk add bash

RUN chmod +x wait-for-it.sh

CMD ["./wait-for-it.sh" , "database:3306" , "--strict" , "--timeout=300" , "--" , "go", "run", "main.go"]

