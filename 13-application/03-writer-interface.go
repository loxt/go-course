package main

import (
	"io"
	"os"
)

func main() {
	_, _ = io.WriteString(os.Stdout, "hello, loxt")
}
