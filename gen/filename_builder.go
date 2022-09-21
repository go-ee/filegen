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
	OutputPath         string
	RelativeToTemplate bool
	RelativeToData     bool
}

func (o *DefaultsFileNameBuilder) BuildFilePathDynamic(
	templateSource string, templateDataSource string) (ret string, err error) {

	fileName := strings.TrimSuffix(filepath.Base(templateSource), filepath.Ext(templateSource))
	return o.BuildFilePath(templateSource, templateDataSource, fileName)
}

func (o *DefaultsFileNameBuilder) BuildFilePath(
	templateSource string, templateDataSource string, fileName string) (ret string, err error) {

	if o.RelativeToTemplate {
		if o.OutputPath == "" {
			ret = path.Dir(templateDataSource)
		} else {
			ret = filepath.Join(path.Dir(templateSource), o.OutputPath)
		}
	} else if o.RelativeToData {
		if o.OutputPath == "" {
			ret = path.Dir(templateDataSource)
		} else {
			ret = filepath.Join(path.Dir(templateDataSource), o.OutputPath)
		}
	} else {
		ret, err = filepath.Abs(o.OutputPath)
	}
	ret = filepath.Join(ret, fileName)
	return
}
