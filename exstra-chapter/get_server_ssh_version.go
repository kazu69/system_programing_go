package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"golang.org/x/crypto/ssh"
)

var host = "<HOSTNAME>"

// 取得してきたサーバーの鍵情報
var hostKeyString string = "<PATH/TO/SERVER/SSH/KEY>"

func main() {
	// 秘密鍵の準備
	key, err := ioutil.ReadFile("<PATH/TO/PRIVATE/SSH/KEY>")
	if err != nil {
		panic(err)
	}

	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		panic(err)
	}

	// サーバーの鍵の準備
	hostKey, _, _, _, err := ssh.ParseAuthorizedKey([]byte(hostKeyString))
	if err != nil {
		panic(err)
	}

	if hostKey == nil {
		log.Fatalf("no hostkey for %s", host)
	}

	// 接続設定
	config := &ssh.ClientConfig{
		User: "<USER>",
		Auth: []ssh.AuthMethod{
			// Use the PublicKeys method for remote authentication.
			ssh.PublicKeys(signer),
		},
		Timeout:         5 * time.Second,
		HostKeyCallback: ssh.FixedHostKey(hostKey),
	}

	// 通信開始
	conn, err := ssh.Dial("tcp", host+":22", config)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	session, err := conn.NewSession()
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// コマンドを実行して出力結果を取得
	output, err := session.CombinedOutput("ssh -v")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(output))
}
