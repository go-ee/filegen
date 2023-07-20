package gen

import (
	"path/filepath"
	"strings"
)

type FileNameBuilder interface {
	BuildFilePathDynamic(templateSource string, templateDataSource string) (string, error)
	BuildFilePath(templateSource string, templateDataSource string, fileName string) (string, error)
}

type DefaultsFileNameBuilder struct {
	RelativePathOrFullPath string
	RelativeToTemplate     bool
	RelativeToData         bool
}

func (o *DefaultsFileNameBuilder) BuildFilePathDynamic(
	templateSource string, dataSource string) (ret string, err error) {

	fileName := strings.TrimSuffix(filepath.Base(templateSource), filepath.Ext(templateSource))
	return o.BuildFilePath(templateSource, dataSource, fileName)
}

func (o *DefaultsFileNameBuilder) BuildFilePath(
	templateSource string, dataSource string, fileName string) (ret string, err error) {
	if o.RelativeToTemplate {
		ret = filepath.Dir(templateSource)
		if o.RelativePathOrFullPath != "" {
			ret = filepath.Join(ret, o.RelativePathOrFullPath)
		}
	} else if o.RelativeToData {
		ret = filepath.Dir(dataSource)
		if o.RelativePathOrFullPath != "" {
			ret = filepath.Join(ret, o.RelativePathOrFullPath)
		}
	} else {
		ret, err = filepath.Abs(o.RelativePathOrFullPath)
	}
	ret = filepath.Join(ret, fileName)
	return
}
