# markdown2confluence

is a markdown to [Confluence](https://www.atlassian.com/software/confluence) xhtml converter library using [mark](https://github.com/a8m/mark) library 

## Supportet Features

- markdown 2 xhtml
- xhtml fixed for confluence
- additional syntax via markdown comments to use macros
-- jira macro implemented

## not yet supportet
- attachments, comments, picture upload + xhtml code generation

If you miss some feature implementation, feel free to open an issue or send pull requests. I will take look as soon as possible.

## Installation

If you already installed GO on your system and configured it properly than its simply:

```
go get github.com/cseeger-epages/markdown2confluence
```

If not follow [these instructions](https://nats.io/documentation/tutorials/go-install/).

## Usage

### Simple example

```
package main

import (
        "fmt"
        "log"

        "github.com/cseeger-epages/markdown2confluence"
)

func main() {

  // multiline markdown
  markdown := `# some markdown
- [md2conf](https://github.com/cseeger-epages/markdown2confluence)
## jira macro using markdown comment
[//]: "jira:JIRA-1234"`

  // Render xhtml
  xhtml, err := md2conf.Render(markdown)
  if err != nil {
    log.Fatal(err)
  }

  // confluence xhtml code
  fmt.Println(xhtml)
}
```

### Advanced examples

see [examples](https://github.com/cseeger-epages/markdown2confluence/tree/master/examples) for some more usage examples

## Code Documentation

You find the full [code documentation here](https://godoc.org/github.com/cseeger-epages/markdown2confluence).

The Confluence Storage Format documentation [can be found here](https://confluence.atlassian.com/doc/confluence-storage-format-790796544.html).
