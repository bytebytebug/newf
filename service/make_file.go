package service

import (
	"github.com/bytebytebug/newf/parser"
	"github.com/bytebytebug/newf/utils"
)

type InputParser interface {
	Parse(base string, files []string) ([]string, error)
}

type FS interface {
	Create(file string) error
}

type MakeFileService struct {
	inputParser InputParser
	fs          FS
}

type OptsValues struct {
	InputParser InputParser
	FS          FS
}

func DefaultOptValues() *OptsValues {
	return &OptsValues{
		InputParser: parser.NewMakeFileInputParser(),
		FS:          utils.CreateFs(),
	}
}

type Opt func(v *OptsValues)

func CreateMakeFileService(opts ...Opt) (*MakeFileService, error) {

	optsValues := DefaultOptValues()

	for _, opt := range opts {
		opt(optsValues)
	}

	return &MakeFileService{
		inputParser: optsValues.InputParser,
		fs:          optsValues.FS,
	}, nil
}

func (service *MakeFileService) Exec(base string, files []string) error {

	var err error

	var parsedFileNames []string
	parsedFileNames, err = service.inputParser.Parse(base, files)
	if err != nil {
		return err
	}

	for _, parsedFileName := range parsedFileNames {
		err := service.fs.Create(parsedFileName)
		if err != nil {
			return err
		}
	}

	return nil
}
