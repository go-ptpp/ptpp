package ptpp

import (
	"archive/zip"
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"sync"
)

// LoadSaver denotes an object that can store and restore its state.
type LoadSaver interface {

	// Load restores the state of the object from r.
	Load(r io.Reader) error

	// Save stores the state of the object into w.
	Save(w io.Writer) error
}

// SpellChecker is a word spell-checker.
type SpellChecker interface {

	// Check finds correct spell suggestions for a word. The resulted suggestion
	// list should not be empty.
	Check(word string) []string

	// Train trains the suggestion model with a list of words.
	Train(words []string)
}

// SemanticMatcher provides selection of the best suggestion by its context.
type SemanticMatcher interface {

	// Match finds the best suggestion based on the context. If the best
	// suggestion correlates with the context, it returns true for matched,
	// otherwise it returns false.
	Match(context string, suggestions []string) (best string, matched bool)

	// Train trains the semantic models with a word and its context.
	Train(context, word string)
}

// Processor is the Persian text preprocessor.
type Processor struct {

	// SpellChecker is the word spell-checker. If this field is nil, the
	// preprocessor will use DefaultSpellChecker.
	SpellChecker SpellChecker

	// SemanticMatcher is the semantic matcher to find the best suggestion. If
	// this field is nil, the preprocessor will use DefaultSemanticMatcher.
	SemanticMatcher SemanticMatcher

	mutex sync.Mutex
}

func (p *Processor) ensureFields() {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if p.SpellChecker == nil {
		p.SpellChecker = &DefaultSpellChecker{}
	}
	if p.SemanticMatcher == nil {
		p.SemanticMatcher = &DefaultSemanticMatcher{}
	}
}

// Train trains the preprocessing model with a list of phrases.
func (p *Processor) Train(phrases []string) {
	p.ensureFields()

	for _, phrase := range phrases {
		words, _ := readWords(strings.NewReader(phrase))
		p.SpellChecker.Train(words)
		for i := 0; i < len(words)-1; i++ {
			p.SemanticMatcher.Train(words[i], words[i+1])
		}
	}
}

// Process does the preprocessing on an input and extracts phrases.
func (p *Processor) Process(r io.Reader) ([]string, error) {
	p.ensureFields()

	words, err := readWords(r)
	if err != nil {
		return nil, err
	}

	phrases := []string{}
	currentPhrase := []string{}

	for _, word := range words {
		suggestions := p.SpellChecker.Check(word)
		if len(currentPhrase) == 0 {
			currentPhrase = append(currentPhrase, suggestions[0])
			continue
		}

		context := currentPhrase[len(currentPhrase)-1]
		best, matched := p.SemanticMatcher.Match(context, suggestions)
		if !matched {
			phrases = append(phrases, strings.Join(currentPhrase, " "))
			currentPhrase = []string{best}
		} else {
			currentPhrase = append(currentPhrase, best)
		}
	}

	if len(currentPhrase) != 0 {
		phrases = append(phrases, strings.Join(currentPhrase, " "))
	}

	return phrases, nil
}

func readWords(r io.Reader) ([]string, error) {
	words := []string{}
	br := bufio.NewReader(r)

	const (
		Start   = 0
		English = 1
		Farsi   = 2
		Number  = 3
	)
	state := Start
	sb := strings.Builder{}

	for {
		ch, _, err := br.ReadRune()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		switch state {
		case Start:
			switch {
			case isEnglishLetter(ch):
				sb.WriteRune(normalize(ch))
				state = English
			case isArabicOrFarsiLetter(ch):
				sb.WriteRune(normalize(ch))
				state = Farsi
			case isDigit(ch):
				sb.WriteRune(normalize(ch))
				state = Number
			default:
				// Drop unknown runes.
			}
		case English:
			if isEnglishLetter(ch) {
				sb.WriteRune(normalize(ch))
			} else {
				words = append(words, sb.String())
				sb.Reset()
				br.UnreadRune()
				state = Start
			}
		case Farsi:
			if isArabicOrFarsiLetter(ch) {
				sb.WriteRune(normalize(ch))
			} else if isTashkil(ch) {
				// Drop tashkils from words.
			} else {
				words = append(words, sb.String())
				sb.Reset()
				br.UnreadRune()
				state = Start
			}
		case Number:
			if isDigit(ch) {
				sb.WriteRune(normalize(ch))
			} else {
				words = append(words, sb.String())
				sb.Reset()
				br.UnreadRune()
				state = Start
			}
		}
	}

	if sb.Len() > 0 {
		words = append(words, sb.String())
	}

	return words, nil
}

const (
	scFileName = "sc.gob"
	smFileName = "sm.gob"
)

// Load restores the state of the processor from filePath.
func (p *Processor) Load(filePath string) error {
	p.ensureFields()

	zr, err := zip.OpenReader(filePath)
	if err != nil {
		return err
	}
	defer zr.Close()

	for _, f := range zr.File {
		switch f.Name {
		case scFileName:
			if err := loadFromZipFile(p.SpellChecker, f); err != nil {
				return err
			}
		case smFileName:
			if err := loadFromZipFile(p.SemanticMatcher, f); err != nil {
				return err
			}
		}
	}

	return nil
}

func loadFromZipFile(v interface{}, f *zip.File) error {
	loader, ok := v.(LoadSaver)
	if !ok {
		return nil
	}

	r, err := f.Open()
	if err != nil {
		return err
	}
	defer r.Close()

	return loader.Load(r)
}

// Save stores the state of the processor into w.
func (p *Processor) Save(filePath string) (err error) {
	p.ensureFields()

	tempFile, err := ioutil.TempFile(path.Dir(filePath), "temp-*")
	if err != nil {
		return err
	}

	zw := zip.NewWriter(tempFile)

	if err = saveToZipFile(p.SpellChecker, zw, scFileName); err != nil {
		goto fail
	}

	if err = saveToZipFile(p.SemanticMatcher, zw, smFileName); err != nil {
		goto fail
	}

	if err = zw.Close(); err != nil {
		goto fail
	}

	tempFile.Close()

	return os.Rename(tempFile.Name(), filePath)

fail:
	tempFile.Close()
	os.Remove(tempFile.Name())

	return err
}

func saveToZipFile(v interface{}, zw *zip.Writer, name string) error {
	saver, ok := v.(LoadSaver)
	if !ok {
		return nil
	}

	w, err := zw.Create(name)
	if err != nil {
		return err
	}

	if err := saver.Save(w); err != nil {
		return err
	}

	return nil
}
