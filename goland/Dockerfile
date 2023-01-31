FROM golang:1.17.7-alpine
WORKDIR /go/ad_serving_program
ENV DSN test_user:password@tcp(db:3306)/test_database?charset=utf8&parseTime=true
RUN apk update && apk add git && apk add mysql-client
RUN apk add alpine-sdk cmake clang
RUN /usr/local/go/bin/go install bitbucket.org/liamstask/goose/cmd/goose@latest