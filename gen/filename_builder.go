package gen

import (
	"github.com/go-ee/utils/lg"
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
	templateSource string, dataSource string) (ret string, err error) {

	fileName := strings.TrimSuffix(filepath.Base(templateSource), filepath.Ext(templateSource))
	return o.BuildFilePath(templateSource, dataSource, fileName)
}

func (o *DefaultsFileNameBuilder) BuildFilePath(
	templateSource string, dataSource string, fileName string) (ret string, err error) {
	if o.RelativeToTemplate {
		ret = path.Dir(templateSource)
		if o.RelativePathOrFullPath != "" {
			ret = filepath.Join(ret, o.RelativePathOrFullPath)
		}
	} else if o.RelativeToData {
		ret = path.Dir(dataSource)
		if o.RelativePathOrFullPath != "" {
			ret = filepath.Join(ret, o.RelativePathOrFullPath)
		}
	} else {
		ret, err = filepath.Abs(o.RelativePathOrFullPath)
	}
	ret = filepath.Join(ret, fileName)
	lg.LOG.Debugf("BuildFilePath: ret=%v, templateSource=%v, dataSource=%v, fileName=%v, RelativeToTemplate=%v, RelativeToData=%v, RelativePathOrFullPath=%v", ret, templateSource, dataSource, fileName, o.RelativeToTemplate, o.RelativeToData, o.RelativePathOrFullPath)
	return
}
