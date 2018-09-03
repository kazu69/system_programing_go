## Systemcall

http://ascii.jp/elem/000/001/267/1267477/

#### 特権モード

CPUの機能が基本的にすべて使える。
メモリ資源がなくなくなりそうなときのOOMキラーなど。

#### ユーザーモード

機能をCPUレベルでは利用できないようになっている。
システムコールで特権モードの機能を利用している。

システムコールを使わなればデータの外部読み取りや画面出力、共有メモリに書き出すことなどもできない。

### Goにおけるシステムコールの実装

- sysall.Opne()
- sysall.Read()
- sysall.Write()
- sysall.Close()
- sysall.Seek()

#### macの場合 (syscall.Open)

- https://github.com/golang/go/blob/master/src/os/file_unix.go#L190
- https://github.com/golang/go/blob/master/src/syscall/zsyscall_darwin_amd64.go#L874
- https://github.com/golang/go/blob/master/src/syscall/asm_darwin_amd64s#L15

SYSCALLLの中はgoではdebugできない。
SYSCALLの中では、`entersyscall`、`exitsyscall` が実行される。
現在実行中のOSスレッドが時間がかかるシステムコールで時間がかかる場合に`entersyscall`でマークする。そのマークをはずのが`exitsyscall`。

実行すべきタスクが現在のシステムコールでブロックされている場合に、
OSに新しいスレッドを作成するように依頼するため。
これはGo言語特有の機能、

#### linuxの場合 (syscall.Open)

- https://github.com/golang/go/blob/master/src/os/file_unix.go#L190
- https://github.com/golang/go/blob/master/src/syscall/zsyscall_linux_amd64.go#L62
- https://github.com/golang/go/blob/master/src/syscall/syscall_unix.go#L30:6
- https://github.com/golang/go/blob/master/src/syscall/asm_linux_arm.s#L44

#### windows (syscall.Open)

内部実装を公開してないため、DLLファイルをダウンロードして、wind32 APIを実行する。

- https://github.com/golang/go/blob/master/src/syscall/syscall_windows.go#L290
- https://github.com/golang/go/blob/master/src/syscall/zsyscall_windows.go#L294
- https://github.com/golang/go/blob/master/src/runtime/syscall_windows.go#L191
- https://github.com/golang/go/blob/master/src/runtime/cgocall.go#L94

### POSIXとC言語の標準規格

POSIXは5つの基本システムコール(open, read, write, close, lseak)、で構成されているOS間での共通システムコール。アプリケーションの移植性を高めるためのIEEE規格、

C言語の関数ではそれぞれ、open(), read(), write(),close(),lseak()が対応している。

Goのsyscall関数もシステムコールの呼び出し口。
直接sycall関数を使うことは基本的にない。使う場合はC言語用の情報を参照する必要がある。

### システムコールより内側の世界

システムコール関数の定義

https://github.com/torvalds/linux/blob/master/fs/read_write.c#L322

実際に呼ばれるコードは SYSCALL_DEFINEx (xの部分は0-6の数値が入る)

- https://github.com/torvalds/linux/blob/master/include/linux/syscalls.h#L446

`asmlinkage`はCPUのレジスタ経由で渡すようにするためのフラグ。
呼び出し側と呼ばれる側で環境が全く違う。(ユーザーモード領域とカーネルモード領域)
そのため引数をすべてCPUのレジスタ経由で渡すため(スタックを使わない)

### cpuで実行するまで

sys_write()はLinuxカーネルビルド時に生成されるsys_call_table配列に格納されている。配列のインデックスがシステムコールの番号になっている、

https://github.com/torvalds/linux/blob/master/arch/x86/entry/common.c#L272

配列内から`do_syscall_64`呼び出される。

https://github.com/torvalds/linux/blob/master/arch/x86/entry/common.c#L290

システムコールの番号をAXレジスタから取得して、引数として渡している。

