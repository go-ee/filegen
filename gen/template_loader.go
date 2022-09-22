package gen

import (
	"github.com/Masterminds/sprig"
	"path/filepath"
	"text/template"
)

type TemplateLoader interface {
	TemplateSource() string
	LoadTemplate() (ret *template.Template, err error)
}

type NextTemplateProvider struct {
	macrosTemplates []string
	templates       []string
	names           []string
	tmpl            *template.Template

	currentName   string
	currentSource string
	currentIndex  int
}

func NewNextTemplateProvider(
	templates []string, macrosTemplates []string) (ret *NextTemplateProvider, err error) {

	ret = &NextTemplateProvider{
		macrosTemplates: macrosTemplates,
		templates:       templates,
		currentIndex:    0,
	}
	err = ret.init()
	return
}

func (o *NextTemplateProvider) init() (err error) {
	o.tmpl = template.New("filegen").Funcs(FuncMap).Funcs(sprig.TxtFuncMap())

	if _, err = o.tmpl.ParseFiles(o.macrosTemplates...); err != nil {
		return
	}
	if _, err = o.tmpl.ParseFiles(o.templates...); err != nil {
		return
	}

	o.names = make([]string, len(o.templates))
	for i, templateFile := range o.templates {
		o.names[i] = filepath.Base(templateFile)
	}
	return
}

func (o *NextTemplateProvider) Next() (ret TemplateLoader) {
	if o.currentIndex < len(o.templates) {
		ret = o
		o.currentName = o.names[o.currentIndex]
		o.currentSource = o.templates[o.currentIndex]

		o.currentIndex++
	}
	return
}

func (o *NextTemplateProvider) Reset() {
	o.currentIndex = 0
}

func (o *NextTemplateProvider) TemplateSource() string {
	return o.currentSource
}

func (o *NextTemplateProvider) LoadTemplate() (ret *template.Template, err error) {
	ret = o.tmpl.Lookup(o.currentName)
	return
}
