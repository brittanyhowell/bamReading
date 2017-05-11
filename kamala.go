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
	"github.com/biogo/biogo/io/seqio"
	"github.com/biogo/biogo/io/seqio/fasta"
	"github.com/biogo/biogo/seq/linear"
	"github.com/biogo/hts/bam"
	"github.com/biogo/hts/sam"
)

var (
	index      string
	bamFile    string
	bedFile    string
	outPath    string
	outName    string
	genome     string
	seqOutName string
	logo5Name  string
	logo3Name  string
	readName   string
	numSplice  int
	cSplice    int
)

func main() {
	flag.StringVar(&index, "index", "", "name index file")
	flag.StringVar(&bamFile, "bam", "", "name bam file")
	flag.StringVar(&bedFile, "intervalsBed", "", "BED3 of required intervals")
	flag.StringVar(&outPath, "outPath", "", "path to output dir")
	flag.StringVar(&outName, "outName", "", "out file name")
	flag.StringVar(&seqOutName, "seqOutName", "", "sequence containing out file name")
	flag.StringVar(&logo5Name, "logo5Name", "", "sequence containing webLogo 5' file name")
	flag.StringVar(&logo3Name, "logo3Name", "", "sequence containing webLogo 3' file name")
	flag.StringVar(&genome, "refGen", "", "reference genome")
	flag.StringVar(&readName, "readName", "", "read information file")
	flag.Parse()

	fmt.Println("Loading genome")
	gen, err := os.Open(genome)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v.", err)
		os.Exit(1)
	}
	defer gen.Close()

	in := fasta.NewReader(gen, linear.NewSeq("", nil, alphabet.DNA))
	sc := seqio.NewScanner(in)
	AllSeqs := map[string]*linear.Seq{}

	for sc.Next() {
		s := sc.Seq().(*linear.Seq)

		AllSeqs[s.Name()] = s
	}
	fmt.Println("Genome loaded")

	// read index
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

	// Creating files for the output
	file := fmt.Sprintf("%v%v", outPath, outName)
	out, err := os.Create(file)
	if err != nil {
		log.Fatalf("failed to create %s: %v", file, err)
	}
	defer out.Close()

	seqFile := fmt.Sprintf("%v%v", outPath, seqOutName)
	seqOut, err := os.Create(seqFile)
	if err != nil {
		log.Fatalf("failed to create %s: %v", seqFile, err)
	}
	defer seqOut.Close()

	threeFile := fmt.Sprintf("%v%v", outPath, logo3Name)
	logo3out, err := os.Create(threeFile)
	if err != nil {
		log.Fatalf("failed to create %s: %v", threeFile, err)
	}
	defer out.Close()

	fiveFile := fmt.Sprintf("%v%v", outPath, logo5Name)
	logo5out, err := os.Create(fiveFile)
	if err != nil {
		log.Fatalf("failed to create %s: %v", fiveFile, err)
	}
	defer out.Close()

	readFile := fmt.Sprintf("%v%v", outPath, readName)
	readOut, err := os.Create(readFile)
	if err != nil {
		log.Fatalf("failed to create %s: %v", readFile, err)
	}
	defer out.Close()

	//// Commence reading

	lr, err := bed.NewReader(loc, 3)
	if err != nil {
		log.Printf("error in NewReader: %s, %v", loc, err)
	}

	var numRead int

	fsc := featio.NewScanner(lr)
	for fsc.Next() {

		numRead = 0 // reset number of reads in element

		f := fsc.Feat().(*bed.Bed3)
		fmt.Printf("\nL1: %v - %v\n", f.Start(), f.End())

		fSplice := false // if element has a spliced read in it, it becomes positive
		countSplice := false

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
			numRead++
			fmt.Printf("Read: %v, \t %v: Coords: %v-%v\n", r.Name, numRead, r.Start(), r.End())

			var (
				startInL1 int
				endInL1   int
				startGap  int
				endGap    int

				gapLen      int
				overlap     int
				extra       int
				readOverlap int
			)
			for _, co := range r.Cigar {

				cSplice = 0 // reset spliced read count

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

				startInL1 = r.Start() - f.Start()
				endInL1 = r.End() - f.Start()
				//	fmt.Printf("Possible splice: \tL1: %v:%v-%v \t Start: %v \tEnd: %v \tLength: %v \t%v\n",
				//	f.Chrom, f.Start(), f.End(), startInL1+overlap, startInL1+overlap+gapLen, gapLen, r.Cigar)
				startGap = startInL1 + overlap
				endGap = startInL1 + overlap + gapLen

				switch co.Type() {
				case sam.CigarSkipped, sam.CigarInsertion:
					fSplice = true
					countSplice = true
					if gapLen > 4 {

						seq := r.Seq.Expand()
						letter := alphabet.Letters(alphabet.BytesToLetters(seq))
						beginsplice := letter.Slice(readOverlap-2, readOverlap)
						endSplice := letter.Slice(readOverlap, readOverlap+2)

						genStartGap := startGap + f.Start()
						genEndGap := endGap + f.Start()
						nucs := AllSeqs[f.Chrom].Slice()
						fiveSJ := nucs.Slice(genStartGap-3, genStartGap+3)
						threeSJ := nucs.Slice(genEndGap-3, genEndGap+3)
						fmt.Fprintf(out, "%v \t%v \t %v \t %v \t %v \t %v \t %v \t %v \t %v \t %v \t %v \t %v \t %v\n",
							r.Name,      // read name
							f.Chrom,     // chromosome of L1
							f.Start(),   // L1 genomic start
							f.End(),     // L1 genomic end
							startGap,    // start position of gap relative to L1
							endGap,      // end position of gap relative to L1
							fiveSJ,      // letters at begin of splice
							threeSJ,     // nucs at end
							beginsplice, // two nucs in read at 5' end of SJ
							endSplice,   // two nucs in reads at 3' end of SJ
							gapLen,      // length of gap
							r.Cigar,     // cigar string of read
							r.Flags,     // flags relative to read
						)

						fmt.Fprintf(seqOut, "%v \t%v \t %v \t %v \n",
							f.Chrom,   // chromosome of L1
							f.Start(), // L1 genomic start
							f.End(),   // L1 genomic end
							startGap,  // start position of gap relative to L1
							endGap,    // end position of gap relative to L1
							nucs.Slice(genStartGap-3, genEndGap+3), //
						)

						fmt.Fprintf(logo5out, ">Logo-5'%v:%v-%v\n%v\n",
							f.Chrom,  // chromosome name
							startGap, // start position of gap relative to L1
							endGap,   // end position of gap relative to L1
							fiveSJ,   // letters at begin of splice
						)

						fmt.Fprintf(logo3out, ">Logo-3'%v:%v-%v\n%v\n",
							f.Chrom,  // chromosome name
							startGap, // start position of gap relative to L1
							endGap,   // end position of gap relative to L1
							threeSJ,  // letters at begin of splice
						)
					}
					extra = gapLen // adds to overlap
				}
			}
			fmt.Printf("Read information: L1 relative: %v\t%v\t genome relative: %v\t%v\t%v\n", startInL1, endInL1, f.Chrom, r.Start(), r.End())

			fmt.Fprintf(readOut, "%v\t%v\t%v\t%v\t%v\n",
				startInL1,
				endInL1,
				f.Chrom,
				r.Start(),
				r.End(),
			)
			if countSplice {
				cSplice++
			}
		}
		fmt.Printf("There were %v reads for that L1 (%v - %v)(%v were spliced)\n", numRead, f.Start(), f.End(), cSplice)
		err = i.Close()
		if err != nil {
			log.Fatal(err)
		}
		if fSplice {
			numSplice++
		}
	}
	err = fsc.Error()
	if err != nil {
		log.Fatalf("bed scan failed: %v", err)
	}
	fmt.Printf("There were %v intervals with at least one spliced read\n", numSplice)
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
