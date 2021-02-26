package strand

import "strings"

func ToRNA(dna string) string {
	conv := map[rune]string{'G': "C", 'C': "G", 'T': "A", 'A': "U"}
	rna := strings.Builder{}
	dnanucs := []rune(dna)
	for _, dnanuc := range dnanucs {
		rnanuc, ok := conv[dnanuc]
		if ok {
			rna.WriteString(rnanuc)
		}
	}
	return rna.String()
}
