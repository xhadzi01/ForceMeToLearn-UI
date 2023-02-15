package language

import game "github.com/xhadzi01/ForceMeToLearn/Language/games"

type LanguageWordScore uint32
type LanguageWordsRatio float64

// ILanguageScore is interface representing score for different languages
type ILanguageScore interface {
	// GetTotalScore returns correctly answered words, total number of asked words and ratio of correct words
	GetTotalScore() (LanguageWordScore, LanguageWordScore, LanguageWordsRatio, error)

	// GetCurrentScore returns correctly answered words, total number of asked words and ratio of correct words for current game
	GetCurrentScore() (LanguageWordScore, LanguageWordScore, LanguageWordsRatio, error)

	// ResetScore resets counters for current game
	ResetScore() error
}

// ILanguage is interface representing different languages
type ILanguage interface {
	// GetName return name of the language
	GetName() (string, error)

	// ILanguageScore contains score for current game and also for total
	*ILanguageScore

	// GetNewGuessingGame returns guessing game
	GetNewGuessingGame() (game.IGuessingGame, error)

	// GetNewTranslationGame returns simple translation game
	GetNewTranslationGame() (game.ISimpleTranslationGame, error)
}
