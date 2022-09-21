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
	TemplateDataLoader
	TemplateLoader
	HasMultipleFiles bool
}

func (o *Generator) Generate() (err error) {
	var tmpl *template.Template
	if tmpl, err = o.LoadTemplate(); err != nil {
		return
	}

	var byteValue []byte
	if byteValue, err = o.LoadData(); err != nil {
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

	for _, item := range multiFileData.Files {
		var outputFile string
		if outputFile, err = o.BuildFilePath(o.TemplateLabel(), o.DataLabel(), item.FileName); err != nil {
			return
		}

		if err = generateFile(tmpl, outputFile, item.Data); err != nil {
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
	var outputFile string
	if outputFile, err = o.BuildFilePathDynamic(o.TemplateLabel(), o.DataLabel()); err != nil {
		return
	}
	err = generateFile(tmpl, outputFile, data)
	return
}

func generateFile(template *template.Template, outputFileName string, data interface{}) (err error) {
	lg.LOG.Infof("generate '%v'", outputFileName)
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
