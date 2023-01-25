
## Packages

モジュールのこと。`main`から最初に読み込まれる。mainがないと動かない

- import: 複数の場合は()で囲む

```go
package go

import (
  "fmt"
  "math"
)

func main() {
  ...
}
```

## Types

- bool
- string
- int
- byte(uint8)
- float32, float64
- complex64, complex128

### Type Cast

型を関数のようにすることで、型キャストができる

```go
i := 3
var j float32 = float32(i)
```

## Variables

大文字ではじまる変数は外部から参照可能。小文字の変数はそのパッケージからしか参照できない。
静的に型をつけなくても、自動で型推論される。型宣言なしで代入する場合:=を使う。
変数に何も代入しない時はゼロ値が暗黙的に代入される

- ゼロ値
  - int: 0
  - bool: false
  - string: ""

```go
var c, python bool // false false
var i int // 0
var j int = 3 // 3
const Str = "string"
func foo() {
  x := "str" // 型宣言なしで代入。関数内でしか使えない
}
```

## Control Flow

### For Loop

初期値と条件式を囲む()がいらない

```go
func main() {
  for i := 0; i<3; i++ {
    fmt.Println(i)
  }
}
```

`while`は存在せず、以下のように書ける

```go
func main() {
  cnt := 0 // カウンタ変数
  for cnt<3 {
    fmt.Println(cnt)
    cnt++
  }
}
```

`loop`は存在せず、以下のように書ける

```go
func main() {
  for {
    // loop infinity
  }
}
```

### If Conditional

```go
func main() {
  if height < 170 {
    fmt.Println("not good")
  } else if height < 180 {
    fmt.Println("great")
  } else {
    fmt.Println("excellent!")
  }
}
```

変数宣言付きショートハンド

```go
func main() {
  if height := 182; height > 180 {
    fmt.Println("excellent!")
  }
}
```

#### Switch



## Functions

```go
func add(x int, y int) int {
  return x + y
}

// equivalent
func add(x, y int) int {
  return x + y
}
```

- naked return: 戻り値を命名する

```go
func split(sum int) (x, y int) {
  x = sum * 2
  y = sum + 3
  return // x, yを省略している
}
```

## Commands

- build: コンパイル
- fmt: フォーマット
- get: 外部パッケージの取得
- mod: モジュールメンテナンス
  - init: go.modの初期化
- run: コンパイルした後、実行
- test: テスト実行

