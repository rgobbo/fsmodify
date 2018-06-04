FROM ubuntu:16.04

COPY fsmodify /usr/app/
RUN mkdir /usr/app/test
WORKDIR /usr/app

CMD ["/usr/app/fsmodify"]