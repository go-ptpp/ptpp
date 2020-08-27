package ptpp

// DefaultSemanticMatcher is a SemanticMatcher which uses a simple lookup table
// to find the best suggestion.
type DefaultSemanticMatcher struct {
	contexts map[string]wordList
}

// Match finds the best suggestion based on the context.
func (sm *DefaultSemanticMatcher) Match(context string, suggestions []string) (string, bool) {
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
	if sm.contexts == nil {
		sm.contexts = make(map[string]wordList)
	}

	if _, ok := sm.contexts[context]; !ok {
		sm.contexts[context] = make(wordList)
	}

	sm.contexts[context].Add(word)
}
