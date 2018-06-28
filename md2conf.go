package md2conf

import (
	"regexp"
	"strings"

	"github.com/a8m/mark"
)

// Render creates confluence xhtml code from markdown
func Render(markdown string) (string, error) {
	macroed, err := jiraMacro(markdown)
	if err != nil {
		return "", err
	}

	xhtml := mark.Render(macroed)
	xhtml = strings.Replace(xhtml, "<hr>", "<hr />", -1)
	return xhtml, nil

}

// jiraMacro creates jira macro xhtml code
func jiraMacro(html string) (string, error) {
	re := regexp.MustCompile(`\[//\]: "jira:(?P<GM>.*)"`)
	if !re.MatchString(html) {
		return html, nil
	}

	macroTemplate := `
<ac:structured-macro ac:name="jira">
	<ac:parameter ac:name="columns">key,summary,status</ac:parameter>
	<ac:parameter ac:name="key">###JIRA###</ac:parameter>
</ac:structured-macro>`

	list := re.FindAllString(html, -1)
	for _, v := range list {
		jiraCase := strings.Split(v, ":")
		jiraMacro := strings.Replace(macroTemplate, "###JIRA###", jiraCase[2][:len(jiraCase[2])-1], 1)
		html = strings.Replace(html, v, jiraMacro, 1)
	}

	return html, nil
}
