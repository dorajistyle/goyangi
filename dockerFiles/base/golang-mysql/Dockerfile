# FROM ubuntu:14.04
FROM debian:jessie
# FROM debian:wheezy
MAINTAINER JoongSeob Vito Kim <dorajissanai@nate.com>

# Run upgrades and install basic packages
RUN echo "deb http://ftp.kr.debian.org/debian/ jessie main contrib non-free" >> /etc/apt/sources.list && echo "deb-src http://ftp.kr.debian.org/debian/ jessie main contrib non-free" >> /etc/apt/sources.list && rm -rf /var/lib/apt/lists/* && apt-get update && apt-get --no-install-recommends -qq -y --force-yes install \
# RUN apt-get update -y && apt-get -qq -y install \
automake \
build-essential \
bzr \
ca-certificates \
curl \
gcc \
gccgo \
gobject-introspection \
git \
libxml2-dev \
libc6-dev \
make \
libssl-dev \
libglib2.0-dev \
libjpeg62-turbo-dev \
libpng12-dev \
gtk-doc-tools \
mercurial \
#    nodejs \
libmysqlclient-dev \
# nginx \
# supervisor \
uuid-runtime \
openssh-server \
openssh-client \
vim
#    && rm -rf /var/lib/apt/lists/*

ENV LC_ALL C.UTF-8

ENV GOLANG_VERSION 1.4.2

RUN curl -sSL https://golang.org/dl/go$GOLANG_VERSION.src.tar.gz \
| tar -v -C /usr/src -xz

RUN cd /usr/src/go/src && ./make.bash --no-clean 2>&1

ENV PATH /usr/src/go/bin:$PATH

RUN mkdir -p /go/src
ENV GOPATH /go
ENV PATH /go/bin:$PATH
RUN go get github.com/nitrous-io/goop

# WORKDIR /go

#COPY go-wrapper /usr/local/bin/

# From docker-library / golang

# From Google Golang
#RUN mkdir /goroot && curl go$GOLANG_VERSION.linux-amd64.tar.gz | tar xvzf - -C /goroot --strip-components=1
# RUN mkdir /gopath

# ENV GOROOT /goroot
# ENV GOPATH /gopath
# ENV PATH $PATH:$GOROOT/bin:$GOPATH/bin

# ENV NODE_VERSION v0.11.13
ENV NODE_VERSION v0.10.37
RUN git clone https://github.com/creationix/nvm.git && \
echo 'source /nvm/nvm.sh && nvm install ${NODE_VERSION}' | bash -l
ENV PATH /nvm/${NODE_VERSION}/bin:${PATH}
RUN npm install -g bower && \
npm install -g gulp
RUN \
DEBIAN_FRONTEND=noninteractive apt-get install -y mysql-server && \
rm -rf /var/lib/apt/lists/* && \
sed -i 's/^\(bind-address\s.*\)/# \1/' /etc/mysql/my.cnf && \
sed -i 's/^\(log_error\s.*\)/# \1/' /etc/mysql/my.cnf && \
echo "mysqld_safe &" > /tmp/config && \
echo "mysqladmin --silent --wait=30 ping || exit 1" >> /tmp/config && \
echo "mysql -e 'GRANT ALL PRIVILEGES ON *.* TO \"root\"@\"%\" WITH GRANT OPTION;'" >> /tmp/config && \
bash /tmp/config && \
rm -f /tmp/config

RUN git clone https://github.com/jcupitt/libvips.git
WORKDIR /libvips
# USER docker
RUN ./bootstrap.sh && \
./configure --enable-debug=no --without-python --without-fftw --without-libexif \
--without-libgf --without-little-cms --without-orc --without-pango --prefix=/usr && \
make && \
make install && \
ldconfig
