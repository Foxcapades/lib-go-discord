package main

import (
	"os"
	"strings"
	"text/template"

	"github.com/foxcapades/lib-go-discord/tools/fields/internal/data"
	"github.com/foxcapades/lib-go-discord/tools/fields/internal/types"
)

const (
	tplPath = "tools/fields/"
	outPath = "pkg/dlib/"
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
