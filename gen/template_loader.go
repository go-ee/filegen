package gen

import (
	"github.com/Masterminds/sprig"
	"path"
	"text/template"
)

type TemplateLoader interface {
	TemplateLabel() string
	LoadTemplate() (ret *template.Template, err error)
}

func NewFileTemplateLoader(templateFile string) *FileTemplateLoader {
	return &FileTemplateLoader{
		File: templateFile,
	}
}

type FileTemplateLoader struct {
	File string
}

func (o *FileTemplateLoader) TemplateLabel() string {
	return o.File
}

func (o *FileTemplateLoader) LoadTemplate() (ret *template.Template, err error) {
	ret, err = template.New(path.Base(o.File)).Funcs(FuncMap).Funcs(sprig.TxtFuncMap()).
		ParseFiles(o.File)
	return
}
