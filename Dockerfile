FROM golang

# ソースをコピーする
# ADD . /go
ENV $APP_PATH = /go/src/github.com/mox692/sushita_serve
WORKDIR $APP_PATH

ADD . $APP_PATH

# # プログラム実行
# CMD ["go","run","/go/sushita_serve/cmd/main.go"]


# 8080 ポートを解放
EXPOSE 8080