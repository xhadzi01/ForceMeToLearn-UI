package views

import (
	"net/http"
	"strconv"

	language "github.com/xhadzi01/ForceMeToLearn/Language"
)

type QuizTranslationWord struct {
	Word string
	Hash uint32
}

type QuizDataStruct struct {
	LanguageName            string
	TotalScoreNominator     language.LanguageWordScore
	TotalScoreDenominator   language.LanguageWordScore
	TotalScoreRatio         language.LanguageWordsRatio
	CurrentScoreNominator   language.LanguageWordScore
	CurrentScoreDenominator language.LanguageWordScore
	CurrentScoreRatio       language.LanguageWordsRatio
	Word_to_translate       string
	Possible_translations   []QuizTranslationWord
}

func QuizQuestion(res http.ResponseWriter, req *http.Request) {

	if req.Method == http.MethodPost {
		answerHash := req.URL.Query().Get("answerHash")
		hashVal, _ := strconv.Atoi(answerHash)
		quessingGameHandler.Answer(game.LanguageWordHash(hashVal))
	}

	name, _ := lang.GetName()
	totalCorrect, totalTotal, totalRatio, _ := lang.GetTotalScore()
	currentCorrect, currentTotal, currentRatio, _ := lang.GetCurrentScore()
	wordPack, _ := lang.GetNewWordPack()

	possibleTranslations := make([]QuizTranslationWord, 0)
	for _, incorrectWord := range wordPack.IncorrectWords {
		possibleTranslations = append(possibleTranslations, QuizTranslationWord{Word: incorrectWord.Word, Hash: uint32(incorrectWord.Hash)})
	}
	possibleTranslations = append(possibleTranslations, QuizTranslationWord{Word: wordPack.CorrectWord.Word, Hash: uint32(wordPack.CorrectWord.Hash)})

	entrypointTemplate, _ := GetRootedTemplate("entrypoint", "index.html", "quiz_template.html")
	ds := &QuizDataStruct{
		LanguageName:            name,
		TotalScoreNominator:     totalCorrect,
		TotalScoreDenominator:   totalTotal,
		TotalScoreRatio:         totalRatio,
		CurrentScoreNominator:   currentCorrect,
		CurrentScoreDenominator: currentTotal,
		CurrentScoreRatio:       currentRatio,
		Word_to_translate:       wordPack.AskedForWord.Word,
		Possible_translations:   possibleTranslations,
	}
	entrypointTemplate.Execute(res, ds)

}
