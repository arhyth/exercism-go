// Package strand exposes ToRNA function.
package strand

type RNA = rune
type DNA = rune

const (
	Cr RNA = 'C'
	Gr RNA = 'G'
	Ar RNA = 'A'
	Ur RNA = 'U'
)

const (
	Gd DNA = 'G'
	Cd DNA = 'C'
	Td DNA = 'T'
	Ad DNA = 'A'
)


// ToRNA takes a DNA string and returns RNA complement string.
func ToRNA(dna string) string {
	if dna == "" {
		return dna
	}

	complement := []RNA{}

	for _, ntide := range dna {
		switch ntide {
		case Gd:
			complement = append(complement, Cr)
		case Cd:
			complement = append(complement, Gr)
		case Td:
			complement = append(complement, Ar)
		case Ad:
			complement = append(complement, Ur)
		}
	}

	return string(complement)
}
