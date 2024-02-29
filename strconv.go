package stalw

// CyrToLat returns a string in which Latin letters are substituted for Cyrillic letters.
// The replacement condition is valid only for visually identical letters,
// otherwise the letter in the string will remain unchanged.
//
// For example, 'йцукенг' ('ЙЦУКЕНГ') will turn into 'йцykehг' ('ЙЦYKEHГ'),
// where 'у', 'к', 'е', 'н' ('У', 'К', 'Е', 'Н') have become Latin 'y', 'k', 'e', 'h' ('Y', 'K', 'E', 'H').
func CyrToLat(input string) string {

}
