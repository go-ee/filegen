package gen

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ghodss/yaml"
	"os"
	"path"
	"text/template"
	"unicode"
)

type Generator struct {
	OutputFileNameBuilder
	TemplateDataLoader
	TemplateFactory
	HasMultipleFiles bool
}

func (o *Generator) Generate() (err error) {
	var tmpl *template.Template
	if tmpl, err = o.CreateTemplate(); err != nil {
		return
	}

	var byteValue []byte
	if byteValue, err = o.LoadTemplateData(); err != nil {
		return
	}

	if o.HasMultipleFiles {
		err = o.generateMultipleFiles(tmpl, byteValue)
	} else {
		err = o.generateSingleFile(tmpl, byteValue)
	}
	return
}

func (o *Generator) generateMultipleFiles(tmpl *template.Template, byteValue []byte) (err error) {
	var multiFileData MultipleFileData

	if err = json.Unmarshal(byteValue, &multiFileData); err != nil {
		return
	}

	for _, outputFile := range multiFileData.Files {
		data := outputFile.Data
		outputFileName := o.BuildOutputFileNameForFileName(o.TemplateLabel(), outputFile.FileName)
		if err = generateFile(tmpl, outputFileName, data); err != nil {
			return
		}
	}
	return
}

func (o *Generator) generateSingleFile(tmpl *template.Template, byteValue []byte) (err error) {
	var data interface{}
	if err = json.Unmarshal(byteValue, &data); err != nil {
		return
	}
	outputFileName := o.BuildOutputFileName(o.TemplateLabel())
	err = generateFile(tmpl, outputFileName, data)
	return
}

func generateFile(template *template.Template, outputFileName string, data interface{}) (err error) {
	if err = os.MkdirAll(path.Dir(outputFileName), os.ModePerm); err != nil {
		return
	}

	var file *os.File
	if file, err = os.Create(outputFileName); err != nil {
		return
	}
	defer file.Close()
	if err != nil {
		fmt.Println(err)
	}
	err = template.Execute(file, data)
	if err != nil {
		fmt.Println(err)
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
