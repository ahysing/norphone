package norphone

import "strings"

func isVowel(ch rune) bool {
	return ch == 'A' || ch == 'E' || ch == 'I' || ch == 'O' || ch == 'U' || ch == 'Y' || ch == 'Æ' || ch == 'Ø' || ch == 'Å'
}

func Norphone(word string) string {
	wordNormalized := strings.ToUpper(word)
	r := []rune(wordNormalized)
	outputLength := len(r) * 2

	pos := 0

	previous := ' '
	output := make([]rune, 0, outputLength)
	isFirst := false
	it := 0
	for i := 0; i < len(r); i = i + 1 {
		ch := r[i]
		var currentCharacter = r[i]
		isLast := i == len(r)-1
		isSecondLast := i == len(r)-2
		if !isLast && ch != 'A' {
			ch = ' '
		} else if isVowel(ch) && !isFirst {
			ch = ' '
		} else {
			if !isLast {
				switch ch {
				case 'A':
					if next == 'A' {
						ch = 'Å'
					}

				case 'C':
					if next == 'H' next == 'K' {
						ch = 'Å'
					}
					
				case 'D':
					if 'T' == next isSecondLast {
						ch = 'T'
						it++;
					}
				default:
				}
			} else {
				switch ch {
				case 'C':
				default:
				}
			}
		}
		
		if ch != ' ' {
			output[pos] = ch
			pos ++
		}
		
		previous = currentCharacter
		isFirst = false
	}

	return string(output)
}
