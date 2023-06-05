package gen

import (
	"path"
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
	templateSource string, templateDataSource string) (ret string, err error) {

	fileName := strings.TrimSuffix(filepath.Base(templateSource), filepath.Ext(templateSource))
	return o.BuildFilePath(templateSource, templateDataSource, fileName)
}

func (o *DefaultsFileNameBuilder) BuildFilePath(
	templateSource string, templateDataSource string, fileName string) (ret string, err error) {

	if o.RelativeToTemplate {
		ret = path.Dir(templateSource)
		if o.RelativePathOrFullPath != "" {
			ret = filepath.Join(ret, o.RelativePathOrFullPath)
		}
	} else if o.RelativeToData {
		ret = path.Dir(templateDataSource)
		if o.RelativePathOrFullPath != "" {
			ret = filepath.Join(ret, o.RelativePathOrFullPath)
		}
	} else {
		ret, err = filepath.Abs(o.RelativePathOrFullPath)
	}
	ret = filepath.Join(ret, fileName)
	return
}
