/*
Copyright © 2022 Eugen Eisler <eoeisler@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"github.com/go-ee/filegen/gen"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var cfgFile string
var dataPath string
var dataFileName string
var dataFileRecursive = false
var templateFiles []string
var macrosTemplateFiles []string
var outPath string
var outRelativeToData = false
var outRelativeToTemplate = false

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "filegen",
	Short: "File generation based on Go templateFiles",

	RunE: func(cmd *cobra.Command, args []string) (err error) {
		return generate()
	},
}

func generate() (err error) {
	var templateProvider *gen.NextTemplateProvider
	if templateProvider, err = gen.NewNextTemplateProvider(templateFiles, macrosTemplateFiles); err != nil {
		return
	}

	var dataFiles []string
	if dataFiles, err = collectDataFiles(); err != nil {
		return
	}
	templateDataProvider := &gen.ArrayNextProvider[gen.DataLoader]{
		Items: gen.FilesToTemplateDataLoaders(dataFiles),
	}
	generator := &gen.Generator{
		FileNameBuilder: &gen.DefaultsFileNameBuilder{
			OutputPath: outPath, RelativeToTemplate: outRelativeToTemplate, RelativeToData: outRelativeToData},
		NextTemplateLoader:     templateProvider,
		NextTemplateDataLoader: templateDataProvider,
	}
	err = generator.Generate()
	return
}

func collectDataFiles() (ret []string, err error) {
	if dataFileRecursive {
		ret, err = gen.CollectFilesRecursive(dataPath)
	} else {
		ret = []string{dataPath}
	}
	return
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.filegen.yaml)")

	_ = rootCmd.MarkPersistentFlagRequired(
		FlagDataPath(rootCmd.PersistentFlags(), &dataPath))
	_ = rootCmd.MarkPersistentFlagRequired(
		FlagTemplateFiles(rootCmd.PersistentFlags(), &templateFiles))

	FlagMacrosTemplatesFiles(rootCmd.PersistentFlags(), &macrosTemplateFiles)
	FlagDataFileRecursive(rootCmd.PersistentFlags(), &dataFileRecursive)

	FlagOutputPath(rootCmd.PersistentFlags(), &outPath)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".filegen" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".filegen")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
