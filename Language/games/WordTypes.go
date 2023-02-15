package game

// LanguageWord represents one word and also its corresponding hash, that will be used for validity check
type LanguageWord struct {
	Word string
	Hash WordHash
}

func NewLanguageWord(word string) (lang LanguageWord) {
	lang = LanguageWord{
		Hash: GetLanguageWordHash(word),
		Word: word,
	}
	return
}

type TranslatedWordsPair struct {
	Word       LanguageWord
	Translated LanguageWord
}
