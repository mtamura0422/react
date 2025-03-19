ARG GO_VERSION=1.24-rc-alpine

FROM golang:${GO_VERSION}

# 作業ディレクトリを設定
WORKDIR /usr/src/application


# Go Modulesの使用を許可（必要な場合）
# ENV GO111MODULE=on

# ローカルのモジュールキャッシュを最適化（必要な場合）
# COPY go.mod .
# COPY go.sum .
# RUN go mod download

RUN go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest
RUN go install  github.com/amacneil/dbmate@latest 
RUN go install  github.com/google/wire/cmd/wire@latest 
RUN go install  go.uber.org/mock/mockgen@latest


RUN apk add make

# アプリケーションをビルド
#RUN go build -o myapp

# 実行可能ファイルをデフォルトのコマンドとして設定
#CMD ["/bin/sh", "-c", "go build -o myapp && while true; do sleep 1; done"]
CMD [ "/bin/sh", "-c", "while true; do sleep 1; done" ]
