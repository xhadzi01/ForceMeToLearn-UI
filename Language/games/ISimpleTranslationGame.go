package game

// SimpleTranslationWordPack contains correct word and incorrect word for a single translating game
type SimpleTranslationWordPack struct {
	AskedForWord   LanguageWord
	TranslatedWord LanguageWord
}

type ISimpleTranslationGame interface {
	GetNewSimpleTranslationWordPack() (SimpleTranslationWordPack, error)
	Answer(WordHash) (bool, error)
	StopGame()
}
