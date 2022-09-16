package filegen

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Masterminds/sprig"
	"github.com/ghodss/yaml"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"
	"unicode"
)

type MultipleFileData struct {
	Files []struct {
		FileName string
		Data     interface{}
	}
}

type TemplateData interface {
	Label() string
	LoadJsonData() ([]byte, error)
}

type FileDataLoader struct {
	FileName string
}

func (o *FileDataLoader) Label() string {
	return o.FileName
}

func (o *FileDataLoader) LoadJsonData() (ret []byte, err error) {
	if ret, err = os.ReadFile(o.FileName); err == nil {
		ret, err = ToJSON(ret)
	}
	return
}

type Generator struct {
	DataLoader       TemplateData
	TemplateFileName string
	OutputDirectory  string
	HasMultipleFiles bool
}

func (o *Generator) Generate() (err error) {
	var tmpl *template.Template
	if tmpl, err = tmpl.New(path.Base(o.TemplateFileName)).Funcs(FuncMap).Funcs(sprig.TxtFuncMap()).ParseFiles(o.TemplateFileName); err != nil {
		return
	}

	var byteValue []byte
	if byteValue, err = o.DataLoader.LoadJsonData(); err != nil {
		return
	}

	if o.HasMultipleFiles {
		var multiFile MultipleFileData

		err = json.Unmarshal(byteValue, &multiFile)
		if err != nil {
			fmt.Println(err)
		}

		for _, mfile := range multiFile.Files {
			outputFileName := path.Join(path.Dir(o.DataLoader.Label()), mfile.FileName)
			mdata := mfile.Data
			generateFile(tmpl, o.OutputDirectory, outputFileName, mdata)
		}

	} else {
		var data interface{}
		err = json.Unmarshal(byteValue, &data)
		if err != nil {
			fmt.Println(err)
		}
		outputFileName := strings.TrimSuffix(o.DataLoader.Label(), filepath.Ext(o.DataLoader.Label())) + ".generated.txt"
		generateFile(tmpl, o.OutputDirectory, outputFileName, data)
	}
	return
}

func generateFile(template *template.Template, outputDirectory string, outputFileName string, data interface{}) {
	absOutputFileName := path.Join(outputDirectory, outputFileName)
	os.MkdirAll(path.Dir(absOutputFileName), os.ModePerm)
	outputFile, err := os.Create(absOutputFileName)
	fmt.Println("Generating file : " + absOutputFileName)
	defer outputFile.Close()
	if err != nil {
		fmt.Println(err)
	}
	err = template.Execute(outputFile, data)
	if err != nil {
		fmt.Println(err)
	}
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
