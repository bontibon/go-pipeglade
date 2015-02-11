package main

import (
	"fmt"
	"flag"

	"github.com/bontibon/go-pipeglade"
)

var (
	uiFile = flag.String("u", "pipeglade.ui", "glade user interface file")
)

func main() {
	flag.Parse()

	ui, err := pipeglade.Exec(*uiFile)
	if err != nil {
		panic(err)
	}
	for {
		e, ok := <-ui.C
		if !ok {
			break
		}
		fmt.Printf("%+v\n", e)
	}
}
