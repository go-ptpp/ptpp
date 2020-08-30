package ptpp

import (
	"encoding/gob"
	"io"
	"sync"
)

// DefaultSemanticMatcher is a SemanticMatcher which uses a simple lookup table
// to find the best suggestion.
type DefaultSemanticMatcher struct {
	contexts map[string]wordList
	mutex    sync.RWMutex
}

// Match finds the best suggestion based on the context.
func (sm *DefaultSemanticMatcher) Match(context string, suggestions []string) (string, bool) {
	sm.mutex.RLock()
	defer sm.mutex.RUnlock()

	if sm.contexts == nil {
		return suggestions[0], false
	}

	if list, ok := sm.contexts[context]; ok {
		for _, suggestion := range suggestions {
			if list.Has(suggestion) {
				return suggestion, true
			}
		}
	}

	return suggestions[0], false
}

// Train trains the semantic models with a word and its context.
func (sm *DefaultSemanticMatcher) Train(context, word string) {
	sm.mutex.Lock()
	defer sm.mutex.Unlock()

	if sm.contexts == nil {
		sm.contexts = make(map[string]wordList)
	}

	if _, ok := sm.contexts[context]; !ok {
		sm.contexts[context] = make(wordList)
	}

	sm.contexts[context].Add(word)
}

// Load restores the state of the semantic-matcher from r.
func (sm *DefaultSemanticMatcher) Load(r io.Reader) error {
	sm.mutex.Lock()
	defer sm.mutex.Unlock()

	if sm.contexts == nil {
		sm.contexts = make(map[string]wordList)
	}

	return gob.NewDecoder(r).Decode(&sm.contexts)
}

// Save stores the state of the semantic-matcher into w.
func (sm *DefaultSemanticMatcher) Save(w io.Writer) error {
	sm.mutex.Lock()
	defer sm.mutex.Unlock()

	if sm.contexts == nil {
		sm.contexts = make(map[string]wordList)
	}

	return gob.NewEncoder(w).Encode(sm.contexts)
}
