FROM golang:1.25.4-alpine3.22

RUN apk update && \
    apk add --no-cache \
    git \
    build-base \
    tzdata
# ログのタイムゾーンを指定
ENV TZ /usr/share/zoneinfo/Asia/Tokyo
# コンテナ内の作業ディレクトリを指定
WORKDIR /app
# go.modとgo.sumを先にコピー（依存関係のキャッシュを効率化）
COPY ./src/go.mod ./src/go.sum* ./
# /app/go.modに記載された依存関係の解決＋必要なパッケージのダウンロードを実行
RUN go mod download
# ソースコードをコンテナ内にコピー
COPY ./src ./
COPY ./scripts ./
# Airのバイナリをインストール
RUN go install github.com/go-delve/delve/cmd/dlv@v1.24.1 && \
    go install github.com/cosmtrek/air@v1.47.0 && \
    go install github.com/stephenafamo/bob/gen/bobgen-mysql@v0.40.2 && \
    go install github.com/99designs/gqlgen@v0.17.83 && \
    CGO_CFLAGS="-D_LARGEFILE64_SOURCE" go install github.com/rubenv/sql-migrate/...@latest

# コンテナの公開するポートを指定
EXPOSE 5050
# 起動時のコマンド(airを使用するため)
CMD ["air"]
