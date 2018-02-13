// ===========================================================================
//
// File Name:  seqconv.go
//
// Author:  Jonathan Kans
//
// ==========================================================================

package seqconv

import (
	"bytes"
	"github.com/cwvhogue/ncbigo/pkg/stack"
)

/*
EXTENDED NUCLEOTIDE 4-BIT AMBIGUITY ALPHABET

BASE   NUM   BITS    SET   COMP   MNEMONIC

  X     0    0000            X    unknown

  A     1    0001    A       T    Adenine
  C     2    0010    C       G    Cytosine
  G     4    0100    G       C    Guanine
  T     8    1000    T       A    Thymine

  M     3    0011    AC      K    aMino
  R     5    0101    AG      Y    puRine
  S     6    0110    CG      S    Strong
  W     9    1001    AT      W    Weak
  Y    10    1010    CT      R    pYrimidine
  K    12    1100    GT      M    Keto

  V     7    0111    ACG     B    not T
  H    11    1011    ACT     D    not G
  D    13    1101    AGT     H    not C
  B    14    1110    CGT     V    not A

  N    15    1111    ACGT    N    unkNown
*/

// byte-to-string lookup tables
var ncbi2naToIupacna [256]string
var ncbi4naToIupacna [256]string

// reciprocal string-to-byte maps
var iupacnaToNcbi2na = make(map[string]byte)
var iupacnaToNcbi4na = make(map[string]byte)

// init2na - initialize the 2NA converter.
func init2na() {
	letters := "ACGT"
	base := []byte{0, 0, 0, 0}
	i := byte(0)
	// iterate over all 4-letter unambiguous nucleotide combinations
	for _, a := range letters {
		base[0] = byte(a)
		for _, b := range letters {
			base[1] = byte(b)
			for _, c := range letters {
				base[2] = byte(c)
				for _, d := range letters {
					base[3] = byte(d)
					str := string(base[:4])
					iupacnaToNcbi2na[str] = i
					ncbi2naToIupacna[i] = str
					i++
				}
			}
		}
	}
}

// init4na - initialize the 4NA converter.
func init4na() {
	letters := "XACMGRSVTWYHKDBN"
	base := []byte{0, 0}
	i := byte(0)
	// iterate over all 2-letter nucleotide ambiguity combinations
	for _, a := range letters {
		base[0] = byte(a)
		for _, b := range letters {
			base[1] = byte(b)
			str := string(base[:2])
			iupacnaToNcbi4na[str] = i
			ncbi4naToIupacna[i] = str
			i++
		}
	}
}

// TODO - cleanup and revcomp maps include U, and accept lower-case input data
var iupac2Map = map[byte]byte{
	'A': 'A',
	'C': 'C',
	'G': 'G',
	'T': 'T',
	'U': 'T',
	'a': 'A',
	'c': 'C',
	'g': 'G',
	't': 'T',
	'u': 'T',
}

var iupac4Map = map[byte]byte{
	'A': 'A',
	'B': 'B',
	'C': 'C',
	'D': 'D',
	'G': 'G',
	'H': 'H',
	'K': 'K',
	'M': 'M',
	'N': 'N',
	'R': 'R',
	'S': 'S',
	'T': 'T',
	'U': 'T',
	'V': 'V',
	'W': 'W',
	'X': 'X',
	'Y': 'Y',
	'a': 'A',
	'b': 'B',
	'c': 'C',
	'd': 'D',
	'g': 'G',
	'h': 'H',
	'k': 'K',
	'm': 'M',
	'n': 'N',
	'r': 'R',
	's': 'S',
	't': 'T',
	'u': 'T',
	'v': 'V',
	'w': 'W',
	'x': 'X',
	'y': 'Y',
}

