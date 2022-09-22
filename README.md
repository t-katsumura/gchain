# gchain

[![GoDoc](https://godoc.org/github.com/t-katsumura/gchain?status.svg)](http://godoc.org/github.com/t-katsumura/gchain)
[![Go Report Card](https://goreportcard.com/badge/github.com/t-katsumura/gchain)](https://goreportcard.com/report/github.com/t-katsumura/gchain)
[![MIT licensed](https://img.shields.io/badge/license-MIT-blue.svg)](./LICENSE)
[![Test](https://github.com/t-katsumura/gchain/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/t-katsumura/gchain/actions/workflows/test.yml?query=branch%3Amain)
[![Codecov](https://codecov.io/gh/t-katsumura/gchain/branch/main/graph/badge.svg?token=P5J4J1F6RN)](https://codecov.io/gh/t-katsumura/gchain)

Crate a method chain using Generics in Go.  
Any types of method can be chained.

> **Warning**
> Generics is introduced since go 1.18.

In particular, when creating method chain, it usually becomes

```go
chain = firstFunc(secondFunc(thirdFunc(finalFunc)))
```

gchain makes it simple

```go
chain = gchain.NewChainXtoX(firstFunc, secondFunc, thirdFunc).
        Chain(finalFunc)
```

or

```go
chain = gchain.NewChainXtoX(firstFunc).
        Append(secondFunc).
        Append(thirdFunc).
        Chain(finalFunc)
```

## Installation

```
go get github.com/t-katsumura/gchain@latest
```

## Usage

**Create a new chain**

Create a new chain.  
Foe example, when creating $x_{n+1} = f^{(n)}(x_{n})$ types of chaining

```go
// create an empty chain
chain := gchain.NewChainXtoX()

or

// create a chain with methods
chain := gchain.NewChainXtoX(f1, f2, f3)
```

**Add some methods to the chain**

Adding methods to the chain.

```go
chain.Append(f1)
chain.Append(f2)
```

**Get the method chain**

Get the method chain.

```go
method_chain := chain.Chain(x0)
```

## Example

This is an example.

It creates middleware chain for http server.

```go
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

    // this handler is the same as firstHandler(secondHandler(thirdHandler))
	handler := gchain.NewChainXtoX(secondHandler, firstHandler).
		Chain(http.HandlerFunc(thirdHandler))

    // This server returns
    //     Hi from first handler
    //     Hi from second handler
    //     Hi from third handler
	http.ListenAndServe(":8080", handler)

}
```

## Structs

Structs' name is in the format of `ChainXYZAtoX`.  
Here, `X`, `Y`, `Z`, `A` represents a Data Type.`A` means an array.  
`XYZA` is arguments and the `X` is the type of returned value.

This means `ChainXYZAtoX` is equivalent to the following equation.

$x_{i+1} = f^{(i)}(x_{i}, y, z, a)$

```go
// when using ChainXYZAtoX
// methods must have the following signature
func(x X, y Y, z Z, a ...A) X
```

```go
// when using ChainXYZA
// methods must have the following signature
func(x X, y Y, z Z, a ...A)
```

```go
// when using ChainXYZ
// methods must have the following signature
func(x X, y Y, z Z)
```

## Questions and support

All bug reports, questions and suggestions should go though Github Issues.

## Contributing

1. Fork it
1. Create feature branch (`git checkout -b feature/new-feature`)
1. Write codes on feature branch
1. Commit your changes (`git commit -m "Added new feature"`)
1. Push to the branch (`git push -u origin feature/new-feature`)
1. Create new Pull Request on Github

## Development

- Write codes
- `go fmt -x ./...` - format codes
- `go test -v -cover ./...` - run test and the coverage should always be 100%
