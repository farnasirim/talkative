FROM busybox:ubuntu-14.04
EXPOSE 80

MAINTAINER Mohammad Nasirifar "far.nasiri.m@cafebazaar.ir"

WORKDIR /

ADD talk talk

ENTRYPOINT ["/talk"]
