package main

import (
	"io"
	"log"
	"os"

	"github.com/golang/snappy"
)

func main() {
	r := os.Stdin
	if len(os.Args) >= 2 {
		srcPath := os.Args[1]
		fr, err := os.Open(srcPath)
		if err != nil {
			log.Fatal(err)
		}
		defer fr.Close()
		r = fr
	}

	cw := snappy.NewBufferedWriter(os.Stdout)
	_, _ = io.Copy(cw, r)
	_ = cw.Close()
}
