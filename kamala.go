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
)

func main() {
	flag.StringVar(&index, "index", "", "name index file")
	flag.StringVar(&bamFile, "bam", "", "name bam file")
	flag.StringVar(&bedFile, "intervalsBed", "", "BED3 of required intervals")
	flag.StringVar(&outPath, "outPath", "", "path to output dir")
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
	file := fmt.Sprintf("%vSplitReads.bed", outPath)
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
		fmt.Printf("\nStart L1 interval: %v, end L1 interval: %v \n", f.Start(), f.End())
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

			fmt.Printf("Start record: %v, end record: %v \n", r.Start(), r.End())
			// check overlap is significant
			if overlaps(r, f) {
				//	fmt.Printf("in interval \n")
			} else {
				//	fmt.Printf("not in interval, skipping\n")
				continue
			}

			var hasDel bool
			for _, co := range r.Cigar {
				switch co.Type() {
				case sam.CigarSkipped, sam.CigarDeletion:
					hasDel = true
				}
			}
			if hasDel {
				fmt.Printf("Possible splice: %s, chromosome: %v start: %v, end: %v, length: %v \n", r.Name, f.Chrom, r.Pos, r.Pos+r.Len(), r.Len())
				fmt.Fprintf(out, "%s \t %d \t %d \n", f.Chrom, r.Pos, r.Pos+r.Len())
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
	return f.Start() < r.End()-50 && f.End() > r.Start()+50
}
