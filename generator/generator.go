package main

import (
	"flag"
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type ModelsConfig struct {
	Models []ModelConfig `yaml:"models,flow"`
}

type ModelConfig map[string]string

var inputConfigFile = flag.String("file", "model.yml", "Input model config yaml file")

func main() {
	flag.Parse()
	for i := 0; i != flag.NArg(); i++ {
		fmt.Printf("arg[%d]=%s\n", i, flag.Arg(i))
	}

	mf, _ := ioutil.ReadFile(*inputConfigFile)
	mc := ModelsConfig{}
	merr := yaml.Unmarshal(mf, &mc)
	if merr != nil {
		fmt.Println("error:", merr)
	}

	fmt.Println("=== Models: ", mc.Models)
	for _, m := range mc.Models {
		fmt.Println("\n--", m["model"])
		for k, v := range m {
			if k != "model" {
				fmt.Println(k, v)
			}
		}
	}
}
