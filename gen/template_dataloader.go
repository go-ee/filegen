package gen

import (
	"io/fs"
	"os"
	"path/filepath"
)

type DataLoader interface {
	DataSource() string
	LoadData() ([]byte, error)
}

func NewJsonFileDataLoader(dataFile string) *JsonFileDataLoader {
	return &JsonFileDataLoader{
		File: dataFile,
	}
}

type JsonFileDataLoader struct {
	File string
}

func (o *JsonFileDataLoader) DataSource() string {
	return o.File
}

func (o *JsonFileDataLoader) LoadData() (ret []byte, err error) {
	if ret, err = os.ReadFile(o.File); err == nil {
		ret, err = ToJSON(ret)
	}
	return
}

func CollectFilesRecursive(baseFile string) (ret []string, err error) {
	ret = []string{}
	fileName := filepath.Base(baseFile)
	baseFolder := filepath.Dir(baseFile)
	err = filepath.Walk(baseFolder, func(path string, info fs.FileInfo, err error) (walkErr error) {
		if !info.IsDir() && info.Name() == fileName {
			ret = append(ret, path)
		}
		return
	})
	return
}

func FilesToTemplateDataLoaders(templateDataFiles []string) (ret []DataLoader) {
	ret = make([]DataLoader, len(templateDataFiles))
	for i, file := range templateDataFiles {
		ret[i] = NewJsonFileDataLoader(file)
	}
	return
}
