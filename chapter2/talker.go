package main

import (
	"fmt"
)

type Talker interface {
	Talk()
}

type Greeter struct {
	name string
}

func (g Greeter) Talk() {
	fmt.Printf("Hello my name is %s\n", g.name)
}

func main() {
	// golang インターフェイス
	var talker Talker
	talker = &Greeter{"Wozozo"}
	talker.Talk() // => Hello my name is Wozozo
}
