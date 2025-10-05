package service

import "fmt"

type MakeFileService struct {
}

func CreateMakeFileService() (*MakeFileService, error) {
	return nil, nil
}

func (service *MakeFileService) Exec(base string, files []string) error {
	fmt.Println(base, files)
	return nil
}
