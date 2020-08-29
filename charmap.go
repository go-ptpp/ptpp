package ptpp

const (
	englishLetter = iota
	englishDigit
	arabicOrFarsiLetter
	arabicOrFarsiDigit
	tashkil
)

var charmap = map[rune]struct {
	Class      int
	Normalized rune
}{
	'A':      {Class: englishLetter, Normalized: 'a'},
	'B':      {Class: englishLetter, Normalized: 'b'},
	'C':      {Class: englishLetter, Normalized: 'c'},
	'D':      {Class: englishLetter, Normalized: 'd'},
	'E':      {Class: englishLetter, Normalized: 'e'},
	'F':      {Class: englishLetter, Normalized: 'f'},
	'G':      {Class: englishLetter, Normalized: 'g'},
	'H':      {Class: englishLetter, Normalized: 'h'},
	'I':      {Class: englishLetter, Normalized: 'i'},
	'J':      {Class: englishLetter, Normalized: 'j'},
	'K':      {Class: englishLetter, Normalized: 'k'},
	'L':      {Class: englishLetter, Normalized: 'l'},
	'M':      {Class: englishLetter, Normalized: 'm'},
	'N':      {Class: englishLetter, Normalized: 'n'},
	'O':      {Class: englishLetter, Normalized: 'o'},
	'P':      {Class: englishLetter, Normalized: 'p'},
	'Q':      {Class: englishLetter, Normalized: 'q'},
	'R':      {Class: englishLetter, Normalized: 'r'},
	'S':      {Class: englishLetter, Normalized: 's'},
	'T':      {Class: englishLetter, Normalized: 't'},
	'U':      {Class: englishLetter, Normalized: 'u'},
	'V':      {Class: englishLetter, Normalized: 'v'},
	'W':      {Class: englishLetter, Normalized: 'w'},
	'X':      {Class: englishLetter, Normalized: 'x'},
	'Y':      {Class: englishLetter, Normalized: 'y'},
	'Z':      {Class: englishLetter, Normalized: 'z'},
	'a':      {Class: englishLetter, Normalized: 'a'},
	'b':      {Class: englishLetter, Normalized: 'b'},
	'c':      {Class: englishLetter, Normalized: 'c'},
	'd':      {Class: englishLetter, Normalized: 'd'},
	'e':      {Class: englishLetter, Normalized: 'e'},
	'f':      {Class: englishLetter, Normalized: 'f'},
	'g':      {Class: englishLetter, Normalized: 'g'},
	'h':      {Class: englishLetter, Normalized: 'h'},
	'i':      {Class: englishLetter, Normalized: 'i'},
	'j':      {Class: englishLetter, Normalized: 'j'},
	'k':      {Class: englishLetter, Normalized: 'k'},
	'l':      {Class: englishLetter, Normalized: 'l'},
	'm':      {Class: englishLetter, Normalized: 'm'},
	'n':      {Class: englishLetter, Normalized: 'n'},
	'o':      {Class: englishLetter, Normalized: 'o'},
	'p':      {Class: englishLetter, Normalized: 'p'},
	'q':      {Class: englishLetter, Normalized: 'q'},
	'r':      {Class: englishLetter, Normalized: 'r'},
	's':      {Class: englishLetter, Normalized: 's'},
	't':      {Class: englishLetter, Normalized: 't'},
	'u':      {Class: englishLetter, Normalized: 'u'},
	'v':      {Class: englishLetter, Normalized: 'v'},
	'w':      {Class: englishLetter, Normalized: 'w'},
	'x':      {Class: englishLetter, Normalized: 'x'},
	'y':      {Class: englishLetter, Normalized: 'y'},
	'z':      {Class: englishLetter, Normalized: 'z'},
	'0':      {Class: englishDigit, Normalized: '0'},
	'1':      {Class: englishDigit, Normalized: '1'},
	'2':      {Class: englishDigit, Normalized: '2'},
	'3':      {Class: englishDigit, Normalized: '3'},
	'4':      {Class: englishDigit, Normalized: '4'},
	'5':      {Class: englishDigit, Normalized: '5'},
	'6':      {Class: englishDigit, Normalized: '6'},
	'7':      {Class: englishDigit, Normalized: '7'},
	'8':      {Class: englishDigit, Normalized: '8'},
	'9':      {Class: englishDigit, Normalized: '9'},
	'\u0622': {Class: arabicOrFarsiLetter, Normalized: '\u0627'},
	'\u0623': {Class: arabicOrFarsiLetter, Normalized: '\u0627'},
	'\u0624': {Class: arabicOrFarsiLetter, Normalized: '\u0648'},
	'\u0625': {Class: arabicOrFarsiLetter, Normalized: '\u0627'},
	'\u0626': {Class: arabicOrFarsiLetter, Normalized: '\u06CC'},
	'\u0627': {Class: arabicOrFarsiLetter, Normalized: '\u0627'},
	'\u0628': {Class: arabicOrFarsiLetter, Normalized: '\u0628'},
	'\u0629': {Class: arabicOrFarsiLetter, Normalized: '\u062A'},
	'\u062A': {Class: arabicOrFarsiLetter, Normalized: '\u062A'},
	'\u062B': {Class: arabicOrFarsiLetter, Normalized: '\u062B'},
	'\u062C': {Class: arabicOrFarsiLetter, Normalized: '\u062C'},
	'\u062D': {Class: arabicOrFarsiLetter, Normalized: '\u062D'},
	'\u062E': {Class: arabicOrFarsiLetter, Normalized: '\u062E'},
	'\u062F': {Class: arabicOrFarsiLetter, Normalized: '\u062F'},
	'\u0630': {Class: arabicOrFarsiLetter, Normalized: '\u0630'},
	'\u0631': {Class: arabicOrFarsiLetter, Normalized: '\u0631'},
	'\u0632': {Class: arabicOrFarsiLetter, Normalized: '\u0632'},
	'\u0633': {Class: arabicOrFarsiLetter, Normalized: '\u0633'},
	'\u0634': {Class: arabicOrFarsiLetter, Normalized: '\u0634'},
	'\u0635': {Class: arabicOrFarsiLetter, Normalized: '\u0635'},
	'\u0636': {Class: arabicOrFarsiLetter, Normalized: '\u0636'},
	'\u0637': {Class: arabicOrFarsiLetter, Normalized: '\u0637'},
	'\u0638': {Class: arabicOrFarsiLetter, Normalized: '\u0638'},
	'\u0639': {Class: arabicOrFarsiLetter, Normalized: '\u0639'},
	'\u063A': {Class: arabicOrFarsiLetter, Normalized: '\u063A'},
	'\u0641': {Class: arabicOrFarsiLetter, Normalized: '\u0641'},
	'\u0642': {Class: arabicOrFarsiLetter, Normalized: '\u0642'},
	'\u0643': {Class: arabicOrFarsiLetter, Normalized: '\u06A9'},
	'\u0644': {Class: arabicOrFarsiLetter, Normalized: '\u0644'},
	'\u0645': {Class: arabicOrFarsiLetter, Normalized: '\u0645'},
	'\u0646': {Class: arabicOrFarsiLetter, Normalized: '\u0646'},
	'\u0647': {Class: arabicOrFarsiLetter, Normalized: '\u0647'},
	'\u0648': {Class: arabicOrFarsiLetter, Normalized: '\u0648'},
	'\u0649': {Class: arabicOrFarsiLetter, Normalized: '\u06CC'},
	'\u064A': {Class: arabicOrFarsiLetter, Normalized: '\u06CC'},
	'\u067E': {Class: arabicOrFarsiLetter, Normalized: '\u067E'},
	'\u0686': {Class: arabicOrFarsiLetter, Normalized: '\u0686'},
	'\u0698': {Class: arabicOrFarsiLetter, Normalized: '\u0698'},
	'\u06A9': {Class: arabicOrFarsiLetter, Normalized: '\u06A9'},
	'\u06AF': {Class: arabicOrFarsiLetter, Normalized: '\u06AF'},
	'\u06CC': {Class: arabicOrFarsiLetter, Normalized: '\u06CC'},
	'\u0660': {Class: arabicOrFarsiDigit, Normalized: '0'},
	'\u0661': {Class: arabicOrFarsiDigit, Normalized: '1'},
	'\u0662': {Class: arabicOrFarsiDigit, Normalized: '2'},
	'\u0663': {Class: arabicOrFarsiDigit, Normalized: '3'},
	'\u0664': {Class: arabicOrFarsiDigit, Normalized: '4'},
	'\u0665': {Class: arabicOrFarsiDigit, Normalized: '5'},
	'\u0666': {Class: arabicOrFarsiDigit, Normalized: '6'},
	'\u0667': {Class: arabicOrFarsiDigit, Normalized: '7'},
	'\u0668': {Class: arabicOrFarsiDigit, Normalized: '8'},
	'\u0669': {Class: arabicOrFarsiDigit, Normalized: '9'},
	'\u06F0': {Class: arabicOrFarsiDigit, Normalized: '0'},
	'\u06F1': {Class: arabicOrFarsiDigit, Normalized: '1'},
	'\u06F2': {Class: arabicOrFarsiDigit, Normalized: '2'},
	'\u06F3': {Class: arabicOrFarsiDigit, Normalized: '3'},
	'\u06F4': {Class: arabicOrFarsiDigit, Normalized: '4'},
	'\u06F5': {Class: arabicOrFarsiDigit, Normalized: '5'},
	'\u06F6': {Class: arabicOrFarsiDigit, Normalized: '6'},
	'\u06F7': {Class: arabicOrFarsiDigit, Normalized: '7'},
	'\u06F8': {Class: arabicOrFarsiDigit, Normalized: '8'},
	'\u06F9': {Class: arabicOrFarsiDigit, Normalized: '9'},
	'\u0621': {Class: tashkil},
	'\u0640': {Class: tashkil},
	'\u064B': {Class: tashkil},
	'\u064C': {Class: tashkil},
	'\u064D': {Class: tashkil},
	'\u064E': {Class: tashkil},
	'\u064F': {Class: tashkil},
	'\u0650': {Class: tashkil},
	'\u0651': {Class: tashkil},
	'\u0652': {Class: tashkil},
}

func isEnglishLetter(r rune) bool {
	if attrib, ok := charmap[r]; ok {
		return attrib.Class == englishLetter
	}
	return false
}

func isArabicOrFarsiLetter(r rune) bool {
	if attrib, ok := charmap[r]; ok {
		return attrib.Class == arabicOrFarsiLetter
	}
	return false
}

func isTashkil(r rune) bool {
	if attrib, ok := charmap[r]; ok {
		return attrib.Class == tashkil
	}
	return false
}

func isDigit(r rune) bool {
	if attrib, ok := charmap[r]; ok {
		return attrib.Class == englishDigit || attrib.Class == arabicOrFarsiDigit
	}
	return false
}

func normalize(r rune) rune {
	if attrib, ok := charmap[r]; ok {
		return attrib.Normalized
	}
	return r
}
