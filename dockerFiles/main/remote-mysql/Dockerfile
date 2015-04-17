# FROM ubuntu:14.04
FROM dorajistyle/golang-mysql-base
# FROM debian:wheezy
MAINTAINER JoongSeob Vito Kim <dorajissanai@nate.com>

ENV USER_PATH /go/src/github.com/dorajistyle
ENV PROJECT_PATH $USER_PATH/go-gin-canjs-i18n

RUN mkdir -p /root/.ssh && str="-----BEGIN RSA PRIVATE KEY-----\n Your SSH Key-----END RSA PRIVATE KEY-----" && echo | sed "i$str"  > /root/.ssh/id_rsa && \
echo "ssh-rsa publickey" > /root/.ssh/id_rsa.pub && \
chmod 600 /root/.ssh/id_rsa && \
printf "Host github.com\n\tStrictHostKeyChecking no\n" >> /root/.ssh/config

WORKDIR /go/src
RUN mkdir -p $USER_PATH
WORKDIR $USER_PATH
RUN git clone git@github.com:dorajistyle/goyangi.git
WORKDIR $PROJECT_PATH
ENV GOPATH $PROJECT_PATH/.vendor:$GOPATH
ENV PATH $PROJECT_PATH/.vendor/bin:$PATH
RUN goop install
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
