package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"

	"gopkg.in/yaml.v2"
)

var inputConfigFile = flag.String("file", "model.yml", "Input model config yaml file")

type ModelAttr struct {
	Model  string
	Table  string
	Keys   []string
	Values []string
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
		keys := make([]string, 0)
		values := make([]string, 0)
		for _, v := range j {
			if v.Key != "model" {
				keys = append(keys, v.Key.(string))
				values = append(values, v.Value.(string))
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
		m := ModelAttr{modelname, table, keys, values}
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
