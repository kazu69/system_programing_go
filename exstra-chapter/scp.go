package main

import "fmt"

func main() {
	go func() {
		w, _ := session.StdinPipe()
		defer w.Close()

		content := []byte("Go言語でシステムプログラミング\n")
		fmt.Fprintln(w, "D0755", 0, "testdir") // mkdir
		fmt.Fprintln(w, "C0644", len(content), "testfile1")
		w.Write(cotennt)
		fmt.Fprintln(w, "\x00")
		fmt.Fprintln(w, "C0644", len(content), "testfile2")
		w.Write(cotennt)
		fmt.Fprint(w, "\x00")
	}()
	err = session.Run("/usr/bin/scp - tr ./")
	if err != nil {
		panic(err)
	}
}
