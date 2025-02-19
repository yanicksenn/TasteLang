package parser

import (
	"reflect"
	"testing"

	"github.com/yanicksenn/TasteLang/go/shared"
	"github.com/yanicksenn/TasteLang/go/taste_parser"
	"github.com/yanicksenn/TasteLang/go/taste_tokenizer"
)

func TestParse(t *testing.T) {
	content := `

namespace Tokenizer.Test;

type Token {
	string value;
	int counter;
}

	`
	actualFile, err := taste_parser.Parse(taste_tokenizer.Tokenize(content))
	if err != nil {
		t.Fatalf("%w", err)
	}
	expectedFile := &shared.File{
		Namespace: "Tokenizer.Test",
		Types: []shared.Type{
			{
				Name: "Token",
				Fields: []shared.Field{
					{
						Name: "value",
						Type: "string",
						IsArray: false,
					},
					{
						Name: "counter",
						Type: "int",
						IsArray: false,
					},
				},
			},
		},

	}
	if !reflect.DeepEqual(expectedFile, actualFile) {
		t.Fatalf("Expected: %+v but got: %+v\n", expectedFile, actualFile)
	}		
}
