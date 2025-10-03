package main

import (
	"bytes"
	"fmt"
	"go/format"
	"os"
	"strings"
	"text/template"

	"github.com/scallywaag/pigment/ansi"
)

var composerTemplate = template.Must(template.New("composer").Parse(`
type composer struct {
	node *styleNode
	fg *composer
	bg *composer
{{- range . }}
	{{ . }} *composer
{{- end }}
}
`))

var composerFuncsTemplate = template.Must(template.New("composer_funcs").Parse(`
func (c *composer) With{{ .name }}() style {
	if c.{{ .style }} == nil {
		sp := createStylePair(ansi.{{ .name }}.Code, ansi.{{ .name }}.Reset)
		c.{{ .style }} = createComposer(c, sp.code, sp.reset)
	}
	return c.{{ .style }}
}

func (c *composer) {{ .name }}(str ...string) string {
	return c.With{{ .name }}().String(str...)
}
`))

func generateComposer() {
	var buf bytes.Buffer

	fmt.Fprint(&buf, "package pigment\n")
	fmt.Fprint(&buf, `import "github.com/scallywaag/pigment/ansi"`+"\n")

	composerFields := make([]string, len(ansi.Styles))
	for i, s := range ansi.Styles {
		composerFields[i] = strings.ToLower(s.Name)
	}
	if err := composerTemplate.Execute(&buf, composerFields); err != nil {
		panic(err)
	}
	fmt.Fprintln(&buf)

	for _, s := range ansi.Styles {
		args := map[string]string{
			"name":  s.Name,
			"style": strings.ToLower(s.Name),
		}
		if err := composerFuncsTemplate.Execute(&buf, args); err != nil {
			panic(err)
		}
		fmt.Fprintln(&buf)
	}

	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		panic(fmt.Sprintf("failed to format generated code: %v", err))
	}

	f, err := os.Create("composer.go")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if _, err := f.Write(formatted); err != nil {
		panic(err)
	}

	fmt.Println("Generated and formatted: composer.go")
}
