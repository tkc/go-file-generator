package swift

import (
	"fmt"
	"github/tkc/go-file-generator/app/params"
)

func Create(param params.Param) Swift {
	return Swift{
		FileName: fileName,
		Param:    param,
	}
}

type Swift struct {
	FileName string
	Param    params.Param
}

func (s Swift) GetParam() params.Param {
	return s.Param
}

func (s Swift) GetFileName() string {
	return s.FileName
}

func (s Swift) Generate() string {
	return fmt.Sprintf(
		template,
		s.Param.Account,
	)
}

const fileName = "common.swift"

const template = `

 let ACCOUNT = "%s";

`
