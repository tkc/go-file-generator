package android

import (
	"fmt"
	"github/tkc/go-file-generator/app/params"
)

func Create(param params.Param) Android {
	return Android{
		FileName: fileName,
		Param:    param,
	}
}

type Android struct {
	FileName string
	Param    params.Param
}

func (a Android) GetParam() params.Param {
	return a.Param
}

func (a Android) GetFileName() string {
	return a.FileName
}

func (a Android) Generate() string {
	return fmt.Sprintf(
		template,
		a.Param.Account,
	)
}

const fileName = "common.jave"

const template = `
public class Common {
 public static final String ACCOUNT = "%s";
}

`
