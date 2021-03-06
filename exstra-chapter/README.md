## セキュリティ関連とssh

### 乱数

#### 乱数生成アルゴリズム

Goの提供している乱数アルゴリズム

- math/rand
- crypto/rand

##### math/rand

擬似乱数生成器。シードにより乱数が変わるがシードが同じなら乱数は同じになる。
シードがわかる乱数が把握される。よって、`math/rand`をセキュリティに使ってはダメである。

##### crypto/rand

暗号のための乱数生成。外乱要因(外部由来の不安定要素)を使うことで乱数を生成する。
外乱要因と擬似乱数生成器をつかって生成する。(外乱要因だけだと観測により把握されるため)

#### 乱数のシステムコール

math/randはユーザーモードで動作する。crypto/randはOSで提供されている暗号論的擬似乱数生成器を使う。

OS | 乱数生成器
--- | ---
Linux/Windows/OpenBSD | /dev/urandom
Linux | getrandom() システムコール
Windows | CryptGenRandom()

/dev/urandom はファイルとしてアクセスできるが実態はなくカーネルが提供する擬似デバイス。
/dev/randomに乱数生成に必要なエントロピーが集められる。

getrandom()は/dev/urandomをシステムコールとして利用できるようにしたもの。
ファイルとしてアクセスしないのでファイルディスクリプたを消費しない。

#### 乱数の使い方

math/randは整数、浮動小数などの様々な方の乱数を生成する関数が揃っている。

```
Golangの連想配列(Map)はランダムに取り出される。
一見ソートされているように見えて、後々ソートされてないことに気がつかなくていいように
あらかじめランダムに取り出されるようにデザインされている
````

### TLS(Transport Layer Security)

一般的にOpenSSLを使う言語が多いが、Goでは標準ライブラリが用意されている。

#### ルート証明書

`なりすまし`でないことを証明するために公開鍵暗号基盤(PKI)を使っている。
PKIを利用するにはデジタル署名が確認出来る公開鍵の入った信用できるルート証明書が必要。

証明書に含まれる情報

- issuer 発行した認証局
- subject 発行対象(所有者)
- 公開鍵
- デジタル署名(発行した認証局の秘密鍵を使って施したデジタル署名)

連鎖的証明書の身元を確認していく。
最終的にはルート証明書に行き着く。ルート証明書はブラウザベンダーなどがあらかじめバンドルして配布している。

Golangでも極力ブラウザのルート証明書に従っている。

#### ルート証明書の取得

Goでは証明書をパースできるcrypto/x509パッケージが用意されている。
クロスコンパイル時には/usr/bin/securityをつかって証明書を取得する。

### ssh(Secure Shell)

TLSの他にサーバーに安全に接続するプロトコルとしてsshがある。
sshでは通信内容が見えないほかに、第三者が改変したり同じ命令を再送できない状態を実現できる。
Goでは標準でsshライブラリをサポート。

#### sshの基本的な流れ

ssh, TLSで使う暗号化アルゴリズムはオープン。
特定のベンダー間でのみしか通信できないということはない。
アルゴリズムが公開されていても通信を保護できるのは「鍵」を切り替えているから。

現在よく使われているのはDH鍵共有(Difiie-Hellman鍵共有)という仕組みをベースにしたもの。
鍵交換システムでもサーバー側がなりすますことがある。sshでは公開鍵方式を使ったデジタル署名によるサーバー認証を採用している。
サーバー側でクライアントの認証する際は接続できるユーザーは不特定多数ということはないのでホワイトリスト形式である。

#### Goによるssh接続

https://github.com/golang/crypto を使う。
認証し、ssh.Dialで通信開始する。
通信開始したらNewSession()でセッション開始。


- Run(cmd): 指定したプログラムを実行して終了を待つ
- Start(cmd): 指定したプログレムを実行、終了してWait()を待つ
- Output(cmd): 指定したプログラムを実行して終了を待ちつつ、リモートのs標準出力を待つ
- CombinedOutput(cmd): 指定した王ログラムを実行して終了を待ち地つつ、リモートの標準出力と標準エラー出力内容を返す

認証にパスワード使う場合

```golang
config := &ssh_ClientConfig{
  User: "root",
  Auth: []ssh.AuthMethod(
    ssh.Password("password")
  )
}
```

### keychain

漏洩すると危険になる情報、例えばユーザーIDやパスワード、クレジットカードの情報など集中管理する仕組みをキーチェーンという。

Go言語で簡単にキーチェインを扱うためのパッケージとして [keyring](https://github.com/tmc/keyring)などがある。
