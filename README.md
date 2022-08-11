# gchain

[![GoDoc](https://godoc.org/github.com/t-katsumura/gchain?status.svg)](http://godoc.org/github.com/t-katsumura/gchain)
[![Go Report Card](https://goreportcard.com/badge/github.com/t-katsumura/gchain)](https://goreportcard.com/report/github.com/t-katsumura/gchain)
[![MIT licensed](https://img.shields.io/badge/license-MIT-blue.svg)](./LICENSE)
[![Test](https://github.com/t-katsumura/gchain/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/t-katsumura/gchain/actions/workflows/test.yml?query=branch%3Amain)
[![Coverage](https://gocover.io/_badge/github.com/t-katsumura/gchain)](https://gocover.io/github.com/t-katsumura/gchain)
[![GitHub release](https://img.shields.io/github/release/t-katsumura/gchain/all.svg?style=flat-square)](https://github.com/t-katsumura/gchain/releases)


gchain is a simple go library for creation of generalized method chain.
Any types of method chain can be created using gchain which is utilizing go Generics.

In particular, when creating method chain, it usually becomes

```go
chain = firstFunc(secondFunc(thirdFunc(finalFunc)))
```

gchain makes it like

```go
chain = gchain.NewChainXtoX(firstFunc, secondFunc, thirdFunc).Chain(finalFunc)
```

## Example

This is an example of gchain.

It creating http handler with using gchain (`handler2`, `handler3`) and without using gchain (`handler1`).

```
package main

import (
	"net/http"

	"github.com/t-katsumura/gchain"
)

func firstHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hi from first handler\n"))
		next.ServeHTTP(w, r)
	})
}

func secondHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hi from second handler\n"))
		next.ServeHTTP(w, r)
	})
}
func thirdHandler(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("Hi from third handler\n"))
}

func main() {
	// Not using gchain
	handler1 := firstHandler(secondHandler(http.HandlerFunc(thirdHandler)))

	// Using gchain
	handler2 := gchain.NewChainXtoX(firstHandler, secondHandler).
		Chain(http.HandlerFunc(thirdHandler))

	// Using gchain
	chain := gchain.NewChainXtoX[http.Handler]()
	chain.Append(firstHandler)
	chain.Append(secondHandler)
	handler3 := chain.Chain(http.HandlerFunc(thirdHandler))

    // Run http server with handler1, handler2, handler3
	http.Handle("/h1", handler1)
	http.Handle("/h2", handler2)
	http.Handle("/h3", handler3)
	http.ListenAndServe(":8080", nil)
}
```

## Questions and support
All bug reports, questions and suggestions should go though Github Issues.

## Contributing
1. Fork it
1. Create feature branch (`git checkout -b feature/new-feature`)
1. Write codes on feature branch
1. Commit your changes (`git commit -m "Added new feature"`)
1. Push to the branch (`git push origin feature/new-feature`)
1. Create new Pull Request on Github

## Development
- Write codes
- `go fmt -x  ./...` - format codes
- `go test -v -cover ./...` - run test and the coverage should always be 100%