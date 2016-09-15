// attempt at reading a bam.
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/biogo/biogo/io/featio"
	"github.com/biogo/biogo/io/featio/bed"
	"github.com/biogo/hts/bam"
	"github.com/biogo/hts/sam"
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
	refs := make(map[string]*sam.Reference, h)
	for _, r := range br.Header().Refs() {
		refs[r.Name()] = r
	}

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

	fsc := featio.NewScanner(lr)
	for fsc.Next() {
		f := fsc.Feat().(*bed.Bed3)
		fmt.Printf("NUM: %v. WORD. ref for %s is %#v\n", f, f.Chrom, refs[f.Chrom])

		chunks, err := bai.Chunks(refs[f.Chrom], f.Start(), f.End())
		fmt.Printf("%#v -> %+v %v\n", f, chunks, err)
		i, err := bam.NewIterator(br, chunks)
		if err != nil {
			log.Fatal(err)
		}
		for i.Next() {
			fmt.Println(i.Record())
		}
		err = i.Close()
		if err != nil {
			log.Fatal(err)
		}
	}
	err = fsc.Error()
	if err != nil {
		log.Fatalf("bed scan failed: %v", err)
	}

	//bam.Chunks()

	//fmt.Printf("BR: %v \n\n", br)
	//fmt.Printf("R: %v \n \n", r)

}
