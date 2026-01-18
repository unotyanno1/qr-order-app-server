# ベースイメージとして公式のGolangイメージを使用
# go.modの要求バージョン（1.24.0）に合わせて1.25を使用
FROM golang:1.25-alpine

# gitをインストール（go installに必要）
RUN apk add --no-cache git

# 作業ディレクトリを設定
WORKDIR /app

# go.modとgo.sumをコピー（依存関係のキャッシュを活用）
COPY go.mod go.sum ./
RUN go mod download

# airをインストール（依存関係取得後）
# リポジトリがgithub.com/air-verse/airに移行したため、新しいパスを使用
RUN go install github.com/air-verse/air@latest

# PATHにGOPATH/binを追加
ENV PATH=$PATH:/go/bin

# ソースコードをコンテナにコピー
COPY . .

# airをデフォルトのコマンドとして設定
CMD ["air", "-c", ".air.toml"]