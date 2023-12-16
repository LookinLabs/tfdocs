package OptionsData

type OptionDetails struct {
	Name        string
	Default     string
	Description string
}

type Options struct {
	DocType   *string
	Path      *string
	OutputDir *string
	OneFile   bool
}
