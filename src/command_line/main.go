package command_line

import (
	"encoding/json"
	flagsData "lhv_docs_generator/m/v2/structures/command_line"
	"log"
	"os"

	"github.com/spf13/pflag"
)

func InitFlags() (flagsData.Flags, map[string]*string) {
	flags := flagsData.Flags{
		DocType:   new(string),
		Path:      new(string),
		OutputDir: new(string),
		OneFile:   false,
	}

	// Map the flag names to the pointers of the corresponding fields in the flags struct
	flagMap := map[string]*string{
		"type":       flags.DocType,
		"path":       flags.Path,
		"output-dir": flags.OutputDir,
	}

	return flags, flagMap
}

func ParseFlags() flagsData.Flags {
	flags, flagMap := InitFlags()

	// Read the flag details from the JSON file
	file, err := os.Open("flagDetails.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	var flagDetails []flagsData.FlagDetails
	err = decoder.Decode(&flagDetails)
	if err != nil {
		log.Fatal(err)
	}

	// Parse the flags
	for _, detail := range flagDetails {
		if detail.Name == "one-file" {
			pflag.BoolVar(&flags.OneFile, detail.Name, detail.Default == "true", detail.Description)
		} else {
			pflag.StringVar(flagMap[detail.Name], detail.Name, detail.Default, detail.Description)
		}
	}

	pflag.Parse()

	return flags
}
