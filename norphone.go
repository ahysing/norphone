// Norphone
// Norphone is a nice sound hasher which hashes words. It is designed to produce the same hash for words with similar sound in norwegian.
//
// Test suite
//		Aarrestad & Årrestad => ÅRST
//		Andreasen & Andreassen => ANDRSN
//		Arntsen & Arntzen => ARNTSN
//		Bache & Bakke => BK
//		Frank & Franck => FRNK
//		Christian & Kristian => KRSTN
//		Kielland & Kjelland => XLND
//		Krogh & Krog => KRG
//		Krog & Krohg => KRG
//		Jendal & Jendahl => JNDL
//		Jendal & Hjendal => JNDL
//		Jendal & Gjendal => JNDL
//		Vold & Wold => VL
//		Thomas & Tomas => TMS
//		Aamodt & Aamot => ÅMT
//		Aksel & Axel => AKSL
//		Kristoffersen & Christophersen => KRSTFRSN
//		Voll & Vold => VL
//		Granli & Granlid => GRNL
//		Gjever & Giever => JVR
package main

import "strings"

func isVowel(ch rune) bool {
	return ch == 'A' || ch == 'E' || ch == 'I' || ch == 'O' || ch == 'U' || ch == 'Y' || ch == 'Æ' || ch == 'Ø' || ch == 'Å'
}

func Norphone(word string) string {
	if len(word) <= 1 {
		return ""
	}

	wordNormalized := strings.ToUpper(word)
	r := []rune(wordNormalized)
	outputLength := len(r) * 2

	pos := 0

	previous := ' '
	output := make([]rune, 0, outputLength)
	isFirst := true

	for i := 0; i < len(r); i++ {
		isLast := i == len(r)-1
		isSecondLast := i == len(r)-2
		ch := r[i]
		currentCharacter := r[i]
		var oneAhead rune
		if !isLast {
			oneAhead = r[i+1]
		}

		if !isLast && ch == oneAhead && ch != 'A' {
			ch = ' '
		} else if isVowel(ch) && !isFirst {
			ch = ' '
		} else if !isLast {
			switch ch {
			case 'A':
				if oneAhead == 'A' {
					ch = 'Å'
				}

			case 'C':
				if oneAhead == 'H' || oneAhead == 'K' {
					ch = 'K'
					i++
				} else {
					ch = 'K'
				}

			case 'D':
				if 'T' == oneAhead || isSecondLast {
					ch = 'T'
					i++
				} else if isVowel(previous) && isLast {
					ch = ' '
				}

			case 'G':
				if 'H' == oneAhead {
					i++ // 'H' is silent
				} else if ('J' == oneAhead || 'I' == oneAhead) && isFirst {
					ch = 'J'
					i++
				}

			case 'H':
				if 'J' == oneAhead {
					ch = 'J'
					i++
				} else if 'L' == oneAhead {
					ch = 'L'
					i++
				} else if 'G' == oneAhead {
					ch = 'G'
					i++
				} else if 'R' == oneAhead {
					ch = 'R'
					i++
				}

			case 'K':
				if 'I' == oneAhead || 'J' == oneAhead {
					ch = 'X'
					i++
				}

			case 'L':
				if 'D' == oneAhead && isSecondLast {
					ch = 'L'
					i++
				}

			case 'N':
				if 'D' == oneAhead {
					i++
				}

			case 'P':
				if 'H' == oneAhead {
					ch = 'F'
					i++ // eat the 'H'
				}

			case 'T':
				if 'H' == oneAhead {
					i++ // 'H' is silent
				}

			case 'W':
				ch = 'V'

			case 'X':
				output = append(output, 'K')
				pos++
				ch = 'S'

			case 'Z':
				ch = 'S'
			}
		} else {
			switch ch {
			case 'D':
				if isVowel(previous) && isLast {
					ch = ' '
				}

			case 'W':
				ch = 'V'

			case 'X':
				output = append(output, 'K')
				pos++
				ch = 'S'

			case 'Z':
				ch = 'S'
			}
		}

		if ch != ' ' {
			output = append(output, ch)
			pos++
		}

		previous = currentCharacter
		isFirst = false
	}

	var outputFinal = output[:pos]
	return string(outputFinal)
}
