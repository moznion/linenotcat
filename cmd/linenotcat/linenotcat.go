package main

import (
	"os"

	"github.com/moznion/linenotcat"
)

func main() {
	linenotcat.Run(os.Args[1:])
}
