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

const (
	refStart = 0
	refEnd   = 45
)

type costPos struct {
	ref, query int
	cost       float64
}

func main() {

	// Read index
	ind, err := os.Open("sample5Change.bam.bai")
	if err != nil {
		log.Printf("error: could not open %s to read %v", ind, err)
	}
	defer ind.Close()
	bai, err := bam.ReadIndex(ind)
	h := bai.NumRefs()

	f, err := os.Open("sample5Change.bam")
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
		//	fmt.Printf("\n \n %v.  ref for %s is %#v\n", f, f.Chrom, refs[f.Chrom])
		fmt.Printf("\ncoord: %v\n", f)
		chunks, err := bai.Chunks(refs[f.Chrom], f.Start(), f.End())
		//	fmt.Printf("%#v -> %+v err:%v\n", f, chunks, err)

		i, err := bam.NewIterator(br, chunks)
		if err != nil {
			log.Fatal(err)
		}

		for i.Next() {

			r := i.Record()

			fmt.Printf("Start chunk: %v, end chunk: %v \n", f.Start(), f.End())
			fmt.Printf("Start record: %v, end record: %v \n", r.Start(), r.End())
			o := Overlaps(f.Start(), f.End(), r)
			if o == true {
				fmt.Printf("in chunk? %v \n", o)
			}
			if o == false {
				fmt.Printf("not in chunk, exiting, %v \n", o)
				continue
			}

			var hasDel bool
			for _, co := range r.Cigar {
				switch co.Type() {
				case sam.CigarSkipped, sam.CigarDeletion:
					hasDel = true
				}
			}
			if hasDel == true {
				fmt.Printf("Possible splice: %s, chromosome: %v start: %v, end: %v, length: %v \n", r.Name, f.Chrom, r.Pos, r.Pos+r.Len(), r.Len())
			}
			//fmt.Println(r.Cigar, hasDel)
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

}
func Overlaps(start int, end int, r *sam.Record) bool {
	return start < r.End()-50 && end > r.Start()+50
}
