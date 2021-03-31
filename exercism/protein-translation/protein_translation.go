package protein

import (
	"errors"
)

var ErrStop error = errors.New("Stop Error")
var ErrInvalidBase error = errors.New("Invalid Base Error")

var proteins = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP",
	"UAG": "STOP",
	"UGA": "STOP",
}

func FromCodon(codon string) (string, error) {
	protein := proteins[codon]
	var err error = nil
	if protein == "STOP" {
		protein = ""
		err = ErrStop
	} else if protein == "" {
		err = ErrInvalidBase
	}
	return protein, err
}

func FromRNA(rna string) ([]string, error) {
	proteins := []string{}
	err := error(nil)
	i := 0
	for {
		if i == len(rna) {
			break
		} else if i+3 > len(rna) {
			err = ErrInvalidBase
			break
		}

		codon := rna[i : i+3]
		protein := ""
		protein, err = FromCodon(codon)
		if err != nil {
			if err == ErrStop {
				err = nil
			}
			break
		}
		proteins = append(proteins, protein)
		i = i + 3
	}
	return proteins, err
}
