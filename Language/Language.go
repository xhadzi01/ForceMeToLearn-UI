package language

import (
	"errors"

	game "github.com/xhadzi01/ForceMeToLearn/Language/games"
)

type Language struct {
	LanguageScore
	name     string
	database ILanguageDatabase
}

func NewLanguage(name string, langData ILanguageDatabase) *Language {
	lang := &Language{}
	lang.current = NewScoreWordPair()
	lang.total = NewScoreWordPair()
	lang.name = name
	lang.database = langData
	return lang
}

func (lang *Language) GetName() (name string, err error) {

	if lang == nil {
		err = errors.New("language object is invalid")
	} else {
		name = lang.name
	}
	return
}

func (lang *Language) GetTotalScore() (correct LanguageWordScore, total LanguageWordScore, ratio LanguageWordsRatio, err error) {
	if lang == nil {
		correct = 0
		total = 0
		ratio = 0.0
		err = errors.New("language object is invalid")
	} else {
		correct = lang.total.correct
		total = lang.total.total
		ratio = lang.total.ratio
	}

	return
}

func (lang *Language) GetCurrentScore() (correct LanguageWordScore, total LanguageWordScore, ratio LanguageWordsRatio, err error) {

	if lang == nil {
		correct = 0
		total = 0
		ratio = 0.0
		err = errors.New("language object is invalid")
	} else {
		correct = lang.current.correct
		total = lang.current.total
		ratio = lang.current.ratio
	}

	return
}

func (lang *Language) ResetScore() (err error) {
	if lang == nil {
		err = errors.New("language object is invalid")
	} else {
		lang.current = NewScoreWordPair()
	}

	return
}

func (lang *Language) GetNewGuessingGame() (retValGame game.IGuessingGame, err error) {
	if lang == nil {
		err = errors.New("language object is invalid")
	} else {
		getNewWordFunctorInst := func() (wordPair game.TranslatedWordsPair, err error) {
			word, err := lang.database.GetRandomWordWithResult()

			wordPair.Word = game.NewLanguageWord(word.word)
			wordPair.Translated = game.NewLanguageWord(word.translatedWord)
			return
		}
		gameInst, errTmp := game.NewLanguageGuessingGame(getNewWordFunctorInst, uint8(5))
		if errTmp != nil {
			err = errTmp
		}
		retValGame = gameInst
	}
	return
}

func (lang *Language) GetNewTranslationGame() (retValGame game.ISimpleTranslationGame, err error) {
	if lang == nil {
		err = errors.New("language object is invalid")
	} else {
		getNewWordFunctorInst := func() (wordPair game.TranslatedWordsPair, err error) {
			word, err := lang.database.GetRandomWordWithResult()

			wordPair.Word = game.NewLanguageWord(word.word)
			wordPair.Translated = game.NewLanguageWord(word.translatedWord)
			return
		}
		gameInst, errTmp := game.NewSimpleTranslationGame(getNewWordFunctorInst)
		if errTmp != nil {
			err = errTmp
		}
		retValGame = gameInst
	}
	return
}

var AllLanguages []*Language

func init() {
	AllLanguages = make([]*Language, 0)

	AllLanguages = append(AllLanguages, NewLanguage("English", DatabaseInst))
}

func GetAllLanguages() []*Language {
	return AllLanguages
}
