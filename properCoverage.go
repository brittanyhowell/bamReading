package main

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"unsafe"

	"github.com/biogo/hts/sam"
)

type reads struct {
	Name       string
	Chrom      string
	ChromStart int
	ChromEnd   int
	L1Start    int
	L1End      int
	ReadStart  int
	ReadEnd    int
	Pos        int
	Cigar      sam.Cigar
}

const (
	NameField = iota
	ChromField
	ChromStartField
	ChromEndField
	L1StartField
	L1EndField
	ReadStartField
	ReadEndField
	PosField
	CigarField
)

var (
	readFile string
	out      string
)

func main() {

	flag.StringVar(&readFile, "readFile", "", "File with read data")
	flag.StringVar(&out, "outFile", "intervals.txt", "cov interval file")
	flag.Parse()

	o := fmt.Sprintf("%v", out)
	oFile, err := os.Create(o)
	if err != nil {
		log.Fatalf("failed to create output file %s: %v", o, err)
	}
	defer oFile.Close()

	f, err := os.Open(readFile)
	if err != nil {
		log.Printf("error: could not open %s to read %v", f, err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	var num int

	for sc.Scan() {

		line := sc.Bytes()

		if err != nil {
			log.Fatalf("failed with %v\n", err)
		}
		num++
		c, err := parseReadFile(line)

		if err != nil {
			if err, ok := err.(*csv.ParseError); ok {
				err.Line = num
				log.Fatalf("failed with %v\n", err)
			}
			fmt.Errorf("%v at line %d", err, num)
		}

		fmt.Println("a line:", c)

		// Get to the consume
		var (
			// startInL1 int
			// endInL1   int
			// startGap  int
			// endGap    int

			intLen  int
			overlap int
			// extra       int
			// readOverlap int
		)
		for _, co := range c.Cigar {
			pos := c.Pos

			t := co.Type()
			con := t.Consumes()
			intLen = co.Len() * con.Reference

			if con.Query == con.Reference {
				o := min(pos+intLen, c.ReadEnd) - max(pos, c.ReadStart)
				if o > 0 {
					overlap += o
				}
				fmt.Println("con stuff:", overlap)
			}
			// 	overlap += extra
			// 	extra = 0
			// fmt.Fprintf(oFile, "%v\t%v", start, end?)
			fmt.Println("Operator: ", co)
			// BRITTANY: Write to file if SAM field says I map.
			// No map? continue. SIMPLE IF STATEMENT. DONESZO.
		}

		// block = append(block, c)

	}

}

// BED format reader type.
type Reader struct {
	r       *bufio.Reader
	BedType int
	line    int
}

func parseReadFile(line []byte) (b reads, err error) {
	const n = 8
	defer handlePanic(b, &err)
	c := bytes.Fields(line)
	if len(c) < n {
		log.Fatalf("wrong number of columns. Have %v, need %v \n", len(c), n)
		return
	}

	b = reads{
		Name:       string(c[NameField]),
		Chrom:      string(c[ChromField]),
		ChromStart: mustAtoi(c[ChromStartField], ChromStartField),
		ChromEnd:   mustAtoi(c[ChromEndField], ChromEndField),
		L1Start:    mustAtoi(c[L1StartField], L1StartField),
		L1End:      mustAtoi(c[L1EndField], L1EndField),
		ReadStart:  mustAtoi(c[ReadStartField], ReadStartField),
		ReadEnd:    mustAtoi(c[ReadEndField], ReadEndField),
		Pos:        mustAtoi(c[PosField], PosField),
		Cigar:      mustCigar(c[CigarField], CigarField),
	}
	return
}

func mustCigar(f []byte, column int) sam.Cigar {
	i, err := sam.ParseCigar(f)
	if err != nil {
		panic(&csv.ParseError{Column: column, Err: err})
	}
	return i
}

func mustAtoi(f []byte, column int) int {
	i, err := strconv.ParseInt(unsafeString(f), 0, 0)
	if err != nil {
		panic(&csv.ParseError{Column: column, Err: err})
	}
	return int(i)
}

func handlePanic(f reads, err *error) {
	r := recover()
	if r != nil {
		e, ok := r.(error)
		if !ok {
			panic(r)
		}
		if _, ok = r.(runtime.Error); ok {
			panic(r)
		}
		*err = e
	}
}

// This function cannot be used to create strings that are expected to persist.
func unsafeString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
