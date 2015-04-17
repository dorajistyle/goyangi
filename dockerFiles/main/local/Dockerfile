# FROM debian:jessie
FROM dorajistyle/golang-base
MAINTAINER JoongSeob Vito Kim <dorajissanai@nate.com>

ENV USER_PATH /go/src/github.com/dorajistyle
ENV PROJECT_PATH $USER_PATH/goyangi

WORKDIR /go/src
RUN mkdir -p $USER_PATH
WORKDIR $USER_PATH
RUN mkdir $PROJECT_PATH
ADD . $PROJECT_PATH
WORKDIR $PROJECT_PATH/frontend/canjs/compiler
RUN gulp
WORKDIR $PROJECT_PATH/migrate
# RUN goop install
# Be sure that you have a database already.
RUN goop exec go run migrate.go

WORKDIR $PROJECT_PATH

EXPOSE 3001
ENTRYPOINT ["./run_server_rds.sh"]
