package main

import (
	"flag"
	"fmt"

	"github.com/neotoolkit/faker"
)

func main() {
	run()
}

func run() {
	flag.Parse()

	a := flag.Arg(0)

	fmt.Println(faker.Faker(a))
}
