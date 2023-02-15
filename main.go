package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	language "github.com/xhadzi01/ForceMeToLearn/Language"
	"github.com/xhadzi01/ForceMeToLearn/views"
)

func main() {
	http.HandleFunc("/", views.ViewScoreboard)

	for language := range language.GetAllLanguages() {
		languageViewHandler, handlerErr := views.GetLanguageViewHandler(language)
		languageName, nameErr := language.GetName()

		if languageViewHandler == nil {
			fmt.Println("Language view is nil")
		} else if handlerErr != nil {
			fmt.Printf("Language view could not be created. error message: %v \n", error.Error())
		} else if nameErr != nil {
			fmt.Printf("Language name could be retrieved. error message: %v \n", error.Error())
		} else if len(languageName) == 0 {
			fmt.Println("Language name is empty")
		} else {
			http.Handle("/"+languageName, languageViewHandler)
		}
	}

	err := http.ListenAndServe(":12500", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("error: server closed")
		os.Exit(1)
	} else if err != nil {
		fmt.Println("error starting server:", err)
		os.Exit(1)
	}
}
