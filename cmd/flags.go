package cmd

import (
	"github.com/spf13/pflag"
)

func FlagDataFile(flagSet *pflag.FlagSet, p *string) (flagName string) {
	flagName = "data"
	flagSet.StringVarP(p, flagName, "d", "", "json or yaml data file")
	return
}

func FlagDataFileRecursive(flagSet *pflag.FlagSet, p *bool) (flagName string) {
	flagName = "data-recursive"
	flagSet.BoolVarP(p, flagName, "", dataFileRecursive, "look for the same data file recursive")
	return
}

func FlagMacrosTemplatesFiles(flagSet *pflag.FlagSet, p *[]string) (flagName string) {
	flagName = "macros"
	flagSet.StringSliceVarP(p, flagName, "m", nil,
		"macros, template files reused by main templates (semicolon separated or multiple flags)")
	return
}

func FlagTemplateFiles(flagSet *pflag.FlagSet, p *[]string) (flagName string) {
	flagName = "templates"
	flagSet.StringSliceVarP(p, flagName, "t", nil,
		"template files (semicolon separated or multiple flags)")
	return
}

func FlagOutputPath(flagSet *pflag.FlagSet, p *string) (flagName string) {
	flagName = "output"
	flagSet.StringVarP(p, flagName, "o", "",
		"output path or relative path to 'output-dir-on-template' or 'output-dir-on-data'")
	return
}
