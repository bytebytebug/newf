package service_test

import (
	"errors"
	"testing"

	"github.com/bytebytebug/newf/service"
	"github.com/bytebytebug/newf/utils"
)

type FSMock struct {
	CreateFn func(file string) error
}

func (fs *FSMock) Create(file string) error {
	return fs.CreateFn(file)
}

type ParserMock struct {
	ParseFn func(base string, files []string) ([]string, error)
}

func (p *ParserMock) Parse(base string, files []string) ([]string, error) {
	return p.ParseFn(base, files)
}

func TestShouldCreateAFile(t *testing.T) {
	var err error
	var sut *service.MakeFileService

	createdFile := ""

	// Opts
	var WithMockedFS service.Opt = func(v *service.OptsValues) {
		v.FS = &FSMock{
			CreateFn: func(file string) error {
				createdFile = file
				return nil
			},
		}
	}
	var WithMockedParser service.Opt = func(v *service.OptsValues) {
		v.InputParser = &ParserMock{
			ParseFn: func(base string, files []string) ([]string, error) {
				newFiles := utils.Map(files, func(file string) string {
					return base + "/" + file
				})
				return newFiles, nil
			},
		}
	}

	if sut, err = service.CreateMakeFileService(
		WithMockedFS,
		WithMockedParser,
	); err != nil {
		t.Fatal(err)
	}

	if err = sut.Exec(".", []string{"hello"}); err != nil {
		t.Fatal(err)
	}

	if createdFile != "./hello" {
		t.Fatal("wrong file created")
	}
}

func TestShouldFailIfFiledIsntCreated(t *testing.T) {
	var err error
	var sut *service.MakeFileService

	var WithMockedFS service.Opt = func(v *service.OptsValues) {
		v.FS = &FSMock{
			CreateFn: func(file string) error {
				return errors.New("errors")
			},
		}
	}

	var WithMockedParser service.Opt = func(v *service.OptsValues) {
		v.InputParser = &ParserMock{
			ParseFn: func(base string, files []string) ([]string, error) {
				newFiles := utils.Map(files, func(file string) string {
					return base + "/" + file
				})
				return newFiles, nil
			},
		}
	}

	if sut, err = service.CreateMakeFileService(
		WithMockedFS,
		WithMockedParser,
	); err != nil {
		t.Fatal(err)
	}

	if err = sut.Exec("./my-dir", []string{"file.txt"}); err == nil {
		t.Fatal("err is nil")
	}
}

func TestShouldReturnErrorIfParsingFail(t *testing.T) {
	var err error
	var sut *service.MakeFileService

	var WithMockedFS service.Opt = func(v *service.OptsValues) {
		v.FS = &FSMock{
			CreateFn: func(file string) error {
				panic(1)
			},
		}
	}

	var WithMockedParser service.Opt = func(v *service.OptsValues) {
		v.InputParser = &ParserMock{
			ParseFn: func(base string, files []string) ([]string, error) {
				return []string{}, errors.New("some error")
			},
		}
	}

	if sut, err = service.CreateMakeFileService(
		WithMockedFS,
		WithMockedParser,
	); err != nil {
		t.Fatal(err)
	}

	if err = sut.Exec("./my-dir", []string{"file.txt"}); err == nil {
		t.Fatal("err is nil")
	}
}
