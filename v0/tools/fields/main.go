// +build ignore

package main

import (
	"os"
	"strings"
	"text/template"

	"github.com/foxcapades/lib-go-discord/v0/tools/fields/internal/data"
	"github.com/foxcapades/lib-go-discord/v0/tools/fields/internal/types"
)

const (
	tplPath = "v0/tools/fields/"
	outPath = "v0/pkg/discord/"
)

func main() {
	fun := template.FuncMap{
		"join": strings.Join,
	}

	tpl := template.Must(new(template.Template).
		Funcs(fun).
		ParseGlob(tplPath + "*.gotpl"))

	files := map[string]string{
		"nullable-fields.go":      "nullable",
		"nullable-fields_test.go": "nullable-test",
		"optional-fields.go":      "optional",
		"tri-state-fields.go":     "tri-state",
	}

	for k, v := range files {
		writeFile(k, v, tpl)
	}
}

func writeFile(fn, name string, tpl *template.Template) {
	file, err := os.Create(outPath + fn)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = tpl.ExecuteTemplate(file, name, types.Wrapper{Types: data.Fields})
	if err != nil {
		panic(err)
	}
}
