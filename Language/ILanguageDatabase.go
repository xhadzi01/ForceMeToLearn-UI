package language

type WordWithResult struct {
	word           string
	translatedWord string
}

type ILanguageDatabase interface {
	GetRandomWordWithResult() (WordWithResult, error)
}
