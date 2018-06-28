package md2conf

import (
	"fmt"
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

	macroed, err = confluencePageMacro(macroed)
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
		if len(jiraCase) < 3 {
			return "", fmt.Errorf("JIRA MACRO ERROR: no jira issue found")
		}
		jiraMacro := strings.Replace(macroTemplate, "###JIRA###", jiraCase[2][:len(jiraCase[2])-1], 1)
		html = strings.Replace(html, v, jiraMacro, 1)
	}

	return html, nil
}

func confluencePageMacro(html string) (string, error) {
	re := regexp.MustCompile(`\[//\]: "confluence:(?P<GM>.*)"`)
	if !re.MatchString(html) {
		return html, nil
	}

	macroTemplate := `
<ac:link>
	<ri:page ri:content-title="###TITLE###" />
</ac:link> `

	list := re.FindAllString(html, -1)
	for _, v := range list {
		data := strings.Split(v, ":")
		if len(data) < 3 {
			return "", fmt.Errorf("CONFLUENCE PAGE MACRO ERROR: no page title found")
		}
		confluencePageMacro := strings.Replace(macroTemplate, "###TITLE###", data[2][:len(data[2])-1], 1)
		html = strings.Replace(html, v, confluencePageMacro, 1)
	}

	return html, nil
}
