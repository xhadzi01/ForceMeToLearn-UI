package language

import (
	"math/rand"
	"time"

	"errors"
)

type LanguageDatabase struct {
	data []WordWithResult
}

var DatabaseInst *LanguageDatabase

func init() {
	rand.Seed(time.Now().Unix())

	DatabaseInst.data = make([]WordWithResult, 0)

	DatabaseInst.data = append(DatabaseInst.data, WordWithResult{word: "car", translatedWord: "auto"})
	DatabaseInst.data = append(DatabaseInst.data, WordWithResult{word: "house", translatedWord: "dom"})
	DatabaseInst.data = append(DatabaseInst.data, WordWithResult{word: "tree", translatedWord: "strom"})
	DatabaseInst.data = append(DatabaseInst.data, WordWithResult{word: "book", translatedWord: "kniha"})
	DatabaseInst.data = append(DatabaseInst.data, WordWithResult{word: "pen", translatedWord: "pero"})
	DatabaseInst.data = append(DatabaseInst.data, WordWithResult{word: "pencil", translatedWord: "ceruzka"})
	DatabaseInst.data = append(DatabaseInst.data, WordWithResult{word: "bottom", translatedWord: "spodok"})
	DatabaseInst.data = append(DatabaseInst.data, WordWithResult{word: "socket", translatedWord: "zástrčka"})
	DatabaseInst.data = append(DatabaseInst.data, WordWithResult{word: "foil", translatedWord: "fólia"})
	DatabaseInst.data = append(DatabaseInst.data, WordWithResult{word: "pig", translatedWord: "prasa"})
	DatabaseInst.data = append(DatabaseInst.data, WordWithResult{word: "cat", translatedWord: "mačka"})
	DatabaseInst.data = append(DatabaseInst.data, WordWithResult{word: "dog", translatedWord: "pes"})
	DatabaseInst.data = append(DatabaseInst.data, WordWithResult{word: "television", translatedWord: "televízia"})
	DatabaseInst.data = append(DatabaseInst.data, WordWithResult{word: "lamp", translatedWord: "lampa"})
	DatabaseInst.data = append(DatabaseInst.data, WordWithResult{word: "chocolate", translatedWord: "čokoláda"})
}

func (db *LanguageDatabase) GetRandomWordWithResult() (wr WordWithResult, err error) {
	if db == nil {
		err = errors.New("Language database object is invalid")
	} else if len(db.data) == 0 {
		err = errors.New("There are no items in database")
	} else {
		randomIdx := rand.Intn(len(db.data))
		wr = db.data[randomIdx]
	}

	return
}
