package cmd

import (
	"github.com/spf13/pflag"
)

func FlagDataFile(flagSet *pflag.FlagSet, p *string) (flagName string) {
	flagName = "data-file"
	flagSet.StringVarP(p, flagName, "d", "", "json or yaml data file")
	return
}

func FlagTemplateFile(flagSet *pflag.FlagSet, p *string) (flagName string) {
	flagName = "template-file"
	flagSet.StringVarP(p, flagName, "t", "", "go template file")
	return
}

func FlagOutputPath(flagSet *pflag.FlagSet, p *string) (flagName string) {
	flagName = "output-path"
	flagSet.StringVarP(p, flagName, "o", "",
		"output path or relative path to 'output-dir-on-template' or 'output-dir-on-data'")
	return
}

func FlagMulti(flagSet *pflag.FlagSet, p *bool) {
	flagName := "multi"
	flagSet.BoolVarP(p, flagName, "m", false,
		"generates one file for each File object in Json (or Yaml) data file")
	return
}
