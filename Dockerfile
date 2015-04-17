# FROM debian:jessie
FROM dorajistyle/golang-mysql-base
MAINTAINER JoongSeob Vito Kim <dorajissanai@nate.com>

ENV USER_PATH /go/src/github.com/dorajistyle
ENV PROJECT_PATH $USER_PATH/goyangi

WORKDIR /go/src
RUN mkdir -p $USER_PATH
WORKDIR $USER_PATH
RUN mkdir $PROJECT_PATH
ADD . $PROJECT_PATH
WORKDIR $PROJECT_PATH
ENV GOPATH $PROJECT_PATH/.vendor:$GOPATH
ENV PATH $PROJECT_PATH/.vendor/bin:$PATH
RUN goop exec go run goyangi.go init
WORKDIR $PROJECT_PATH/frontend/canjs/compiler
RUN gulp
WORKDIR $PROJECT_PATH/migrate
# RUN goop install
RUN /usr/bin/mysqld_safe & \
sleep 10s && \
mysql -uroot -e 'create database goyangi_dev' && \
goop exec go run migrate.go

WORKDIR $PROJECT_PATH
#RUN /usr/bin/mysqld_safe & \
#    sleep 10s && \
#    go run server.go
EXPOSE 3001
ENTRYPOINT ["./run_server.sh"]
