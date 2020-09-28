FROM golang

WORKDIR /go/src/github.com/mox692/sushita_serve

ADD . $APP_PATH

ENV MYSQL_HOST="mysql" \
    MYSQL_PORT="3306" \
    MYSQL_USER="golang-test-user" \
    MYSQL_PASSWORD="golang-test-pass" \
    MYSQL_DATABASE="golang-test-database"

RUN go get -u github.com/go-sql-driver/mysql

CMD ["go","run","./cmd/main.go"]

EXPOSE 8080