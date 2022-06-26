# faker

[![build-img]][build-url]
[![pkg-img]][pkg-url]
[![reportcard-img]][reportcard-url]
[![coverage-img]][coverage-url]
[![version-img]][version-url]

Fake data generator

## Features
- Zero dependencies
- Supports data `boolean`, `username`, `domain`, `email`, `firstname`, `lastname`, `name`, `gender`, `uuid`, `ip`
- Easy to integrate.

## Installation
```shell
go get github.com/neotoolkit/faker
```

## Usage
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

## Documentation

See [these docs][pkg-url].

## License

[MIT License](LICENSE).

[build-img]: https://github.com/neotoolkit/faker/workflows/build/badge.svg
[build-url]: https://github.com/neotoolkit/faker/actions
[pkg-img]: https://pkg.go.dev/badge/neotoolkit/faker
[pkg-url]: https://pkg.go.dev/github.com/neotoolkit/faker
[reportcard-img]: https://goreportcard.com/badge/neotoolkit/faker
[reportcard-url]: https://goreportcard.com/report/neotoolkit/faker
[coverage-img]: https://codecov.io/gh/neotoolkit/faker/branch/main/graph/badge.svg
[coverage-url]: https://codecov.io/gh/neotoolkit/faker
[version-img]: https://img.shields.io/github/v/release/neotoolkit/faker
[version-url]: https://github.com/neotoolkit/faker/releases

## Sponsors
<p>
  <a href="https://evrone.com/?utm_source=github&utm_campaign=neotoolkit">
    <img src="https://raw.githubusercontent.com/neotoolkit/.github/main/assets/sponsored_by_evrone.svg"
      alt="Sponsored by Evrone">
  </a>
</p>
