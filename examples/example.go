package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/c-seeger/markdown2confluence"
)

func main() {
	bytes, err := ioutil.ReadFile("test.md")
	if err != nil {
		log.Fatal(err)
	}

	xhtml, err := md2conf.Render(string(bytes))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(xhtml)
}
