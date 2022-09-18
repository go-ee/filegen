package gen

import (
	"github.com/Masterminds/sprig"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"
)

type TemplateDataLoader interface {
	TemplateLabel() string
	LoadTemplateData() ([]byte, error)
}

type TemplateFactory interface {
	CreateTemplate() (ret *template.Template, err error)
}

type OutputFileNameBuilder interface {
	BuildOutputFileName(templateLabel string) string
	BuildOutputFileNameForFileName(templateLabel string, customFileName string) string
}

type MultipleFileData struct {
	Files []struct {
		FileName string
		Data     interface{}
	}
}

type FileTemplateFactory struct {
	TemplateFile string
}

func (o *FileTemplateFactory) CreateTemplate() (ret *template.Template, err error) {
	ret, err = template.New(path.Base(o.TemplateFile)).Funcs(FuncMap).Funcs(sprig.TxtFuncMap()).
		ParseFiles(o.TemplateFile)
	return
}

type DefaultsOutputFileNameBuilder struct {
	OutputDirectory  string
	DefaultExtension string
}

func (o *DefaultsOutputFileNameBuilder) BuildOutputFileName(templateLabel string) string {
	fileName := strings.TrimSuffix(filepath.Base(templateLabel), filepath.Ext(templateLabel)) + o.DefaultExtension
	return o.BuildOutputFileNameForFileName(templateLabel, fileName)
}

func (o *DefaultsOutputFileNameBuilder) BuildOutputFileNameForFileName(templateLabel string, fileName string) string {
	return path.Join(o.OutputDirectory, path.Dir(templateLabel), fileName)
}

type FileTemplateDataLoader struct {
	FileName string
}

func (o *FileTemplateDataLoader) TemplateLabel() string {
	return o.FileName
}

func (o *FileTemplateDataLoader) LoadTemplateData() (ret []byte, err error) {
	if ret, err = os.ReadFile(o.FileName); err == nil {
		ret, err = ToJSON(ret)
	}
	return
}
