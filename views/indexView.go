package views

import (
	"net/http"
	"strings"

	language "github.com/xhadzi01/ForceMeToLearn/Language"
)

type ScoreDataStruct struct {
	LanguageName            string
	TotalScoreNominator     language.LanguageWordScore
	TotalScoreDenominator   language.LanguageWordScore
	TotalScoreRatio         language.LanguageWordsRatio
	CurrentScoreNominator   language.LanguageWordScore
	CurrentScoreDenominator language.LanguageWordScore
	CurrentScoreRatio       language.LanguageWordsRatio
	Games                   map[string]string
}

func ViewScoreboard(res http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		gameType := req.URL.Query().Get("gameType")
		gameType = strings.ToLower(gameType)

		if gameType == "quizgame" {
			http.Redirect(res, req, "/show-quiz-question", http.StatusSeeOther)
		} else if gameType == "translationgame" {
			http.Redirect(res, req, "/show-translation-question", http.StatusSeeOther)
		} else {
			res.WriteHeader(http.StatusNotFound)
		}
		return
	}

	name, _ := lang.GetName()
	totalCorrect, totalTotal, totalRatio, _ := lang.GetTotalScore()
	currentCorrect, currentTotal, currentRatio, _ := lang.GetCurrentScore()

	entrypointTemplate, _ := GetRootedTemplate("entrypoint", "index.html", "scoreboard.html")

	games := make(map[string]string)
	games["Translation game"] = "/quiz-question"
	games["Quiz game"] = "/translation-question"

	ds := &ScoreDataStruct{
		LanguageName:            name,
		TotalScoreNominator:     totalCorrect,
		TotalScoreDenominator:   totalTotal,
		TotalScoreRatio:         totalRatio,
		CurrentScoreNominator:   currentCorrect,
		CurrentScoreDenominator: currentTotal,
		CurrentScoreRatio:       currentRatio,
		Games:                   games,
	}
	entrypointTemplate.Execute(res, ds)

}
