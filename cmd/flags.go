package cmd

import "github.com/spf13/cobra"

func FlagDataFile(command *cobra.Command, p *string, required bool) {
	flagName := "dataFile"
	command.Flags().StringVarP(p, flagName, "d", "", "json or yaml data file")
	markFlagRequiredIfTrue(command, flagName, required)
}

func FlagTemplateFile(command *cobra.Command, p *string, required bool) {
	flagName := "templateFile"
	command.Flags().StringVarP(p, flagName, "t", "", "go template file")
	markFlagRequiredIfTrue(command, flagName, required)
}

func FlagOutputDir(command *cobra.Command, p *string, required bool) {
	flagName := "outputDir"
	command.Flags().StringVarP(p, flagName, "o", ".", "output directory")
	markFlagRequiredIfTrue(command, flagName, required)
}

func FlagOutputFileExt(command *cobra.Command, p *string, required bool) {
	flagName := "outputFileExt"
	command.Flags().StringVarP(p, flagName, "e", "txt", "file extension for generated file")
	markFlagRequiredIfTrue(command, flagName, required)
}

func FlagMulti(command *cobra.Command, p *bool, required bool) {
	flagName := "multi"
	command.Flags().BoolVarP(p, flagName, "m", false,
		"generates one file for each File object in Json (or Yaml) data file")
	markFlagRequiredIfTrue(command, flagName, required)
}

func markFlagRequiredIfTrue(command *cobra.Command, flagName string, required bool) {
	if required {
		_ = command.MarkFlagRequired(flagName)
	}
}
