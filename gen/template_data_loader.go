package gen

import (
	"os"
)

type TemplateDataLoader interface {
	DataLabel() string
	LoadData() ([]byte, error)
}

func NewFileTDataLoader(dataFile string) *FileDataLoader {
	return &FileDataLoader{
		File: dataFile,
	}
}

type FileDataLoader struct {
	File string
}

func (o *FileDataLoader) DataLabel() string {
	return o.File
}

func (o *FileDataLoader) LoadData() (ret []byte, err error) {
	if ret, err = os.ReadFile(o.File); err == nil {
		ret, err = ToJSON(ret)
	}
	return
}

type MultipleFileData struct {
	Files []struct {
		FileName string
		Data     interface{}
	}
}
