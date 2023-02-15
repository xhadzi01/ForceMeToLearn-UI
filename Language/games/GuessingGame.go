package game

import (
	"errors"
	"math/rand"
)

type GuessingGame struct {
	getNewWordFunctor    func() (TranslatedWordsPair, error)
	availableAnswersSize uint8
	reversed             bool
	lastCorrectWord      *LanguageWord
}

func NewLanguageGuessingGame(getNewWordFunctorInst func() (TranslatedWordsPair, error), availableAnswers uint8) (game *GuessingGame, err error) {
	if getNewWordFunctorInst == nil {
		err = errors.New("word retrieving functor is invalid")
		return
	}

	game = &GuessingGame{
		getNewWordFunctor:    getNewWordFunctorInst,
		availableAnswersSize: availableAnswers,
		reversed:             false,
		lastCorrectWord:      nil,
	}
	return
}

func (game *GuessingGame) GetNewGuessingWordPack() (wordPack GuessingWordPack, err error) {
	if game == nil || game.getNewWordFunctor == nil {
		err = errors.New("game object or word retrieval functor is invalid")
	} else {
		randomAnswerIdx := rand.Intn(int(game.availableAnswersSize))
		answers := make([]LanguageWord, game.availableAnswersSize)
		for i := 0; i < int(game.availableAnswersSize); i++ {
			wordPairTmp, errTmp := game.getNewWordFunctor()
			if errTmp != nil {
				err = errTmp
				return
			}

			if !game.reversed {
				answers[i] = wordPairTmp.Word
				if randomAnswerIdx == i {
					wordPack.AskedForWord = wordPairTmp.Word
					wordPack.TranslatedWord = wordPairTmp.Translated
					game.lastCorrectWord = &wordPairTmp.Translated
				}

			} else {
				answers[i] = wordPairTmp.Translated
				if randomAnswerIdx == i {
					wordPack.AskedForWord = wordPairTmp.Translated
					wordPack.TranslatedWord = wordPairTmp.Word
					game.lastCorrectWord = &wordPairTmp.Word
				}
			}
		}

		wordPack.IncorrectlyTranslatedWords = answers
	}
	return
}

func (game *GuessingGame) Answer(hash WordHash) (correct bool, correctWord *LanguageWord, err error) {
	if game == nil {
		err = errors.New("game object is invalid")
	} else if game.lastCorrectWord == nil {
		err = errors.New("Game did not start yet")
	} else {
		correct = hash == game.lastCorrectWord.Hash
		correctWord = game.lastCorrectWord
	}
	return
}

func (game *GuessingGame) StopGame() {
	if game != nil {
		game.lastCorrectWord = nil
	}
}
