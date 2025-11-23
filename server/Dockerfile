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
RUN go install github.com/go-delve/delve/cmd/dlv@latest && \
    go install github.com/air-verse/air@latest && \
    go install github.com/rubenv/sql-migrate/sql-migrate@latest && \
    go install github.com/stephenafamo/bob/gen/bobgen-mysql@latest && \
    go install github.com/99designs/gqlgen@latest

# コンテナの公開するポートを指定
EXPOSE 5050
# 起動時のコマンド(airを使用するため)
CMD ["air"]
