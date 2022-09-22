package gen

import (
	"os"
)

type TemplateDataLoader interface {
	DataSource() string
	LoadData() ([]byte, error)
}

func NewFileDataLoader(dataFile string) *FileDataLoader {
	return &FileDataLoader{
		File: dataFile,
	}
}

type FileDataLoader struct {
	File string
}

func (o *FileDataLoader) DataSource() string {
	return o.File
}

func (o *FileDataLoader) LoadData() (ret []byte, err error) {
	if ret, err = os.ReadFile(o.File); err == nil {
		ret, err = ToJSON(ret)
	}
	return
}

type SingleNextProvider[T any] struct {
	Item             T
	alreadyDelivered bool
}

func (o *SingleNextProvider[T]) Next() (ret T) {
	if !o.alreadyDelivered {
		o.alreadyDelivered = true
		ret = o.Item
	}
	return
}
