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
		h, _ := node.(*mark.HeadingNode)
		return headerOverwrite(h)
	})

	xhtml := m.Render()

	// hr fix
	xhtml = strings.Replace(xhtml, "<hr>", "<hr />", -1)

	return xhtml, nil

}
