# ベースイメージを指定
FROM node:22-alpine3.19

# 作業ディレクトリを指定
WORKDIR /app

# 依存関係のファイルをコピー
COPY package*.json ./

# 依存関係をインストール
RUN npm install

# 残りのアプリケーションファイルをコピー
COPY . .

# アプリケーションをビルド
RUN npm run build

# ポートを開放
EXPOSE 3000

# サーバーを起動
CMD [ "npm", "start" ]
