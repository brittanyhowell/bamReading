// attempt at reading a bam.
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/biogo/biogo/io/featio/bed"
	"github.com/biogo/hts/bam"
)

func main() {

	// Read index
	ind, err := os.Open("sample5.bam.bai")
	if err != nil {
		log.Printf("error: could not open %s to read %v", ind, err)
	}
	defer ind.Close()
	bai, err := bam.ReadIndex(ind)
	h := bai.NumRefs()
	fmt.Printf("h is h: %v \n\n", h)

	// Read location bed
	loc, err := os.Open("tiny.bed")
	if err != nil {
		log.Printf("error: could not open %s to read %v", loc, err)
	}
	defer loc.Close()

	lr, err := bed.NewReader(loc, 3)
	if err != nil {
		log.Printf("error in NewReader: %s, %v", loc, err)
	}

	l, err := lr.Read()
	fmt.Printf("Locations: %v\n", l)

	//bam.Chunks()

	f, err := os.Open("sample5.bam")
	if err != nil {
		log.Printf("error: could not open %s to read %v", f, err)
	}
	defer f.Close()

	var br *bam.Reader
	br, err = bam.NewReader(f, 0)
	if err != nil {
		log.Printf("error: %s, %v", br, err)
	}
	defer br.Close()

	r, err := br.Read()
	if err != nil {
		log.Printf("error: %s, %v", r, err)
	}

	hat := r.Cigar

	fmt.Printf("hat: %v \n\n", hat)

	//fmt.Printf("BR: %v \n\n", br)
	//fmt.Printf("R: %v \n \n", r)

}
