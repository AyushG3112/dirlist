package templates

import (
	"bytes"
	"html/template"

	"github.com/ayushg3112/dirlist/walk"

	_ "embed"
)

//go:embed static/singlepage.html
var singlepage string

func GenerateSinglePageTemplateHTML(structure []walk.DirectoryStructure) (string, error) {
	funcMap := template.FuncMap{
		"mul": func(i int, j int) int {
			return i * j
		},
	}

	t, err := template.New("singlepage").Funcs(funcMap).Parse(singlepage)

	if err != nil {
		return "", err
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, structure); err != nil {
		return "", err
	}

	return tpl.String(), nil
}
