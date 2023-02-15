package game

import (
	"errors"
)

type SimpleTranslationGame struct {
	getNewWordFunctor func() (TranslatedWordsPair, error)
	reversed          bool
	lastCorrectWord   *LanguageWord
}

func NewSimpleTranslationGame(getNewWordFunctorInst func() (TranslatedWordsPair, error)) (game *SimpleTranslationGame, err error) {
	if getNewWordFunctorInst == nil {
		err = errors.New("word retrieving functor is invalid")
		return
	}

	game = &SimpleTranslationGame{
		getNewWordFunctor: getNewWordFunctorInst,
		reversed:          false,
		lastCorrectWord:   nil,
	}
	return
}

func (game *SimpleTranslationGame) GetNewSimpleTranslationWordPack() (translationWordPack SimpleTranslationWordPack, err error) {
	if game == nil || game.getNewWordFunctor == nil {
		err = errors.New("game object or word retrieval functor is invalid")
	} else {
		wordPairTmp, errTmp := game.getNewWordFunctor()
		if errTmp != nil {
			err = errTmp
			return
		}

		if !game.reversed {
			translationWordPack.AskedForWord = wordPairTmp.Word
			translationWordPack.TranslatedWord = wordPairTmp.Translated
			game.lastCorrectWord = &wordPairTmp.Translated
		} else {
			translationWordPack.AskedForWord = wordPairTmp.Translated
			translationWordPack.TranslatedWord = wordPairTmp.Word
			game.lastCorrectWord = &wordPairTmp.Word
		}
	}
	return
}

func (game *SimpleTranslationGame) Answer(hash WordHash) (correct bool, err error) {
	if game == nil {
		err = errors.New("game object is invalid")
	} else if game.lastCorrectWord == nil {
		err = errors.New("Game did not start yet")
	} else {
		correct = hash == game.lastCorrectWord.Hash
	}
	return
}

func (game *SimpleTranslationGame) StopGame() {
	if game != nil {
		game.lastCorrectWord = nil
	}
}
