# nats-jetStream-sample

## USGE
| アイテム | 説明 |
| --- | --- |
| Name | 名前 |
| Storage | ストレージの種類, File または Memory |
| Subjects | コンシュームするサブジェクトのリスト |
| Replicas | クラスタ化されたJetStream内のメッセージごとに保持するレプリカの数、最大5 |
| MaxAge | ストリーム内のメッセージの最大経過時間。単位はナノ秒 |
| MaxBytes | ストリーム内のメッセージの最大バイト数。単位はバイト |
| MaxMsgs | ストリーム内のメッセージの数。 |
| MaxMsgSize | ストリームが受け入れる最大のメッセージサイズ |
| MaxConsumers | 特定のストリームに対して定義できるコンシューマーの数。無制限の場合は -1 |
| NoAck | ストリームによって受信されたメッセージの確認を無効 |
| Retention | メッセージ保持の考慮方法、 LimitsPolicy(デフォルト)、InterestPolicyまたは WorkQueuePolicy |
| Discard | ストリームが制限に達すると、 DiscardNewは新しいメッセージを拒否し、DiscardOld \（default \）は古いメッセージを削除します |
| Duplicates | 重複メッセージを追跡するためのウィンドウ。ナノ秒単位で表されます。 |

### Consumer
コンシューマは、ストリームの「クライアント」と考えることができます。
コンシューマは、「サブジェクトフィルタ」と「再生ポリシー」に従って、ストリームに格納されているメッセージのすべてまたはサブセットをコンシュームし、1つまたは複数のクライアントアプリケーションで使用することができます。同じストリームを指す数千個のコンシューマを定義しても構いません。
| アイテム | 説明 |
| --- | --- |
| AckPolicy | AckExplicit, AckNone or AckAll |
| DeliverPolicy | コンシューマーを最初に作成するときに、ストリーム内のどこでメッセージの受信を開始するかを指定できます。 |
| DeliverySubject | 観察されたメッセージを配信する件名。プル型では使用できません。 |
| Durable | サーバーが追跡するコンシューマーの名前です。名前を設定することで、durableなコンシューマーにすることができます。 |
| FilterSubject | ワイルドカードサブジェクトを持つストリームから消費する場合、これにより、メッセージを受信する完全なワイルドカードサブジェクトのサブセットを選択できます。 |
| MaxAckPending | 保留しているAckの最大数、-1で際限なく保存可能 |
| MaxMsgs | ストリーム内のメッセージの数。 |
| FlowControl | フローコントロールを有効にするかどうか |
| IdleHeartbeat | アイドルハートビート期間が設定されている場合、送信する新しいメッセージがない間、サーバーは定期的にステータスメッセージをクライアントに送信します（つまり、期間が経過したとき）。これにより、ストリームにアクティビティがない場合でも、JetStreamサービスがまだ稼働中であることがクライアントに通知されます。メッセージステータスヘッダーのコードは100になります。FlowControlとは異なり、アドレスへの応答はありません。 「アイドルハートビート」のような説明があるかもしれません |
| MaxDeliver | 特定のメッセージが配信される最大回数。 ackポリシーのために再送信されるすべてのメッセージに適用されます。 |
| RateLimit | コンシューマーへのメッセージの配信をビット/秒で調整するために使用されます。 |
| ReplayPolicy | DeliverAll , DeliverByStartSequence or DeliverByStartTime |
| SampleFrequency | 可観測性のためにサンプリングする必要がある確認応答のパーセンテージを設定します。0〜100この値は文字列であり、たとえば、有効な値として30％と30％の両方を許可します。 |

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

### Subscriptionの数など
- http://localhost:8222/subsz
