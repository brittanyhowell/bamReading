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

	// helpful traits
	// cost := [...]float64{
	// 	sam.CigarDeletion: -1,
	// 	sam.CigarSkipped:  -1,
	// }

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
		// var (
		// 	scores []costPos
		// 	ref    = i.Record().Start()
		// 	query  int
		// )
		for i.Next() {
			//	fmt.Println(i.Record())
			// fmt.Printf("Start: %v End: %v \n", i.Record().Start()+1, i.Record().End()+1)
			fmt.Printf("cigar: %v\n", i.Record().Cigar)

			//fmt.Printf("scores: %v, ref: %v, query: %v, cost: %v \n", scores, ref, query, cost)
			// scores = append(scores, costPos{
			// 	ref:   ref,
			// 	query: query,
			// 	cost:  cost[],
			// })

			//cs := i.Record().Cigar

			// fmt.Printf("%q overlaps reference by %d letters\n", i.Record().Name, Overlap(i.Record(), i.Record().Start(), i.Record().Start()+i.Record().Len()))
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

// func Overlap(r *sam.Record, start, end int) int {
// 	var overlap int
// 	pos := r.Pos
// 	for _, co := range r.Cigar {
// 		t := co.Type()
// 		con := t.Consumes()
// 		fmt.Printf("Consumes: %v\n", con)
// 		lr := co.Len() * con.Reference
// 		if con.Query == con.Reference {
// 			o := min(pos+lr, end) - max(pos, start)
// 			if o > 0 {
// 				overlap += o
// 			}
// 		}
// 		pos += lr
// 	}

// 	return overlap
// }
