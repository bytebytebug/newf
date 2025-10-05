package parser

import "path/filepath"

type MakeFileInputParser struct{}

func NewMakeFileInputParser() *MakeFileInputParser {
	return &MakeFileInputParser{}
}

func (parser *MakeFileInputParser) Parse(base string, files []string) ([]string, error) {
	var out []string = make([]string, 0)

	for _, file := range files {
		if filepath.IsAbs(file) {
			out = append(out, file)
		} else {
			out = append(out, filepath.Join(base, file))
		}
	}

	return out, nil
}
