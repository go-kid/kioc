package creator

import (
	"fmt"
	"github.com/go-kid/ioc/cmd/kioc/util"
	"path/filepath"
)

type GoFile struct {
	*File
}

func NewGoFile(template, pkg, fileName string, skipIfExists bool) *GoFile {
	f := &File{
		RelativePath: pkg,
		Template:     template,
		Ext:          "go",
		SkipIfExists: skipIfExists,
		Data:         nil,
		FileName:     strcase.SnakeCase(fileName),
		PostHandler: func(file *File) {
			var cmd util.GoCmd
			cmd.FmtAll(pkg)
		},
	}

	if f.SkipIfExists {
		f.SetAttribute("GenerateTitle", fmt.Sprintf("Package %s Code generated by gi2ctl once; SKIP IF EXISTS.", filepath.Base(pkg)))
	} else {
		f.SetAttribute("GenerateTitle", fmt.Sprintf("Code generated by kioc. DO NOT EDIT."))
	}
	f.SetAttribute("Pkg", filepath.Base(pkg))
	return &GoFile{
		File: f,
	}
}

type StructureGoFile struct {
}
