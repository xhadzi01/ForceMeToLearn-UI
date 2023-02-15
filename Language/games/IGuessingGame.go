package game

// GuessingWordPack contains correct word and incorrect words for a single guessing game
type GuessingWordPack struct {
	AskedForWord               LanguageWord
	TranslatedWord             LanguageWord
	IncorrectlyTranslatedWords []LanguageWord
}

type IGuessingGame interface {
	GetNewGuessingWordPack() (GuessingWordPack, error)
	Answer(WordHash) (bool, *LanguageWord, error)
	StopGame()
}
