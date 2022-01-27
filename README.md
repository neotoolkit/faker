Fake data generator. Written in Go

[![Go Report Card](https://goreportcard.com/badge/github.com/go-dummy/faker)](https://goreportcard.com/report/github.com/go-dummy/faker)
[![codecov](https://codecov.io/gh/go-dummy/faker/branch/main/graph/badge.svg?token=HCOXZU6MFB)](https://codecov.io/gh/go-dummy/faker)

### Installation
Faker requires Go > 1.17
```bash
go get github.com/go-dummy/faker
```
### Usage
```go
package main

import (
	"fmt"

	"github.com/go-dummy/faker"
)

func main() {
	f := faker.NewFaker()

	fmt.Println(f.Person().Name())
}
```
```bash
Elon Musk
```
