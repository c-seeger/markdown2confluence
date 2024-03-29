# markdown2confluence

is a markdown to [Confluence](https://www.atlassian.com/software/confluence) xhtml converter library using [mark](https://github.com/a8m/mark) library

The implementation is just a markdown parser with some markdown comment syntax to add macro support.

> !!! Currently untested !!!

## Supportet Features

- markdown 2 xhtml
- xhtml fixed for confluence
- additional syntax via markdown comments to use macros
  - jira macro implemented
  - confluence page macro implemented
  - table of contents macro

## not yet supported
- attachments, comments, picture upload + xhtml code generation

If you miss some feature implementation, feel free to open an issue or send pull requests. I will take look as soon as possible.

## Installation

If you already installed GO on your system and configured it properly than its simply:

```
go get github.com/c-seeger/markdown2confluence
```

If not follow [these instructions](https://nats.io/documentation/tutorials/go-install/).

## Usage

### Simple example

```
package main

import (
        "fmt"
        "log"

        "github.com/c-seeger/markdown2confluence"
)

func main() {

  markdown := "#some markdown"

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

see [examples](https://github.com/c-seeger/markdown2confluence/tree/master/examples) for some more usage examples

basic macro support using the following syntax

#### Jira issue macro 
```
[//]: "jira:<issue-key>"
e.g. [//]: "jira:JIRA-1234"
```

#### confluence page links
```
[//]: "confluence:<page-title>"
e.g. [//]: "confluence:some page title"
```

### ToC macro

```
[//]: "toc:<max-level>:<type>:<outlined(true/false)>"
e.g: [//]: "toc:3:list:false"
```

[markdown2confluence](https://github.com/c-seeger/markdown2confluence) can also be used in combination with [confluence-go-api](https://github.com/c-seeger/confluence-go-api) see [confluence example](https://github.com/c-seeger/markdown2confluence/blob/master/examples/confluence.go)

## Code Documentation

- [Code documentation](https://godoc.org/github.com/c-seeger/markdown2confluence).
- [Confluence Storage Format](https://confluence.atlassian.com/doc/confluence-storage-format-790796544.html).
- [Macro definitions](https://confluence.pnac.org/display/DOC/Confluence+Storage+Format+for+Macros#ConfluenceStorageFormatforMacros-TableofContentsmacro)