https://github.com/torvalds/linux/blob/master/arch/x86/entry/entry_64.S

`entry_SYSCALL_64()`を呼び出している。

レジスタを構造体に対比しながらカーネル用のスタックに付け替えしながら

https://github.com/torvalds/linux/blob/master/arch/x86/entry/entry_64.S#L238

を呼び出している。

`entry_SYSCALL_64`で呼び出すのはCPUそのもの。

CPUから呼び出されるように `syscall_init()` 内で登録している。
https://github.com/torvalds/linux/blob/master/arch/x86/kernel/cpu/common.c#L1544

CPUのレジスタ`MSR_LSTAR`に登録している。このレジスタに登録された`entry_SYSCALL_64`関数を呼び出し、RAXレジスタからすsテムコール番号をもとに実際処理をする関数レジスタを処理をする関数に引数として渡す。

### Go言語のシステムコールとPOSIX

システムコールと呼び出される関数

| システムコール関数 | Linux                          | FreeBSD                      | macOS | Windows                               |
| ------------------ | ------------------------------ | ---------------------------- | ----- | ------------------------------------- |
| syscall.Open()     | openatシステムコールを呼び出す | openシステムコールを呼び出す | 同左     | Win32 APIのCreateFile()を呼び出す     |
| syscall.Read()     | readシステムコールを呼び出す   | 同左                            | 同左     | Win32 APIのReadFile()を呼び出す       |
| syscall.Write()    | writeシステムコールを呼び出す  | 同左                            | 同左     | Win32 APIのWriteFile()を呼び出す      |
| syscall.Close()    | closeシステムコールを呼び出す  | 同左                            | 同左     | Win32 APIのCloseHandle()を呼び出す    |
| syscall.Seek()     | lseekシステムコールを呼び出す  | 同左                            | 同左     | Win32 APIのSetFilePointer()を呼び出す |


同じシステムコールでもOSごとに番号が異なる。

| システム コール | Linux (x86) | Linux (x64) | Linux (arm64) | FreeBSD (x86/x64/arm) | macOS (x64/arm64) |
| --------------- | ----------- | ----------- | ------------- | --------------------- | ----------------- |
| open            | 5           | 2           |               | 5                     | 5                 |
| openat          | 295         | 257         | 56            | 499                   |                   |
| read            | 3           | 0           | 63            | 3                     | 3                 |
| write           | 4           | 1           | 64            | 4                     | 4                 |
| close           | 6           | 3           | 57            | 6                     | 6                 |
| lseek           | 19          | 8           | 62            | 478                   | 199               |

`POSIXのコンセプトは、「OS間のポータビリティを維持する」`なので、
C言語の関数をそのまま使うべきであり、システムコールを自前で番号指定して呼び出すのはナンセンス。
しかし、Go言語は自前で頑張る道を選ぶことで、他のOSで動くバイナリが簡単に作成できるようになっている。

### システムコールのモニタリング

Goのアプリはmain()関数の呼び出し前のシーケンスの中でも大量のシステムコールを呼び出している。

Linuxなら`strace`、FreeBSDは`truss`、macは`dtruss`。windowsは`API Monitor`というサードパーティのアプリ。

### エラー処理

どのシステムコールも正常の場合は`0`より大きな数値、エラーの場合は`-1`が返る。
レジスタ経由なので帰り値も、エラーも数値しか扱えない。
最も低いレイヤーでは数値だが、アプリケーションレイヤーでは言語の流儀に沿ってエラーを報告する。

Node.jsだと最初のコールバックパラメーターをエラーで返す。
Goの場合は最後の値をエラーインターフェイスとして設定する。

```golang
type writer interface {
  write(p []byte) (n int, err error)
}
```

```golang
// 上記のI/Fにより
// err != nil の場合エラーがあるということになる。
count, err := writer.Write([]byte("hello"))
if err != nil {
  log.Fatal(err)
}
```
