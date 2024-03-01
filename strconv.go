package stalw

// CyrToLat returns a string in which Latin letters are substituted for Cyrillic letters.
// The replacement condition is valid only for visually identical letters,
// otherwise the letter in the string will remain unchanged.
//
// For example, 'йцукенг' ('ЙЦУКЕНГ') will turn into 'йцykehг' ('ЙЦYKEHГ'),
// where 'у', 'к', 'е', 'н' ('У', 'К', 'Е', 'Н') have become Latin 'y', 'k', 'e', 'h' ('Y', 'K', 'E', 'H').
func CyrToLat(input string) string {
	var cyrToLat = map[rune]rune{
		'а': 'a',
		'в': 'b',
		'е': 'e',
		'к': 'k',
		'м': 'm',
		'н': 'h',
		'о': 'o',
		'р': 'p',
		'с': 'c',
		'т': 't',
		'у': 'y',
		'х': 'x',
	}

	runes := []rune(input)

	for i, r := range runes {
		var isCap bool
		symb := r

		if r >= 'А' && r <= 'Я' {
			r += -'А' + 'а'
			isCap = true
		}

		if v, ok := cyrToLat[r]; ok {
			symb = v
			if isCap {
				symb += -'а' + 'А'
			}
		}

		runes[i] = symb
	}

	return string(runes)
}
