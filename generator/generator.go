package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
	"unicode"

	"gopkg.in/yaml.v2"
)

var inputConfigFile = flag.String("file", "model.yml", "Input model config yaml file")

type ModelAttr struct {
	Model   string
	Table   string
	Imports []string
	Attrs   []string
	Keys    []string
	Values  []string
	Columns []string
}

func ucfirst(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}

func main() {
	flag.Parse()
	for i := 0; i != flag.NArg(); i++ {
		fmt.Printf("arg[%d]=%s\n", i, flag.Arg(i))
	}

	mf, _ := ioutil.ReadFile(*inputConfigFile)
	ms := make(map[string][]yaml.MapSlice)
	merr := yaml.Unmarshal(mf, &ms)
	if merr != nil {
		fmt.Println("error:", merr)
	}
	for _, j := range ms["models"] {
		var modelname, table, filename string
		imports := make([]string, 0)
		attrs := make([]string, 0)
		keys := make([]string, 0)
		values := make([]string, 0)
		columns := make([]string, 0)
		imports = append(imports, "fmt")
		for _, v := range j {
			if v.Key != "model" {
				attrs = append(attrs, ucfirst(v.Key.(string)))
				keys = append(keys, v.Key.(string))
				values = append(values, v.Value.(string))
				c := v.Value.(string)
				if c == "string" {
					c = "VARCHAR(255)"
				} else if c == "int64" {
					c = "BIGINT"
				} else if c == "time.Time" {
					c = "DATETIME"
					imports = append(imports, "time")
				}
				columns = append(columns, c)
			} else {
				modelname = v.Value.(string)
				table = strings.ToLower(modelname)
				filename = "model/" + modelname + ".go"
			}
		}
		fmt.Println("-- Generate", filename)
		t, err := template.ParseFiles("generator/model.template")
		if err != nil {
			fmt.Println(err)
			return
		}
		var b bytes.Buffer
		m := ModelAttr{modelname, table, imports, attrs, keys, values, columns}
		t.Execute(&b, m)
		fmt.Println(b.String())

		// Write to file
		f, err := os.Create(filename)
		if err != nil {
			fmt.Println("create file: ", err)
			return
		}
		err = t.Execute(f, m)
		if err != nil {
			fmt.Print("execute: ", err)
			return
		}
		f.Close()

	}

}
