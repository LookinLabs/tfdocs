package main

import (
	"lhv_docs_generator/m/v2/src/command_line"
	"lhv_docs_generator/m/v2/src/generator"

	"github.com/spf13/cobra"
)

func main() {
	flags, flagMap := command_line.InitFlags()

	cmd := &cobra.Command{
		Use:   "docs-generate",
		Short: "Generate docs",
		Long:  `This command generates docs.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Call the GenerateDocs function with the flags
			generator.DocsGenerator() // call the DocsGenerator function
		},
	}

	// Add the string flags to the command
	for name, pointer := range flagMap {
		defaultValue := ""
		if name == "output" {
			defaultValue = "./docs"
		}
		cmd.Flags().StringVar(pointer, name, defaultValue, "Description for "+name)
	}

	// Add the boolean flag to the command
	cmd.PersistentFlags().BoolVar(&flags.OneFile, "one-file", false, "Generate all docs in one file")

	cmd.Execute()
}
