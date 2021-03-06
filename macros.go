package md2conf

import (
	"fmt"
	"regexp"
	"strings"
)

// applyMacros to apply Macros
func applyMacros(markdown string) (string, error) {
	macroed, err := jiraMacro(markdown)
	if err != nil {
		return "", err
	}
	macroed, err = confluencePageMacro(macroed)
	if err != nil {
		return "", err
	}
	macroed, err = tableOfContentsMacro(macroed)
	if err != nil {
		return "", err
	}
	return macroed, nil
}

// jiraMacro creates jira macro xhtml code
// https://confluence.atlassian.com/conf59/jira-issues-macro-792499129.html
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

// confluencePageMacro adds page links to other confluence pages
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

// tableOfContentsMacro adds confluence ToC
// https://confluence.pnac.org/display/DOC/Confluence+Storage+Format+for+Macros#ConfluenceStorageFormatforMacros-TableofContentsmacro
func tableOfContentsMacro(html string) (string, error) {
	re := regexp.MustCompile(`\[//\]: "toc:(?P<GM>.*)"`)
	if !re.MatchString(html) {
		return html, nil
	}

	macroTemplate := `
<ac:macro ac:name="toc">
	<ac:parameter ac:name="printable">true</ac:parameter>
	<ac:parameter ac:name="style">square</ac:parameter>
	<ac:parameter ac:name="maxLevel">###LEVEL###</ac:parameter>
	<ac:parameter ac:name="indent">5px</ac:parameter>
	<ac:parameter ac:name="minLevel">1</ac:parameter>
	<ac:parameter ac:name="class">bigpink</ac:parameter>
	<ac:parameter ac:name="exclude">[1//2]</ac:parameter>
	<ac:parameter ac:name="type">###TYPE###</ac:parameter>
	<ac:parameter ac:name="outline">###OUTLINE###</ac:parameter>
	<ac:parameter ac:name="include">.*</ac:parameter>
</ac:macro>`

	list := re.FindAllString(html, -1)
	for _, v := range list {
		data := strings.Split(v, ":")
		if len(data) < 5 {
			return "", fmt.Errorf("TABLE OF CONTENS MACRO ERROR: not enough arguments")
		}
		tocMacro := strings.Replace(macroTemplate, "###LEVEL###", data[2], 1)
		tocMacro = strings.Replace(tocMacro, "###TYPE###", data[3], 1)
		tocMacro = strings.Replace(tocMacro, "###OUTLINE###", data[4][:len(data[3])-1], 1)
		html = strings.Replace(html, v, tocMacro, 1)
	}

	return html, nil
}
