// attempt at reading a bam.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/biogo/biogo/alphabet"
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
			if overlaps(r, f) {
			} else {
				continue
			}

			var gapLen, overlap, extra, readOverlap int
			for _, co := range r.Cigar {

				pos := r.Pos
				t := co.Type()
				con := t.Consumes()
				gapLen = co.Len() * con.Reference

				if con.Query == con.Reference {
					o := min(pos+gapLen, r.End()) - max(pos, r.Start())
					if o > 0 {
						overlap += o
						readOverlap += o
					}
				}
				overlap += extra
				extra = 0
				switch co.Type() {
				case sam.CigarSkipped, sam.CigarDeletion:

					startInL1 := r.Start() - f.Start()
					fmt.Printf("\n\nPossible splice: %v \tL1: %v:%v-%v \t Start: %v \tEnd: %v \tLength: %v\n", r.Name, f.Chrom, f.Start(), f.End(), startInL1+overlap, startInL1+overlap+gapLen, gapLen)
					startGap := startInL1 + overlap
					endGap := startInL1 + overlap + gapLen
					if gapLen > 4 && gapLen < 2000 {
						seq := r.Seq.Expand()
						notReallyLetter := alphabet.BytesToLetters(seq)
						letter := alphabet.Letters(notReallyLetter)
						beginsplice := letter.Slice(readOverlap-2, readOverlap)
						endSplice := letter.Slice(readOverlap, readOverlap+2)
						fmt.Fprintf(out, "%v \t%v \t %v \t %v \t %v \t %v \t %v \t %v \t %v \t %v \t %v\n", r.Name, f.Chrom, f.Start(), f.End(), startGap, endGap, beginsplice, endSplice, gapLen, r.Cigar, r.Flags)

						fmt.Printf("Begin splice: %v, End splice:%v\n", beginsplice, endSplice)
						fmt.Printf("Alignment coordinate information: %v, readOverlap: %v, Start: %v, Length: %v\n", r.Cigar, readOverlap, startGap, gapLen)
					}
					extra = gapLen // adds to overlap
				}
			}
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
	// read must be entirely within L1 interval
	return r.Start() > f.Start() && r.End() < f.End()
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
