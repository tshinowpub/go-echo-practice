FROM golang:1.12

MAINTAINER tshinow-pub <shino.t.music@gmail.com>

#---------------------------------------------------#
# Migration Server
#---------------------------------------------------#
RUN apt-get -y update && \
    apt-get -y upgrade && \
    apt-get install -y lsb-release

RUN curl -L https://packagecloud.io/golang-migrate/migrate/gpgkey | apt-key add -

RUN echo "deb https://packagecloud.io/golang-migrate/migrate/debian/ $(lsb_release -sc) main" > /etc/apt/sources.list.d/migrate.list

RUN apt-get -y update && \
    apt-get install -y migrate