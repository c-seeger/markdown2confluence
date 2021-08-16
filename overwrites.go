package md2conf

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/a8m/mark"
)

// headerOverwriteParagraph is used to overwrite the standard mark ParagraphNode function to
// add <a name"<text>"></a> for supporting confluence ToC
func headerOverwriteParagraph(node mark.Node) (s string) {
	p, _ := node.(*mark.ParagraphNode)
	for _, node := range p.Nodes {
		s += node.Render()
	}
	return s
}

// headerOverwriteHeading is used to overwrite the standard mark HeadingNode function to
// add <a name"<text>"></a> for supporting confluence ToC
func headerOverwriteHeading(node mark.Node) (s string) {
	n, _ := node.(*mark.HeadingNode)
	re := regexp.MustCompile(`[^\w]+`)
	id := re.ReplaceAllString(n.Text, "-")
	// ToLowerCase
	id = strings.ToLower(id)
	return fmt.Sprintf("<%[1]s id=\"%s\"><a name=\"%s\"></a>%s</%[1]s>", "h"+strconv.Itoa(n.Level), id, s, s)

}
