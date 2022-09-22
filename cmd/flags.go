package cmd

import (
	"github.com/spf13/pflag"
)

func FlagDataFile(flagSet *pflag.FlagSet, p *string) (flagName string) {
	flagName = "data-file"
	flagSet.StringVarP(p, flagName, "d", "", "json or yaml data file")
	return
}

func FlagDataFileRecursive(flagSet *pflag.FlagSet, p *bool) (flagName string) {
	flagName = "data-file-recursive"
	flagSet.BoolVarP(p, flagName, "", dataFileRecursive, "look for the same data file recursive")
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
