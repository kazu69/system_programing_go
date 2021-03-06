## gorutine

### チャネル

queue FIFOのデータ構造。golangのチャネルはqueueに並列アクセスされても正しく処理できる機能を組み合わせたもの。

- データを順序よく受け渡すためのデータ構造
  - 配列と違いランダムアクセスできない
- 並列処理されても正しくデータを受け渡す同期機構
  - 同時に複数のgorutineがデータ読み書きを行っても一つしかアクセスできない
- 読み込み、書き込みで準備ができるまでブロックする
  - チャネルにデータがない状態で取得しに行くと、データが投入されてアクセスできるまでブロックして待つ
  - バッファに空きがない状態で書き込みすると、空きができるまでブロックする

データ入出力、終了通知とタイムアウト管理には `context.Context` を使う

```golang
// バッファなし
tasks := make(chan string)
// バッファあり
tasks := make(chan string, 10)
```

```golang
// データ送信
tasks <- "cmake ..."

// データ受け取り
task := <-tasks
// データ受け取り & close判定
taks, ok := <- tasks
// データ読み捨て
<-wait
```

gorutineはcloseされているのか受信側に確実に知らせる方法がない。
closeされていると`0`がかえるため、数値を送信する場合は注意が必要。(closeされているのかわからない)
受け取り側ではerrorチェックも無視しないで実装する必要がある。
`終了情報のやり取りは別のチャネルを利用するのが良い`
チャネルはcloseされなくてもGCされる。

終了通知するチャネルは確実にcloseしたほうがいい。データのやり取りのチャネルは`0`が紛れこむ可能性がある場合はcloseしないほうがいい。

チャネルと状態

| 操作 | バッファなし make(chan) | バッファあり make(chan, buf) | 閉じたチャネル close(chan)
--- | --- | --- | ---
chan <- val で受信 | 受け取り側が受信操作するまで停止 | バッファがあれば即座に停止なければ左に同じ | panic
var := <- chan で受信 | 送信側がデータを入れるまで停止 | 送信側がデータを入れるまで停止 ｜ 値が残っていればそれを、なければデフォルト値を返す
var, ok := <- chan で受信 | 上記に同じで、okにtrueが入る | 上記に同じで、okにtrueが入る | 上記に同じで、okにfalseが入る
for var := <- chan で受信 | チャネルに値が入るたびにループが回る | チャネルに値が入るたびにループが回る | ループから抜ける

### チャネルとセレクト文

ブロックする複数のチャネルを並列に待ち受けるさいには`select`文を使う。
select節はトリガーすると終わるため、for文の中で使うことが多い。
いずれかのチャネルが応答するまでblockし続ける。

```golang
for {
  select {
    case data := <-reader:
    // 読み込んだデータを利用
    case <-exit
    //ループを抜ける
    break
  }
}
```

読み込むまでポーリングする場合はdefault節を使う

```golang
for {
  select {
    case data := <-reader:
    // ...
    default:
    // ...
    break
  }
}
```

### コンテキスト

コンテキストは深いネスト、派生ジョブなどがあるなど複雑なロジックの中でもチャネルを正しく終了、タイムアウトを実装で切り仕組み。

### golangの通知インターフェイス

- データを順序よく受け渡すためのデータ構造
- 並列処理されも正しくデータを受け渡す同期機構
- 読み込み、書き込みで準備するまでブロックする機構
