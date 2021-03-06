package nucleic

import (
	"github.com/hivdb/nucamino/utils"
	"strings"
)

type NucleicAcid int

const (
	A NucleicAcid = iota
	C
	G
	T
	W
	S
	M
	K
	R
	Y
	B
	D
	H
	V
	N
)

const NumNucleicAcids = 15

var NucleicAcids = [NumNucleicAcids]NucleicAcid{
	A, C, G, T, W, S, M, K, R, Y, B, D, H, V, N,
}

var nucleicAcidLookup = [NumNucleicAcids]string{
	"A",
	"C",
	"G",
	"T",
	"W",
	"S",
	"M",
	"K",
	"R",
	"Y",
	"B",
	"D",
	"H",
	"V",
	"N",
}

var nucleicAcidLookupR = map[rune]NucleicAcid{
	'A': A,
	'C': C,
	'G': G,
	'T': T,
	'W': W,
	'S': S,
	'M': M,
	'K': K,
	'R': R,
	'Y': Y,
	'B': B,
	'D': D,
	'H': H,
	'V': V,
	'N': N,
}

var ambiguousNucleicAcids = [NumNucleicAcids][]NucleicAcid{
	{A},
	{C},
	{G},
	{T},
	{A, T},
	{C, G},
	{A, C},
	{G, T},
	{A, G},
	{C, T},
	{C, G, T},
	{A, G, T},
	{A, C, T},
	{A, C, G},
	{A, C, G, T},
}

func (na NucleicAcid) ToString() string {
	return nucleicAcidLookup[na]
}

func (self NucleicAcid) IsAmbiguous() bool {
	return self > T
}

func ReadString(nucleicAcidSequence string) []NucleicAcid {
	nucleicAcidSequence = strings.ToUpper(
		utils.StripWhiteSpace(nucleicAcidSequence))
	result := make([]NucleicAcid, len(nucleicAcidSequence))
	for idx, runeVal := range nucleicAcidSequence {
		na, present := nucleicAcidLookupR[runeVal]
		if !present {
			na = N
		}
		result[idx] = na
	}
	return result
}

func WriteString(nas []NucleicAcid) string {
	var result string
	for _, na := range nas {
		result += nucleicAcidLookup[na]
	}
	return result
}

func GetUnambiguousNucleicAcids(na NucleicAcid) []NucleicAcid {
	return ambiguousNucleicAcids[na]
}
