package md2conf

import (
	"strings"

	"github.com/a8m/mark"
)

// Render creates confluence xhtml code from markdown
func Render(markdown string) (string, error) {
	macroed, err := applyMacros(markdown)
	if err != nil {
		return "", err
	}

	m := mark.New(macroed, nil)
	// header overwrite to support confluence ToC
	m.AddRenderFn(mark.NodeHeading, func(node mark.Node) string {
		return headerOverwriteHeading(node)
	})
	m.AddRenderFn(mark.NodeParagraph, func(node mark.Node) string {
		return headerOverwriteParagraph(node)
	})

	xhtml := m.Render()

	// hr fix
	xhtml = strings.Replace(xhtml, "<hr>", "<hr />", -1)
	// underscore fix
	xhtml = strings.Replace(xhtml, "<em>", " ", -1)
	xhtml = strings.Replace(xhtml, "</em>", " ", -1)

	return xhtml, nil

}
