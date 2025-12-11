package generate

import (
	"fmt"
	"strconv"
	"strings"
	"text/template"

	"github.com/bbq-yaozi/go-mdi/internal/assets"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var caser = cases.Title(language.English)

const (
	TypesGoFile = "templates/types.go.tmpl"
	FuncsGoFile = "templates/funcs.go.tmpl"
	IconGoFile  = "templates/icon.go.tmpl"
)

func LoadTemplate(name string) (*template.Template, error) {
	templData, err := assets.ReadFile(name)
	if err != nil {
		return nil, err
	}

	funcs := template.FuncMap{
		"PascalCase":      pascalCase,
		"BytesToString":   bytesToString,
		"StringsToString": stringsToString,
	}

	return template.New("model").Funcs(funcs).Parse(string(templData))
}

func pascalCase(name string) string {
	if name == "" {
		return name
	}

	words := strings.Split(name, "-")
	for i := 0; i < len(words); i++ {
		if words[i] != "" {
			words[i] = caser.String(words[i])
		}
	}
	return strings.Join(words, "")
}

func bytesToString(data []byte) string {
	if len(data) == 0 {
		return ""
	}

	parts := make([]string, len(data))
	for i, b := range data {
		parts[i] = strconv.Itoa(int(b))
	}

	return strings.Join(parts, ", ")
}

func stringsToString(data []string) string {
	if len(data) == 0 {
		return ""
	}

	parts := make([]string, len(data))
	for i, b := range data {
		parts[i] = fmt.Sprintf(`"%s"`, b)
	}

	return strings.Join(parts, ", ")
}
