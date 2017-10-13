package main

import (
	"flag"
	"github/tkc/go-file-generator/app/params"
	"github/tkc/go-file-generator/app/template/swift"
	"github/tkc/go-file-generator/app/template/android"
	"github/tkc/go-file-generator/app/helper/generator"
	"github/tkc/go-file-generator/app/helper/yaml"
)

var (
	configName = flag.String("c", "default", "help message for s option")
)

func main() {

	flag.Parse()

	p := params.Param{
		Account: yaml.GetAccount(*configName),
	}

	swift := swift.Create(p)
	android := android.Create(p)

	generator.OutOut(swift)
	generator.OutOut(android)

}
