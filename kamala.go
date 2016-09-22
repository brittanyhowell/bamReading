// attempt at reading a bam.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/biogo/biogo/feat"
	"github.com/biogo/biogo/io/featio"
	"github.com/biogo/biogo/io/featio/bed"
	"github.com/biogo/hts/bam"
	"github.com/biogo/hts/sam"
)

var (
	index   string
	bamFile string
	bedFile string
	outPath string
	outName string
)

func main() {
	flag.StringVar(&index, "index", "", "name index file")
	flag.StringVar(&bamFile, "bam", "", "name bam file")
	flag.StringVar(&bedFile, "intervalsBed", "", "BED3 of required intervals")
	flag.StringVar(&outPath, "outPath", "", "path to output dir")
	flag.StringVar(&outName, "outName", "", "out file name")
	flag.Parse()

	// Read index
	ind, err := os.Open(index)
	if err != nil {
		log.Printf("error: could not open %s to read %v", ind, err)
	}
	defer ind.Close()
	bai, err := bam.ReadIndex(ind)
	h := bai.NumRefs()

	// Read bam
	f, err := os.Open(bamFile)
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

	// store bams
	refs := make(map[string]*sam.Reference, h)
	for _, r := range br.Header().Refs() {
		refs[r.Name()] = r
	}

	// Read location file
	loc, err := os.Open(bedFile)
	if err != nil {
		log.Printf("error: could not open %s to read %v", loc, err)
	}
	defer loc.Close()

	// Creating a file for the output
	file := fmt.Sprintf("%v%v", outPath, outName)
	out, err := os.Create(file)
	if err != nil {
		log.Fatalf("failed to create %s: %v", file, err)
	}
	defer out.Close()

	lr, err := bed.NewReader(loc, 3)
	if err != nil {
		log.Printf("error in NewReader: %s, %v", loc, err)
	}

	fsc := featio.NewScanner(lr)
	for fsc.Next() {
		f := fsc.Feat().(*bed.Bed3)
		//	fmt.Printf("Start L1 interval: %v, end L1 interval: %v \n", f.Start(), f.End())
		// set chunks

		chunks, err := bai.Chunks(refs[f.Chrom], f.Start(), f.End())
		if err != nil {
			fmt.Println(chunks, err)
			continue
		}

		i, err := bam.NewIterator(br, chunks)
		if err != nil {
			log.Fatal(err)
		}
		// iterate over reads
		for i.Next() {

			r := i.Record()

			//		fmt.Printf("Start record: %v, end record: %v \n", r.Start(), r.End())
			// check overlap is significant
			if overlaps(r, f) {
				//	fmt.Printf("in interval \n")
			} else {
				//	fmt.Printf("not in interval, skipping\n")
				continue
			}

			// var hasDel bool
			var gapLen, overlap, extra int

			for _, co := range r.Cigar {
				pos := r.Pos
				t := co.Type()
				con := t.Consumes()
				gapLen = co.Len() * con.Reference

				if con.Query == con.Reference {
					o := min(pos+gapLen, r.End()) - max(pos, r.Start())
					if o > 0 {
						overlap += o
					}
				}
				overlap += extra
				extra = 0
				switch co.Type() {
				case sam.CigarSkipped, sam.CigarDeletion:

					startInL1 := r.Start() - f.Start()
					//	endInL1 := startInL1 + r.Len()
					// fmt.Printf("Read position: %v length of gap,: %v \n", overlap, gapLen)
					// fmt.Printf("%s \t %d \t %d \t %s\n", f.Chrom, startInL1, endInL1, r.Cigar)
					// fmt.Printf("Start L1 interval: %v, end L1 interval: %v \n", f.Start(), f.End())
					fmt.Printf("L1: %v:%v-%v \tPossible splice: Start: %v \tEnd: %v \tLength: %v\n", f.Chrom, f.Start(), f.End(), startInL1+overlap, startInL1+overlap+gapLen, gapLen)
					startGap := startInL1 + overlap
					endGap := startInL1 + overlap + gapLen
					if gapLen > 4 && gapLen < 2000 {
						fmt.Fprintf(out, "%v \t %v \t %v \t %v \t %v \t %v\n", f.Chrom, f.Start(), f.End(), startGap, endGap, gapLen)
					}
					extra = gapLen
				}
			}
			// if hasDel {

			// fmt.Printf("f%v\n", f.Chrom)
			//fmt.Printf("Possible splice: %s, chromosome: %v start: %v, end: %v, length: %v \n", r.Name, f.Chrom, r.Pos, r.Pos+r.Len(), r.Len())
			//if endInL1 < 8000 && startInL1 > 0 {
			//fmt.Println(r.Cigar)

			// fmt.Printf("Possible splice: %s, chromosome: %v start: %v, end: %v, length: %v \n", r.Name, f.Chrom, startInL1, endInL1, r.Len())
			// fmt.Fprintf(out, "%s \t %d \t %d \t %s\n", f.Chrom, startInL1, endInL1, r.Cigar)
			//}
			// }

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
func overlaps(r *sam.Record, f feat.Feature) bool {
	// return f.Start() < r.End()-50 && f.End() > r.Start()+50
	return r.Start() > f.Start() && r.End() < f.End()
}

// func min(a, b int) int {
// 	if a > b {
// 		return b
// 	}
// 	return a
// }

// func max(a, b int) int {
// 	if a < b {
// 		return b
// 	}
// 	return a
// }

// func split(r *sam.Record, start, end int) int {

// 	var overlap int
// 	pos := r.Pos
// 	for _, co := range r.Cigar {
// 		t := co.Type()
// 		con := t.Consumes()
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
