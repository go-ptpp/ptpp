package ptpp

import (
	"sync"
	"unicode/utf8"
)

type wordList map[string]bool

func (w wordList) Add(word string) {
	w[word] = true
}

func (w wordList) Has(word string) bool {
	_, ok := w[word]
	return ok
}

// DefaultSpellChecker is a SpellChecker which uses a distance model to find
// suggestions for a misspelled word.
type DefaultSpellChecker struct {
	lexicon map[int]wordList
	mutex   sync.RWMutex
}

// Check finds correct spell suggestions for a word.
func (sc *DefaultSpellChecker) Check(word string) []string {
	sc.mutex.RLock()
	defer sc.mutex.RUnlock()

	length := utf8.RuneCountInString(word)
	if sc.lexicon == nil || length < 2 {
		return []string{word}
	}

	suggestions := []string{}

	if list, ok := sc.lexicon[length]; ok {
		if list.Has(word) {
			suggestions = append(suggestions, word)
		}

		for w := range list {
			if Levenshtein(word, w) == 1 {
				suggestions = append(suggestions, w)
			}
		}
	}

	if length > 2 {
		if list, ok := sc.lexicon[length-1]; ok {
			for w := range list {
				if Levenshtein(word, w) == 1 {
					suggestions = append(suggestions, w)
				}
			}
		}
	}

	if list, ok := sc.lexicon[length+1]; ok {
		for w := range list {
			if Levenshtein(word, w) == 1 {
				suggestions = append(suggestions, w)
			}
		}
	}

	if len(suggestions) == 0 {
		suggestions = append(suggestions, word)
	}

	return suggestions
}

// Train trains the suggestion model with a list of words.
func (sc *DefaultSpellChecker) Train(words []string) {
	sc.mutex.Lock()
	defer sc.mutex.Unlock()

	for _, w := range words {
		sc.trainWord(w)
	}
}

func (sc *DefaultSpellChecker) trainWord(word string) {
	if sc.lexicon == nil {
		sc.lexicon = make(map[int]wordList)
	}

	len := utf8.RuneCountInString(word)
	if len < 2 {
		return
	}

	if _, ok := sc.lexicon[len]; !ok {
		sc.lexicon[len] = make(wordList)
	}

	sc.lexicon[len].Add(word)
}