var revComp = map[rune]rune{
	'A': 'T',
	'B': 'V',
	'C': 'G',
	'D': 'H',
	'G': 'C',
	'H': 'D',
	'K': 'M',
	'M': 'K',
	'N': 'N',
	'R': 'Y',
	'S': 'S',
	'T': 'A',
	'U': 'A',
	'V': 'B',
	'W': 'W',
	'X': 'X',
	'Y': 'R',
	'a': 'T',
	'b': 'V',
	'c': 'G',
	'd': 'H',
	'g': 'C',
	'h': 'D',
	'k': 'M',
	'm': 'K',
	'n': 'N',
	'r': 'Y',
	's': 'S',
	't': 'A',
	'u': 'A',
	'v': 'B',
	'w': 'W',
	'x': 'X',
	'y': 'R',
}

// Unpack2NA - given a buffer of 2-bit encoded nucleic acid sequence, return a string of IUPAC sequence.
func Unpack2NA(data []byte, length int) string {
	var buffer bytes.Buffer
	for _, d := range data {
		buffer.WriteString(ncbi2naToIupacna[d])
	}
	str := buffer.String()
	// truncate if sequence ends inside the last byte
	if length > 0 && length < len(str) {
		str = str[:length]
	}
	return str
}

// Unpack4NA - given a buffer of 4-bit encoded nucleic acid sequence, return a string of IUPAC sequence.
func Unpack4NA(data []byte, length int) string {
	var buffer bytes.Buffer
	for _, d := range data {
		buffer.WriteString(ncbi4naToIupacna[d])
	}
	str := buffer.String()
	// truncate if sequence ends in middle of last byte
	if length > 0 && length < len(str) {
		str = str[:length]
	}
	return str
}

// Compress2NA - given a string of IUPAC nucleic acid sequence, return a byte array of 2-bit encoded sequence.
func Compress2NA(data string) ([]byte, int, *stack.ErrorT) {
	base := []byte{0, 0, 0, 0}

	max := len(data)

	var arry []byte
	var ok bool

	for i, j := 0, 0; i < max; i = j {
		j += 4
		if j > max {
			j = max
		}
		k := 0
		for p := i; p < j; p++ {
			ch := data[p]
			base[k], ok = iupac2Map[ch]
			if !ok {
				return nil, 0, stack.ErrorF("failed to find base in iupac2map with `%v`", ch)
			}
			k++
		}
		for k < 4 {
			base[k] = 'A'
			k++
		}

		b, ok := iupacnaToNcbi2na[string(base)]
		if !ok {
			return nil, 0, stack.ErrorF("failed to find %s in iupacnaToNcbi2na map", string(base))
		}
		arry = append(arry, b)
	}

	return arry, max, nil
}

// Compress4NA - given a string of IUPAC nucleic acid sequence, return a byte array of 4-bit encoded sequence.
func Compress4NA(data string) ([]byte, int, *stack.ErrorT) {

	base := []byte{0, 0}

	max := len(data)

	var arry []byte
	var ok bool

	for i, j := 0, 0; i < max; i = j {
		j += 2
		if j > max {
			j = max
		}
		k := 0
		for p := i; p < j; p++ {
			ch := data[p]
			base[k], ok = iupac4Map[ch]
			if !ok {
				return nil, 0, stack.ErrorF("failed to find base in iupac4map with `%v`", ch)
			}
			k++
		}
		for k < 2 {
			base[k] = 'X'
			k++
		}

		b, ok := iupacnaToNcbi4na[string(base)]
		if !ok {
			return nil, 0, stack.ErrorF("failed to find %s in iupacnaToNcbi4na map", string(base))
		}
		arry = append(arry, b)
	}

	return arry, max, nil
}


// RevCompIUPAC - given a string of IUPAC nucleic acid sequence, return a string of the reverse complement.
func RevCompIUPAC(seq string) (string, *stack.ErrorT) {

	runes := []rune(seq)

	// reverse sequence letters - middle base in odd-length sequence is not touched, so cannot also complement here
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	var ok bool

	// now complement every base, also upper-casing and handling uracil
	for i, ch := range runes {
		runes[i], ok = revComp[ch]
		if !ok {
			return "", stack.ErrorF("failed to find '%v' in revComp map", ch)
		}
	}

	return string(runes), nil
}
