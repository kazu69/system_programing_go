## io.Reader

### 補助関数

#### ioutil.ReadAll()

終端記号にあたるまで全てのデータを読み込み

#### io.Copy()

io.Reader()からio.Writer()にそのままデータを渡す

#### そのほか

io.Closer I/F ファイルを閉じる
io.Seeker I/F 読み書きの位置を変更
io.ReaderAt I/F 読み込み位置を変更

### 入出力インターフェイスのキャスト

io.ReadCloser を要求されているが io.Readerしか見たいしてないとき
`ioutil.NopCloser()` を使ってダミーのClose()を使う。

```golang
import (

)
var reader io.Reader = strings.NewReader("テスト")
var readCloser io.ReadCloser = ioutil.Nop(reader)
```



