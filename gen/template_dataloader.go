package gen

import (
	"github.com/go-ee/utils/lg"
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
	fileName := filepath.Base(baseFile)
	baseFolder := filepath.Dir(baseFile)
	m := map[string]bool{}
	lg.LOG.Infof("collect '%v' files recursive in '%v'", fileName, baseFolder)
	err = filepath.Walk(baseFolder, func(path string, info fs.FileInfo, err error) (walkErr error) {
		if err == nil && !info.IsDir() && info.Name() == fileName {
			m[info.Name()] = true
		}
		return
	})
	i := 0
	ret = make([]string, len(m))
	for path := range m {
		ret[i] = path
		i++
	}
	lg.LOG.Infof("%v - '%v' files collected recursive in '%v'", len(ret), fileName, baseFolder)
	return
}

func FilesToTemplateDataLoaders(templateDataFiles []string) (ret []DataLoader) {
	ret = make([]DataLoader, len(templateDataFiles))
	for i, file := range templateDataFiles {
		ret[i] = NewJsonFileDataLoader(file)
	}
	return
}
