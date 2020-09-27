FROM golang

ENV SRC_DIR=/go/src/sushita/

ENV GOBIN=/go/bin

WORKDIR $GOBIN

ADD . $SRC_DIR

RUN cd /go/src/;

# mysql のドライバ
RUN go get github.com/go-sql-driver/mysql;

RUN go get github.com/gorilla/mux;

# RUN go install github.com/mohohewo/;

ENTRYPOINT ["./sushita"]

EXPOSE 8080