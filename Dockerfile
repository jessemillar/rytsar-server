FROM golang:1.5

RUN mkdir -p /go/src/github.com/jessemillar
ADD . /go/src/github.com/jessemillar/rytsar

WORKDIR /go/src/github.com/jessemillar/rytsar
RUN go get -d -v
RUN go install -v

CMD ["/go/bin/rytsar"]

EXPOSE 8000
