package md2conf

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/a8m/mark"
)

// headerOverwrite is used to overwrite the standard mark HeadingNode function to
// add <a name"<text>"></a> for supporting confluence ToC
func headerOverwrite(n *mark.HeadingNode) (s string) {
	for _, node := range n.Nodes {
		s += node.Render()
	}

	re := regexp.MustCompile(`[^\w]+`)
	id := re.ReplaceAllString(n.Text, "-")
	// ToLowerCase
	id = strings.ToLower(id)
	return fmt.Sprintf("<%[1]s id=\"%s\"><a name=\"%s\"></a>%s</%[1]s>", "h"+strconv.Itoa(n.Level), id, s, s)

}
