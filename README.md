# buildingBlockchain

Building Blockchain with Go 👣

## Tutorial

### [Start a module](https://golang.org/ref/mod#go-mod-init)

## Development

`rest.go`

1. url 정의
2. router 만들기 (handler)
3. controller 만들기

## Concepts

In Go, `nil` is the **zero value for pointers, interfaces, maps, slices, channels and function types**, representing an uninitialized value.
함수명을 대문자로 시작하면 외부에서 호출 가능

- [Golang functions vs methods](https://www.sohamkamani.com/golang/functions-vs-methods/)
- [Member variable](https://cvml.tistory.com/11) : 대소문자 구분
- [Once](https://pkg.go.dev/sync#Once) : 스레드, 고루틴과 상관 없이 오지 1번만 실행
- [Singleton Pattern](https://en.wikipedia.org/wiki/Singleton_pattern) : 단 하나의 instance만을 공유
- [text/template](https://pkg.go.dev/html/template)
- [Marshal](https://jeonghwan-kim.github.io/dev/2019/01/18/go-encoding-json.html) : interface를 받아 JSON으로 변환
- [Struct field tags](https://pkg.go.dev/encoding/json#Marshal)
- [Stringer](https://pkg.go.dev/golang.org/x/tools/cmd/stringer)
- [MarshalText](https://pkg.go.dev/encoding#TextMarshaler)
- [Gorilla Web Toolkit](https://www.gorillatoolkit.org/)
  - [Gorilla WebSocket](https://pkg.go.dev/github.com/gorilla/websocket)
- [bolt db](https://github.com/boltdb/bolt)
  - [bolt browser](https://github.com/br0xen/boltbrowser)
  - [bolt web](https://github.com/evnix/boltdbweb)
  - [bbolt](https://github.com/etcd-io/bbolt)
- [gob](https://pkg.go.dev/encoding/gob) : value를 byte로 encode or decode

- [godoc](https://pkg.go.dev/golang.org/x/tools/cmd/godoc)

### Channels

```go
// Only Send Channel
func send(c chan<- int) {
	for i := range [10]int{} {
		fmt.Printf(">> sending %d <<\n", i)
		c <- i
		fmt.Printf(">> sent %d <<\n", i)
	}
	close(c) // avoid deadlock
}

// Only Receive Channel
func receive(c <-chan int) {
	for {
		time.Sleep(10 * time.Second)
		a, ok := <-c
		if !ok {
			fmt.Println("Done")
			break
		}
		fmt.Printf("|| received %d ||\n", a)
	}
}

func main() {
	c := make(chan int, 10) // Buffer Channel
	go send(c)
	receive(c)
}
```

## Tools

- [MVP.CSS](https://andybrewer.github.io/mvp/)
