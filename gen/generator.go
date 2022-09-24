package gen

import (
	"bytes"
	"encoding/json"
	"github.com/ghodss/yaml"
	"github.com/go-ee/utils/lg"
	"os"
	"path"
	"strings"
	"text/template"
	"unicode"
)

type Generator struct {
	FileNameBuilder
	NextTemplateLoader     NextProvider[TemplateLoader]
	NextTemplateDataLoader NextProvider[DataLoader]
}

func (o *Generator) Generate() (err error) {
	templateDataLoader := o.NextTemplateDataLoader.Next()
	for templateDataLoader != nil {
		templateLoader := o.NextTemplateLoader.Next()
		for templateLoader != nil {
			if err = o.resolveAndGenerate(templateLoader, templateDataLoader); err != nil {
				return
			}
			templateLoader = o.NextTemplateLoader.Next()
		}
		templateDataLoader = o.NextTemplateDataLoader.Next()
		o.NextTemplateLoader.Reset()
	}
	return
}

func (o *Generator) resolveAndGenerate(
	templateLoader TemplateLoader, templateDataLoader DataLoader) (err error) {

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

func generateFile(tmpl *template.Template, outputFileName string, data interface{}) (err error) {
	if err = os.MkdirAll(path.Dir(outputFileName), os.ModePerm); err != nil {
		return
	}

	perm := os.FileMode(0666)
	if strings.HasSuffix(outputFileName, ".sh") {
		perm |= 0111
	}

	var file *os.File
	if file, err = os.OpenFile(outputFileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, perm); err != nil {
		return
	}
	defer file.Close()

	lg.LOG.Infof("generate: %v", outputFileName)
	if err = tmpl.Execute(file, data); err != nil {
		return
	}

	if strings.HasSuffix(outputFileName, ".sh") {
		var info os.FileInfo
		if info, err = os.Stat(outputFileName); err != nil {
			return
		}
		mode := info.Mode()
		if mode != perm {
			err = os.Chmod(outputFileName, perm)
		}
	}
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
