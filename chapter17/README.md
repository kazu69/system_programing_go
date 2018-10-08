# Go言語とコンテナ

## 仮想化は低レイヤの技術の組み合わせ

ユーザモード、特権モードの下に、ハイパーバイザー用OSのモードを追加し、
それを使うことで、ゲストOSからホストOSへの処理の移譲が必要な操作を効率よくフックできる仕組み

CPUから見ると、ホストOSの下により強力な権限をもったレイヤーが追加されている

#### 準仮想化

ハードウェア上にインストールされたハイパーバイザー（ホストOSはない）上で動作する。 ゲストOSは、自分がハイパーバイザーの上で動作していることを意識している

#### 完全仮想化

ゲストOSがホストOS上にインストールされ、ゲストOSは自分が仮想環境で動いているのを意識する必要がない

## コンテナ

```
アプリケーションが好き勝手にしても全体が壊れないような、他のアプリケーションに干渉しない・されない箱を作る
```

コンテナのことを`OSレベル仮想化`と呼ぶこともある。

#### cgroupsで制限する項目

- CPU
- メモリ
- ブロックデバイス（mmap可能なストレージとほぼ同義）
- ネットワーク
- /dev以下のデバイスファイル

#### namespaceで制限する項目

- プロセスID
- ネットワーク（インタフェース、ルーティングテーブル、ソケットなど）
- マウント（ファイルシステム）
- UTS（ホスト名）
- IPC（セマフォ、MQ、共有メモリなどのプロセス間通信）
- ユーザー（UID、GID）

## libcontainerでコンテナを自作する

Dockerのコアとなっているのは、Go言語で書かれている[libcontainer](https://github.com/docker/libcontainer)

#### Dcokerのコンテナの変遷

- LXCのラッパー
- libcontainer
- runC (libcontainerを含んでいる)

## OSのブートに必要な下準備

Alpine Linuxのイメージに含まれているファイルシステムを使う

#### Prepare

```sh
$ docker pull alpine 
$ docker run --name alpine alpine
$ docker export alpine > alpine.tar
$ docker rm alpine

$ mkdir rootfs
$ tar -C rootfs -xvf alpine.tar
```

#### exec

```
$ docker build -t container .
$ docker run --privileged -it container /bin/ash

# in container 
$ ./container
/bin/sh: can't access tty; job control turned off
$ /bin/hostname

```
