package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/jmhobbs/go-raP"
	"github.com/jmhobbs/go-raP/printer"
)

func main() {
	flag.Usage = func() {
		_, err := fmt.Fprintln(flag.CommandLine.Output(), "usage: rap-decode <file.bin>")
		if err != nil {
			panic(err)
		}
	}
	flag.Parse()

	f, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer func() {
		err := f.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	root, err := raP.Decode(f)
	if err != nil {
		panic(err)
	}

	p := printer.New()
	err = p.File(os.Stdout, root)
	if err != nil {
		panic(err)
	}
}
