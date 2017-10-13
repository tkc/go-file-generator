package generator

import (
	"os"
	"io/ioutil"
	"github/tkc/go-file-generator/app/params"
)

type SourceData interface {
	GetParam() params.Param
	GetFileName() string
	Generate() string
}

func OutOut(s SourceData) {
	source := []byte(s.Generate())
	dir := "./docs/" + s.GetParam().Account
	MakeDir(dir)
	ioutil.WriteFile(dir+"/"+s.GetFileName(), source, 0644)
}

func MakeDir(dirName string) bool {
	if isExistDir(dirName) {
		return true
	}
	if err := os.Mkdir(dirName, 700); err != nil {
		return false
	} else {
		return true
	}
}

func isExistDir(dirName string) bool {
	_, err := os.Stat(dirName)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}
