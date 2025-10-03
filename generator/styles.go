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

var styleTemplate = template.Must(template.New("style").Parse(`
type style interface {
	String(...string) string
	Sprint(...any) string
	Sprintf(string, ...any) string
	Sprintln(...any) string
	Print(...any)
	Printf(string, ...any)
	Println(...any)
	WithFG(uint8) style
	WithBG(uint8) style
{{- range . }}
	With{{ . }}() style
	{{ . }}(...string) string
{{- end }}
}
`))

var styleFuncsTemplate = template.Must(template.New("style_funcs").Parse(`
func With{{ .name }}() style {
	return defaultComposer.With{{.name}}()
}

func {{ .name }}(str ...string) string {
	return With{{ .name }}().String(str...)
}
`))

func generateStyles() {
	var buf bytes.Buffer

	fmt.Fprint(&buf, "package pigment\n")

	styleNames := make([]string, len(ansi.Styles))
	for i, s := range ansi.Styles {
		styleNames[i] = s.Name
	}
	if err := styleTemplate.Execute(&buf, styleNames); err != nil {
		panic(err)
	}
	fmt.Fprintln(&buf)

	for _, s := range ansi.Styles {
		args := map[string]string{
			"name":  s.Name,
			"style": strings.ToLower(s.Name),
		}
		if err := styleFuncsTemplate.Execute(&buf, args); err != nil {
			panic(err)
		}
		fmt.Fprintln(&buf)
	}

	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		panic(fmt.Sprintf("failed to format generated code: %v", err))
	}

	f, err := os.Create("styles.go")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if _, err := f.Write(formatted); err != nil {
		panic(err)
	}

	fmt.Println("Generated and formatted: styles.go")
}
