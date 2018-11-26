## UDP Socket

UDPはTCPと違ってコネクションレス。
一方的にデータを送りつけるのに使われます。 
パケットの到着順序も管理しない。

### UDPを使っている例

ドメイン名からIPアドレスを取得するDNS。
時計合わせのためのプロトコルのNTP。
ストリーミング動画・音声もUDPを利用するものが多い
動画・音声通信プロトコルWebRTC

独自プロトコルを開発するときにUDPが土台として選ばれることもあります

### UDPが使われる場面の変化

セキュリティ上の理由から、VPN接続でも暗号化のためにTLSを経由するSSL-VPNが使われることが増えている。

独自プロトコルを開発する場合も、土台としてUDPを使うということは、通信環境が劣悪な状態での信頼性とか、ネットワークに負荷をかけすぎて他の通信の邪魔をしないか（フェアネス）とか、そういった点について自分たちで作りこみが必要になる

大規模なフィールドテストができる状態でないならば、そもそも独自プロトコルを使わないほうが得策といえます。 TCPプロトコルの制御もバージョンがあがり、輻輳制御は高性能になっています。

現在ではアプリケーションレイヤーで使われるプロトコルの多くがTCPを土台にしている。
DNSも、512バイトを超えるレスポンスの場合にはTCPにフォールバックする仕組み


アプリケーション開発という視点で見れば、`「ロスしても良い、マルチキャストが必要、ハンドシェイクの手間すら惜しいなど、いくつかの特別な条件に合致する場合以外はTCP」`

### UDPのマルチキャストの実装例

マルチキャストでは使える宛先IPアドレスがあらかじめ決められている。
先頭4ビットが1110のアドレス（224.0.0.0 ～ 239.255.255.255）がマルチキャスト用として予約されている。

 ### UDPマルチキャストサーバー

UDPのマルチキャストでは、クライアントがソケットをオープンして待ち受け、
そこにーバがデータを送信する

### TCPとUDPの機能面の違い
 
TCPの場合は接続前のハンドシェイクに1.5RTT分の時間がかかる

- TCPには再送処理とフロー処理がある

TCPでは送信するメッセージにシーケンス番号が入っている。
受信側はメッセージを受け取ると、受信したデータのシーケンス番号とサイズの合計を確認答番号として返信。
送信側は届いたことが確認できない場合は再び送り直す。

- ウインドウ制御

受信側が用意できていない状態で送信リクエストが集中して
通信内容が失われたりするのを防ぐ。
ウインドウサイズは最初のコネクション確立時に決まる。
受信側のデータの読み込み処理が間に合わない場合には、 
受信できるウインドウサイズを受信側から送信側に伝えて通信量を制御する。

UDPには、ウインドウやシーケンス番号を利用した再送処理やフロー処理の機能はない。

### QUIC

HTTP/2では1本のTCP接続上でストリームという単位で並列化できるようになっている。
同じサーバへの接続ではTCPのフロー制御より高度な優先順位による制御方法が組み込まれている。

HTTP/2とTCPで同じようなことを重複して行っている部分を統合し、
UDPを使うことでさらに無駄を減らそうというのが`QUIC`

https://qiita.com/flano_yuki/items/251a350b4f8a31de47f5


- TCPだと1.5RTTかかっていたハンドシェイクなしに一方的な通信で済むので、通信開始までの時間が減らせる
- Wifiから3G/4G通信に切り替わったときに再接続時の遅延がなくせる
- 同一ストリーム内でのみ順序を維持することでオーバーヘッドを減らせる
- TCPのレイヤーとHTTP2のレイヤーで個別に行っていたウインドウサイズの制御が一元化
- IPのレイヤーのパフォーマンスをアプリケーションが最大限に引き出せる
