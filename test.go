package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	language "github.com/xhadzi01/ForceMeToLearn/Language"
	game "github.com/xhadzi01/ForceMeToLearn/Language/games"
	"github.com/xhadzi01/ForceMeToLearn/views"
)

type ViewNewQuizQuestionHandler struct {
	lang *language.Language
	game game.IGuessingGame
}

var quessingGameHandler *ViewNewQuizQuestionHandler

func (view *ViewNewQuizQuestionHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	currentCorrect, currentTotal, currentRatio, _ := quessingGameHandler.lang.GetCurrentScore()
	views.QuizQuestion(res, req)
}

func maind() {
	langInst := language.NewLanguage("English", &language.DatabaseInst)
	gameInst, err := quessingGameHandler.lang.GetNewGuessingGame()

	if langInst != nil && gameInst != nil && err == nil {

		quessingGameHandler := &ViewNewQuizQuestionHandler{
			lang: langInst,
			game: gameInst,
		}

		if quessingGameHandler == nil {
			fmt.Println("Guessing game handler is invalid")
		}

		http.HandleFunc("/", views.ViewScoreboard)
		http.Handle("/quiz-question", quessingGameHandler)
		// http.HandleFunc("/translation-question", views.TranslationQuestion)

		err := http.ListenAndServe(":12500", nil)
		if errors.Is(err, http.ErrServerClosed) {
			fmt.Println("server closed")
		} else if err != nil {
			fmt.Println("error starting server:", err)
			os.Exit(1)
		}
	} else {
		fmt.Println("Could not create Language or game instance")
	}
}
