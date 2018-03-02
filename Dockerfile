FROM library/golang:alpine

RUN apk add --no-cache inotify-tools git
RUN go get golang.org/x/tools/cmd/godoc
