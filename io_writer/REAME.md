### io.Writer memo

- golangではOSのAPIの差異を吸収すべく、ファイルディクリプタを模倣している
- Clangでは出力はバッファーリングするが、golangではバッファリングしない


### godocコマンド

ドキュメントを手元で開く

```sh
$ godoc -http ':6060' -analysis type
# パッケージが多いと時間がかかるので GOPATHを一時的に変更する
$ GOPATH=/ godoc -http ':6060' -analysis type
```

`-analysis type` でインターフェイスの分析

### 入出力API

| | |
--- | ---
ioutil.WriteFile() | ファイル書き込み
ioutil.ReadFile() | ファイル読み込み
http.Get() | http GET でデータ受け取り
http.Post() | http POSTでデータ送出
