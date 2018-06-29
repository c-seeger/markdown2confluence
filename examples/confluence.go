package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/cseeger-epages/confluence-go-api"
	"github.com/cseeger-epages/markdown2confluence"
)

func main() {
	// initialize your confluence api
	api, err := goconfluence.NewAPI("https://<your-domain>.atlassian.net", "<username>", "<password>")
	if err != nil {
		log.Fatal(err)
	}

	// get markdown from file
	filename := "test.md"
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	page := string(buf)

	// render page
	xhtml, err := md2conf.Render(page)

	// create content
	data := &goconfluence.Content{
		Version: goconfluence.Version{
			Number: hist.LastUpdated.Number + 1,
		},
		Type:  "page",
		Title: "title",
		Ancestors: []goconfluence.Ancestor{
			goconfluence.Ancestor{},
		},
		Body: goconfluence.Body{
			Storage: goconfluence.Storage{
				Value:          xhtml,
				Representation: "storage",
			},
		},
		Space: goconfluence.Space{
			Key: "SPACE-KEY",
		},
	}

	c, err := api.createContent(data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", c)

}
