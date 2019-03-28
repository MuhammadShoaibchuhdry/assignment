FROM golang:alpine as builder
RUN apk --update add bash && \
    apk add dos2unix
RUN mkdir -p /go/src/assignment
ADD ./ /go/src/assignment/
WORKDIR $GOPATH/src/assignment
RUN go build -o userService user/cmd/user/main.go
EXPOSE 8070
CMD ["./start.sh"]
