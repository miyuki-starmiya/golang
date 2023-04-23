
## Commands

- build: コンパイル
- fmt: フォーマット
- get: 外部パッケージの取得
- mod: モジュールメンテナンス
  - init: go.modの初期化
- run: コンパイルした後、実行
- test: テスト実行

## 開発の仕方

1. ディレクトリの初期化: go mod init <url>
2. ライブラリの読み込み: go get modules
3. 実装: write code into main.go


## Packages

モジュールのこと。`main`から最初に読み込まれる。mainがないと動かない
Goは`大文字`の変数しか外部から参照できない

- import: 複数の場合は()で囲む

```go
package main

import (
  "fmt"
  "math"
)

func main() {
  ...
}
```

## Variables

大文字ではじまる変数は外部から参照可能。小文字の変数はそのパッケージからしか参照できない。
静的に型をつけなくても、自動で型推論される。型宣言なしで宣言・初期化(代入)する場合`:=`を使う。
変数に何も代入しない時はゼロ値が暗黙的に代入される

### Type

- Basic
  - int
  - float32
  - complex64
  - bool
- String
- Array: [capacity]T
- Slice: variable Array []T
- Struct
- Pointer: address of value
- Interface: 抽象基底メソッド
- Map: key-value
- Function: 
- error: StandardErrorクラス

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

### Print Debug

- Print: フォーマットなし。改行なし
- Println: Print line. フォーマットなし。改行あり
- Printf: Print format. フォーマットあり。改行なし

## Types

- primitive
  - bool
  - string
  - int
  - byte(uint8)
  - float32, float64
  - complex64, complex128
- advanced
  - pointer
  - struct

### Type Cast

型を関数のようにすることで、型キャストができる

```go
i := 3
var j float32 = float32(i)
```

### Array

- Array: [n]T 固定長
- Slice: []T 可変長。Arrayの参照なので、参照を変更すると実体も変更される(参照渡し)。ArrayよりSliceの方がよく使われる

- methods
  - len(): 要素数
  - cap(): capacity. アドレス数
  - make([]T, len): 要素数分のゼロ値を格納したスライスを作成
  - append(slice, ...factors): Sliceに要素を追加

- range: iterableをiterationさせる時に使う

removeする時は癖があり、i番目の前後の要素を再作成する必要がある

```go
func main() {
  // initialize
	var arr = [3]int{1, 2, 3}
	arr2 := make([]int, 2)
	arr2[0] = 1
	arr2[1] = 2
	fmt.Println(arr, arr2)

  // append
  add := append(arr, 4)

  // concat
  c := append(arr, arr2...)

  // remove
  i := 1
  r := append(arr[:i], arr[i+1:]...)

  // range
  arr := []int{1, 2, 3}
	for i, e := range arr {
		fmt.Println(i, e)
	}
}
```

### Map

GoのMapはiterableのコールバックとして使うのではなく、ObjectやHash、Dictの役割

```go
	m := make(map[string]int)
	m["key"] = 0
	fmt.Println(m)

	for k, v := range m {
		fmt.Println(k, v)
	}
```

### Pointer

データの実体ではなく、メモリアドレスを返す

- &: 実体 => ポインタ
- *: ポインタ => 実体

```go
func main() {
  i := 1
  var p *int = &i
  fmt.Println(p) // 0xc000014260
  var data = *p
  fmt.Println(data) // 1
}
```

## Control Flow
### For Loop

初期値と条件式を囲む()がいらない

- break: ループを脱出
- continue: 次のイテレーションへ進む

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

#### range

`for i in iterable`の構文

```go
func main() {
  iterable := []int{1, 2, 3}
  for i, e := range iterable {
    fmt.Println(i, e)
  }

  for _, e := range iterable {
    fmt.Println(e) // discard _
  }
}
```

### If Conditional

三項演算子のショートハンド記法はない

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

// equivalent

func main() {
  switch {
  case height < 170:
    fmt.Println("not good")
  case height < 180:
    fmt.Println("great")
  default:
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

すべてのcase文にbreakが暗黙的に実行されている

```go
func main() {
	fmt.Println("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin": // 暗黙的にbreakが実行
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("it's default")
	}
}
```

### Error Handling

Goにはtry, catchのような構文が存在しない。代わりにerrorクラスを条件分岐させて出力させる



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
### Defer

関数がreturnされた後に実行する

```go
func main() {
  defer fmt.Println("world")

  fmt.Println("hello ")
}
```

deferした関数はLIFO(Last-in-First-out)になる

```go
func main() {
  for i := 0; i<3; i++ {
    defer fmt.Println(i)
  }
}
// 2, 1, 0
```

## Type

### Pointer

データの実体ではなくメモリアドレスを返す。ポインタはアドレスだから軽量
意味論: 関数の引数に渡してメモリを省力化

- &: 実体 to アドレス
- *: アドレス to 実体。ポインタ型や、ポインタからデリファレンスして実体を得る

- デリファレンス: ポインタから実体を得ること

```go
func main() {
  i := 1
  var p *int = &i // iのアドレス
  fmt.Println(p) // 0xc000014260
  var data = *p
  fmt.Println(data) // 1
}
```

### Struct

struct.fieldでアクセスする。pointer.fieldでもアクセス可

```go
type Human struct {
	age  int
	name string
}

func main() {
	h := Human{
    age: 14,
    name: "hitoe"
  }
  p := &h
	fmt.Println(Human{14, "hitoe"})
	fmt.Println(h.Age)
  fmt.Println(p.Age) // (*h).Age

  no_name := Human{Age: 6} // no_name.Name = ""
}
```

### Method

Structに関数をアタッチする

- レシーバ: アタッチされるStruct
  - 値レシーバ: レシーバは変更されない
  - ポインタレシーバ: メソッド内の変更がレシーバに反映されるため、セッターに用いられる

```go
// ポインタレシーバ
func (p *Person) greet(newName string) {
	p.name = newName
}

func main() {
	p := Person{
		age:  16,
		name: "hitoe",
	}

	fmt.Println(p)
	fmt.Println(p.name)
	p.greet("hoo")
	fmt.Println(p.name)
```

### Interface

`抽象基底メソッド`
Goには継承がない代わりにInterfaceでStructの抽象化を行う。
異なるStructに共通の`振る舞い`を持たせることができる。一方で、継承より柔軟で疎結合となる

```go
type Animal interface {
	speak() string
	run()
}

type Person struct {
	age  int
	name string
}

func (p Person) speak() string {
	return "hello"
}
func (p Person) run() {
	fmt.Println("human run")
}

type Dog struct {
	name string
}

func (d Dog) speak() string {
	return "hello"
}
func (d Dog) run() {
	fmt.Println("dog run")
}

func main() {
	var a Animal

	p := Person{
		age:  16,
		name: "hitoe",
	}
	d := Dog{name: "pochi"}
	a = p
	a.run()
	a = d
	d.run()
}
```

### Struct Embeddings

実質Structの継承

```go
type Human struct {
  age int
  name string
}

func (h Human) greet() {
  fmt.Printf("hey, I'm %s", h.name)
}

type Worker struct {
  Human // 継承的な
  occupation string
}

func main() {
  w := Worker{
    Human: Human{
      age: 24,
      name: "hitoe",
    },
    occupation: "engineer",
  }

  w.greet() // "hey, I'm hitoe"
}
```

### Generics

## Concurrency
### Goroutine
### Channel


