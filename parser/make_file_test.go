package parser_test

import (
	"testing"

	"github.com/bytebytebug/newf/parser"
)

func TestHello(t *testing.T) {
	sut := parser.NewMakeFileInputParser()

	files, err := sut.Parse("some/path", []string{"hello/world.txt"})
	if err != nil {
		t.Fatal("parsing error")
	}

	if files[0] != "some/path/hello/world.txt" {
		t.Fatal("\"" + files[0] + "\" should be \"some/path/hello/world.txt\"")
	}
}
