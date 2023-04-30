# nats-jetStream-sample

## ローカルで試す
### natsのインストール
1. brew tap nats-io/nats-tools
2. brew install nats-io/nats-tools/nats
3. brew install nats-server

### natsサーバー起動
- `nats server run --jetstream`

### コンテキストの生成
- `nats context select nats_development`

### ストリーム作成
- `nats stream add`

### 作成されたストリーム確認
- `nats stream ls`

### 注文メッセージ発行
-  `nats pub orders.us "{{.Count}}" --count 1000`
-  `nats pub orders.eu "{{.Count}}" --count 1000`


## Docker
### network作成
- `docker network create natsnet`
### 状態確認
- http://localhost:8222/
- http://localhost:8282/
