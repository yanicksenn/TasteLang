package taste_writer

import (
	"errors"
	"os"
	"text/template"

	"github.com/yanicksenn/TasteLang/go/shared"
)


func Write(file *shared.File, recipe *shared.Recipe, outputFilePath string, overrideOutputFile bool) (error) {
	if file == nil {
		return errors.New("file must not be nil")
	}

	if recipe == nil {
		return errors.New("recipe must not be nil")
	}

	template, err := template.New("t").Parse(recipe.Content)
	if err != nil {
		return errors.New("Failed to parse recipe")
	}

	_, err = os.Stat(outputFilePath)
	if err == nil {
			if !overrideOutputFile {
				return errors.New("Output file already exists and cannot be overridden")
			}
		
			if overrideOutputFile {
				os.Remove(outputFilePath)
			}
	}

	outputFile, err := os.Create(outputFilePath)
	err = template.Execute(outputFile, file)
	if err != nil {
		return errors.New("Failed to build recipe")
	}
	
	return nil
}
