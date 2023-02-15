package views

import (
	"errors"
	"html/template"
	"os"
	"path/filepath"

	language "github.com/xhadzi01/ForceMeToLearn/Language"
)

func GetRootedTemplate(entrypoint string, templateNames ...string) (template *template.Template, err error) {
	wd, err := os.Getwd()
	if err != nil {
		return
	}

	expandedPaths := make([]string, len(templateNames))
	for idx, fileName := range templateNames {
		expandedPaths[idx] = filepath.Join(wd, "templates", fileName)
	}

	templateTmp, err := template.ParseFiles(expandedPaths...)
	if err != nil {
		return
	}

	templateLookedUp := templateTmp.Lookup(entrypoint)
	if templateLookedUp == nil {
		err = errors.New("could not retrive the correct template entrypoint")
		return
	}

	template = templateLookedUp
	return
}

func GetLanguageViewHandler(lang language.ILanguage) {

}
