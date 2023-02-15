package language

type ScoreWordPair struct {
	correct LanguageWordScore
	total   LanguageWordScore
	ratio   LanguageWordsRatio
}

func NewScoreWordPair() *ScoreWordPair {
	wordPair := &ScoreWordPair{
		correct: 0,
		total:   0,
		ratio:   0.0,
	}

	return wordPair
}

type LanguageScore struct {
	current *ScoreWordPair
	total   *ScoreWordPair
}
