// ===========================================================================
//
// File Name:  seqconv_test.go
//
// Author:  Jonathan Kans
//
// ==========================================================================

package seqconv

import (
	"fmt"
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type SeqConvTesterT struct {
	// dir string   a transient /tmp directory for testing - if needed for bigger tests
}

var _ = Suite(&SeqConvTesterT{})

func (s *SeqConvTesterT) SetUpTest(c *C) {
	// s.dir = c.MkDir()   - if you need a /tmp dir it goes here
	// fmt.Printf("Working Directory: %s", s.dir)
}

func (s *SeqConvTesterT) TearDownSuite(c *C) {
}

func (s *SeqConvTesterT) TestSeqConv(c *C) {

	init2na()
	init4na()

	fmt.Printf("Table: ncbi2naToIupacna\n")
	for j := 0; j < 256; j++ {
		fmt.Printf("%3d %s\n", j, ncbi2naToIupacna[j])
	}

	fmt.Printf("ACGT %d\n", iupacnaToNcbi2na["ACGT"])
	fmt.Printf("AC %d\n", iupacnaToNcbi2na["AC"])
	for j := 0; j < 256; j++ {
		fmt.Printf("%3d %s\n", j, ncbi4naToIupacna[j])
	}
	fmt.Printf("AC %d\n", iupacnaToNcbi4na["AC"])
	fmt.Printf("AN %d\n", iupacnaToNcbi4na["AN"])

	before := "ACGTTGCAAC"
	data, length, cerr := Compress2NA(before)
	fmt.Printf("Compress2NA errors? [%v]\n", cerr)
	c.Assert(cerr, IsNil)
	after := Unpack2NA(data, length)
	c.Assert(before, Equals, after)
	fmt.Printf("BEFORE: '%s'\nAFTER:  '%s'\nLENGTH: %3d\n", before, after, length)

	before = "ACGTMRWSYKVHDBNX"
	data, length, cerr = Compress4NA(before)
	fmt.Printf("Compress4NA errors? [%v]\n", cerr)
	c.Assert(cerr, IsNil)
	after = Unpack4NA(data, length)
	c.Assert(before, Equals, after)
	fmt.Printf("BEFORE: '%s'\nAFTER:  '%s'\nLENGTH: %3d\n", before, after, length)

	before = "GAUcRYN"
	after, cerr = RevCompIUPAC(before)
	fmt.Printf("RevCompIUPAC errors? [%v]\n", cerr)
	c.Assert(cerr, IsNil)
	expect := "NRYGATC"
	c.Assert(after, Equals, expect)
	fmt.Printf("BEFORE: '%s'\nAFTER:  '%s'\n", before, after)
}
