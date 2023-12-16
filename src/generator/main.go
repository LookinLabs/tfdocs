package generator

import (
	OptionsData "lhv_docs_generator/m/v2/structures/command_line"
	"log"
	"os"
	"strings"
)

func DocsGenerator() {
	flags := .ParseFlags()
	path := flags.Path

	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
		log.Printf("Unable to receive path. Directory or input files not found.")
	}

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".tf") {
			content, err := os.ReadFile(path + file.Name())
			if err != nil {
				log.Fatal(err)
				log.Printf("Unable to read file %s", file.Name())
			}

			// parse the file content to extract the necessary information
			// this is a placeholder and should be replaced with your actual parsing logic
			parsedContent := parseFileContent(string(content))

			// format the parsed content into a Markdown format
			formattedContent := formatContent(parsedContent)

			// write the formatted content to a new README file
			err = os.WriteFile(path+"README.md", []byte(formattedContent), 0640)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func parseFileContent(content string) string {
	// placeholder function, replace with your actual parsing logic
	return content
}

func formatContent(content string) string {
	// placeholder function, replace with your actual formatting logic
	return "# " + content
}
