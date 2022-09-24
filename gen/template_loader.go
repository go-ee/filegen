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
	macrosSources    []string
	templatesSources []string
	templatesNames   []string
	tmpl             *template.Template

	currentName   string
	currentSource string
	currentIndex  int
}

func NewNextTemplateProviderFromFiles(
	templatesFiles []string, macrosTemplatesFiles []string) (ret *NextTemplateProvider, err error) {

	tmpl := template.New("file-gen").Funcs(FuncMap).Funcs(sprig.TxtFuncMap())

	if _, err = tmpl.ParseFiles(macrosTemplatesFiles...); err != nil {
		return
	}
	if _, err = tmpl.ParseFiles(templatesFiles...); err != nil {
		return
	}

	names := make([]string, len(templatesFiles))
	for i, templateFile := range templatesFiles {
		names[i] = filepath.Base(templateFile)
	}

	ret = &NextTemplateProvider{
		macrosSources:    macrosTemplatesFiles,
		templatesSources: templatesFiles,
		currentIndex:     0,
		templatesNames:   names,
	}
	return
}

type TemplateSource struct {
	Text   string
	Source string
}

func NewNextTemplateProviderFromText(
	templates []*TemplateSource, macrosTemplates []*TemplateSource) (ret *NextTemplateProvider, err error) {

	tmpl := newTemplate("file-gen")

	macrosSources := make([]string, len(macrosTemplates))
	for i, childTmplSource := range macrosTemplates {
		macrosSources[i] = childTmplSource.Source
		fileName := filepath.Base(childTmplSource.Source)
		childTmpl := newTemplate(fileName)
		if childTmpl, err = childTmpl.Parse(childTmplSource.Text); err != nil {
			return
		}
		if _, err = tmpl.AddParseTree(childTmpl.Name(), childTmpl.Tree); err != nil {
			return
		}
	}

	names := make([]string, len(templates))
	templatesSources := make([]string, len(templates))
	for i, childTmplSource := range templates {
		templatesSources[i] = childTmplSource.Source
		fileName := filepath.Base(childTmplSource.Source)
		childTmpl := newTemplate(fileName)
		if childTmpl, err = childTmpl.Parse(childTmplSource.Text); err != nil {
			return
		}
		if _, err = tmpl.AddParseTree(childTmpl.Name(), childTmpl.Tree); err != nil {
			return
		}
		names[i] = fileName
	}

	ret = &NextTemplateProvider{
		macrosSources:    macrosSources,
		templatesSources: templatesSources,
		templatesNames:   names,
		tmpl:             tmpl,
	}
	return
}

func newTemplate(templateName string) *template.Template {
	return template.New(templateName).Funcs(FuncMap).Funcs(sprig.TxtFuncMap())
}

func (o *NextTemplateProvider) Next() (ret TemplateLoader) {
	if o.currentIndex < len(o.templatesSources) {
		ret = o
		o.currentName = o.templatesNames[o.currentIndex]
		o.currentSource = o.templatesSources[o.currentIndex]

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
