Fake data generator. Written in Go

[![Go Report Card](https://goreportcard.com/badge/github.com/neotoolkit/faker)](https://goreportcard.com/report/github.com/neotoolkit/faker)
[![codecov](https://codecov.io/gh/neotoolkit/faker/branch/main/graph/badge.svg?token=HCOXZU6MFB)](https://codecov.io/gh/neotoolkit/faker)

### Installation
Faker requires Go > 1.17
```bash
go get github.com/neotoolkit/faker
```
### Usage
```go
package main

import (
	"fmt"

	"github.com/neotoolkit/faker"
)

func main() {
	f := faker.NewFaker()

	fmt.Println(f.Person().Name())
}
```
```bash
Elon Musk
```
