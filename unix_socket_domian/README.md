## Unix Socket Domain

### Unixドメインソケットの基本

TCPとUDPによるソケット通信は、外部のネットワークに繋がるインタフェースに接続。
Unixドメインソケットではカーネル内部で完結する高速なネットワークインタフェースを作成。

Unixドメインソケットを開くには、ファイルシステムのパスを指定。 
サーバプロセスを起動すると、ファイルシステム上の指定した位置にファイルができ、クライアントは、 ファイルパスを使って通信相手を探す。 

Unixドメインソケットで作成されるのは、ソケットファイルという特殊なファイルであり、
通常のファイルのような実態はない。

### Unixドメインソケットの使い方

#### クライアント

```go
conn, err := net.Dial("unix", "socketfile")
if err != nil {
  panic(err)
}
// conn を使った読み書き
```

net.Dial()を使う
第一引数が"tcp"ではなく"unix"
第二引数がファイルのパス

#### サーバー

```go
listener, err := net.Listen("unix", "socketfile")
if err != nil {
  panic(err)
}

defer listener.Close()

conn, err := listener.Accept()
if err != nil {
  // エラー処理
}
// conn を使った読み書き
```

TCPと同様にnet.Listen()を使う
第一引数が"tcp"ではなく"unix"
第二引数がファイルのパス

### UnixドメインソケットとTCPのベンチマーク

```
Unixドメインソケットの場合はほぼ、カーネル内のバッファにデータをコピーして、
そこからサーバプロセス（の裏側のカーネル内）のバッファに書き込む程度の負荷しかかかりません。
```

// TCP localhostとUnix Domain Socketはどちらが速いのか？
https://qiita.com/ma2shita/items/154ad8f55d75051234c6

TCP overheadがない分unix domain socketのほうが早い