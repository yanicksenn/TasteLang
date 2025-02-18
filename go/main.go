package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/yanicksenn/TasteLang/go/shared"
	"github.com/yanicksenn/TasteLang/go/taste_parser"
	"github.com/yanicksenn/TasteLang/go/taste_tokenizer"
	"github.com/yanicksenn/TasteLang/go/taste_writer"
)

var tasteFilePathFlag string
var recipeFilePathFlag string
var outputFilePathFlag string
var overrideOutputFileFlag bool 

func main() {
	flag.StringVar(&tasteFilePathFlag, "taste_file_path", "", "Path to the taste file.")
	flag.StringVar(&recipeFilePathFlag, "recipe_file_path", "", "Path to the recipe file.")
	flag.StringVar(&outputFilePathFlag, "output_file_path", "", "Path to the output file.")
	flag.BoolVar(&overrideOutputFileFlag, "override_output", false, "Whether the output file should be overriden if it already exists.")
	flag.Parse()

	if len(tasteFilePathFlag) == 0 {
		fmt.Fprintln(os.Stderr, errors.New("Flag --taste_file_path not set"))
		os.Exit(1)
	}

	tasteFile, err := read(tasteFilePathFlag)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}

	if len(recipeFilePathFlag) == 0 {
		fmt.Fprintln(os.Stderr, errors.New("Flag --recipe_file_path not set"))
		os.Exit(3)
	}

	recipeFile, err := read(recipeFilePathFlag)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(4)
	}

	if len(outputFilePathFlag) == 0 {
		fmt.Fprintln(os.Stderr, errors.New("Flag --output_file_path not set"))
		os.Exit(5)
	}

	tokens := taste_tokenizer.Tokenize(tasteFile)
	file, err := taste_parser.Parse(tokens)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(6)
	}

	recipe := shared.Recipe{
		Content: recipeFile,
	}
	err = taste_writer.Write(file, &recipe, outputFilePathFlag, overrideOutputFileFlag)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(7)
	}
}

func read(path string) (string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

