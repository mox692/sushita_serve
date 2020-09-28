FROM golang

# ソースをコピーする
# ADD . /go


# # プログラム実行
# CMD ["go","run","/go/sushita_serve/cmd/main.go"]


# 8080 ポートを解放
EXPOSE 8080