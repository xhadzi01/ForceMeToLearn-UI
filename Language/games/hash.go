package game

import (
	"hash/fnv"
)

type WordHash uint32

func GetLanguageWordHash(languageWord string) WordHash {
	h := fnv.New32a()
	h.Write([]byte(languageWord))
	return WordHash(h.Sum32())
}
