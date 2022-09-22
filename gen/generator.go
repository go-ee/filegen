package gen

import (
	"bytes"
	"encoding/json"
	"github.com/ghodss/yaml"
	"github.com/go-ee/utils/lg"
	"os"
	"path"
	"text/template"
	"unicode"
)

type Generator struct {
	FileNameBuilder
	NextTemplateLoader     func() TemplateLoader
	NextTemplateDataLoader func() TemplateDataLoader
}

func (o *Generator) Generate() (err error) {
	templateLoader := o.NextTemplateLoader()
	for templateLoader != nil {
		templateDataLoader := o.NextTemplateDataLoader()
		for templateDataLoader != nil {
			if err = o.resolveAndGenerate(templateLoader, templateDataLoader); err != nil {
				return
			}
			templateDataLoader = o.NextTemplateDataLoader()
		}
		templateLoader = o.NextTemplateLoader()
	}
	return
}

func (o *Generator) resolveAndGenerate(
	templateLoader TemplateLoader, templateDataLoader TemplateDataLoader) (err error) {

	var tmpl *template.Template
	if tmpl, err = templateLoader.LoadTemplate(); err != nil {
		return
	}

	var byteValue []byte
	if byteValue, err = templateDataLoader.LoadData(); err != nil {
		return
	}

	var data interface{}
	if err = json.Unmarshal(byteValue, &data); err != nil {
		return
	}
	var outputFile string
	if outputFile, err = o.BuildFilePathDynamic(
		templateLoader.TemplateSource(), templateDataLoader.DataSource()); err != nil {
		return
	}
	err = generateFile(tmpl, outputFile, data)
	return
}

func generateFile(template *template.Template, outputFileName string, data interface{}) (err error) {
	lg.LOG.Infof("resolveAndGenerate '%v'", outputFileName)
	if err = os.MkdirAll(path.Dir(outputFileName), os.ModePerm); err != nil {
		return
	}

	var file *os.File
	if file, err = os.Create(outputFileName); err != nil {
		return
	}
	defer file.Close()
	err = template.Execute(file, data)
	return
}

func ToJSON(data []byte) ([]byte, error) {
	if hasJSONPrefix(data) {
		return data, nil
	}
	return yaml.YAMLToJSON(data)
}

var jsonPrefix = []byte("{")

func hasJSONPrefix(buf []byte) bool {
	trim := bytes.TrimLeftFunc(buf, unicode.IsSpace)
	return bytes.HasPrefix(trim, jsonPrefix)
}
