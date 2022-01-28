package main

import (
	"flag"
	"fmt"

	"github.com/go-dummy/faker"
)

func main() {
	run()
}

func run() {
	f := faker.NewFaker()

	flag.Parse()

	a := flag.Arg(0)

	fmt.Println(f.ByName(a))
}
