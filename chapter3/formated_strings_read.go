package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

var csvSource = `
13101,"100  ","1000000","ﾄｳｷｮｳﾄ","ﾁﾖﾀﾞｸ","ｲｶﾆｹｲｻｲｶﾞﾅｲﾊﾞｱｲ","東京都","千代田区","以下に掲載がない場合",0,0,0,0,0,0
13101,"102  ","1020072","ﾄｳｷｮｳﾄ","ﾁﾖﾀﾞｸ","ｲｲﾀﾞﾊﾞｼ","東京都","千代田区","飯田橋",0,0,1,0,0,0
13101,"102  ","1020082","ﾄｳｷｮｳﾄ","ﾁﾖﾀﾞｸ","ｲﾁﾊﾞﾝﾁｮｳ","東京都","千代田区","一番町",0,0,0,0,0,0
13101,"101  ","1010032","ﾄｳｷｮｳﾄ","ﾁﾖﾀﾞｸ","ｲﾜﾓﾄﾁｮｳ","東京都","千代田区","岩本町",0,0,1,0,0,0
`

func main() {
	reader := strings.NewReader(csvSource)
	// func NewReader(r io.Reader)
	csvReader := csv.NewReader(reader)
	for {
		line, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		fmt.Println(line[2], line[6:9])
	}
}
