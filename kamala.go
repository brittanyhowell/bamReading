// attempt at reading a bam.
package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
	"time"

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
	report      string
	index       string
	bamFile     string
	bedFile     string
	outPath     string
	outName     string
	genome      string
	seqOutName  string
	seqOut      *os.File
	logo5Name   string
	logo3Name   string
	logo5out    *os.File
	logo3out    *os.File
	readName    string
	readOut     *os.File
	readSumName string
	readSum     *os.File
	SJMap3      string
	SJMap5      string
	numSplice   int
	cSplice     int
)

func main() {
	flag.StringVar(&report, "report", "untitled_report.txt", "name of report file")
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
	flag.StringVar(&readSumName, "readSumName", "", "read summary file")
	flag.StringVar(&SJMap5, "SJMap5", "", "5' SJ Classification map")
	flag.StringVar(&SJMap3, "SJMap3", "", "3' SJ Classification map")
	flag.Parse()

	fmt.Println("Begin")

	// Make summary report file
	report := fmt.Sprintf("%v%v", outPath, report)
	rep, err := os.Create(report)
	if err != nil {
		log.Fatalf("failed to create readSum %s: %v", report, err)
	}
	defer rep.Close()

	fmt.Fprintf(rep, "Commence file initialisation at %s\n", time.Now())

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
		log.Fatalf("failed to create out %s: %v", file, err)
	}
	defer out.Close()

	if seqOutName != "" {
		fmt.Fprintf(rep, "File for full intron provided - will report full intron sequence in %s\n", seqOutName)
		seqFile := fmt.Sprintf("%v%v", outPath, seqOutName)
		seqOut, err = os.Create(seqFile)
		if err != nil {
			log.Fatalf("failed to create seqOut %s: %v", seqFile, err)
		}
		defer seqOut.Close()
	} else {
		fmt.Fprintf(rep, "File for full intron not provided - will not report full intron sequence\n")
	}

	if logo3Name != "" {
		fmt.Println("3' weblogo file provided - will produce 5' splicelogo fasta")
		threeFile := fmt.Sprintf("%v%v", outPath, logo3Name)
		logo3out, err = os.Create(threeFile)
		if err != nil {
			log.Fatalf("failed to create logo3out %s: %v", threeFile, err)
		}
		defer logo3out.Close()
	} else {
		fmt.Fprintf(rep, "File for 3' splice logo not provided - will not produce 3' Splice logo .fasta\n")
	}

	if logo5Name != "" {
		fmt.Println("5' weblogo file provided - will produce 5' splicelogo fasta")

		fiveFile := fmt.Sprintf("%v%v", outPath, logo5Name)
		logo5out, err = os.Create(fiveFile)
		fmt.Println(reflect.TypeOf(logo5out))
		if err != nil {
			log.Fatalf("failed to create fiveFile %s: %v", fiveFile, err)
		}
		defer logo5out.Close()
	} else {
		fmt.Fprintf(rep, "File for 5' splice logo not provided - will not produce 5' Splice logo .fasta\n")
	}

	if readName != "" {
		fmt.Fprintf(rep, "File for read record provided - will report all L1 relative read information in %s\n", readName)
		readFile := fmt.Sprintf("%v%v", outPath, readName)
		readOut, err = os.Create(readFile)
		if err != nil {
			log.Fatalf("failed to create readOut %s: %v", readFile, err)
		}
		defer readOut.Close()
	} else {
		fmt.Fprintf(rep, "File for read record not provided - will not produce L1 relative read information\n")
	}

	if readSumName != "" {
		fmt.Fprintf(rep, "File for read summary record provided - will create L1 read summary in %s\n", readSumName)
		readSumFile := fmt.Sprintf("%v%v", outPath, readSumName)
		readSum, err = os.Create(readSumFile)
		if err != nil {
			log.Fatalf("failed to create readSum %s: %v", readSumFile, err)
		}
		defer readSum.Close()
	} else {
		fmt.Fprintf(rep, "File for read summary not provided - will not produce L1 relative read information\n")
	}

	// read SJ data
	SpJu3, err := os.Open(SJMap3)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v.", err)
		os.Exit(1)
	}
	defer SpJu3.Close()

	scSJ3 := bufio.NewScanner(SpJu3)
	All3SJ := map[string]int{}

	var count int
	count = 1

	// scan into map
	for scSJ3.Scan() {
		sj3 := scSJ3.Text()
		All3SJ[sj3] = count
		count++
	}

	// read SJ data
	SpJu5, err := os.Open(SJMap5)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v.", err)
		os.Exit(1)
	}
	defer SpJu5.Close()

	scSJ5 := bufio.NewScanner(SpJu5)
	All5SJ := map[string]int{}

	//reset counter
	count = 1

	// scan into map
	for scSJ5.Scan() {
		sj5 := scSJ5.Text()
		All5SJ[sj5] = count
		count++
	}

	fmt.Fprintln(rep, "Loading genome")
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
	fmt.Fprintln(rep, "Genome loaded")

	//// Commence reading

	lr, err := bed.NewReader(loc, 3)
	if err != nil {
		log.Printf("error in NewReader: %s, %v", loc, err)
	}

	var numRead int

	fsc := featio.NewScanner(lr)
	for fsc.Next() {
		cSplice = 0 // reset spliced read count
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
			fmt.Printf("Overlaping read: %v, \t %v: Coords: %v-%v\n", r.Name, numRead, r.Start(), r.End())

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
				startGap = startInL1 + overlap
				endGap = startInL1 + overlap + gapLen

				switch co.Type() {
				case sam.CigarSkipped, sam.CigarInsertion:

					fSplice = true
					countSplice = true
					if gapLen > 4 {

						// genomic position of gap in read
						genStartGap := startGap + f.Start()
						genEndGap := endGap + f.Start()

						// genomic position of splice junction
						genStartSJ := genStartGap - 2
						genEndSJ := genEndGap - 2

						// Reading the nucleotides of the SJ in string form
						var buffer5 bytes.Buffer
						for i := genStartSJ; i < genStartSJ+4; i++ {
							buffer5.WriteString(string(AllSeqs[f.Chrom].At(i).L))
						}
						sFiveSJ := strings.ToUpper(buffer5.String())

						var buffer3 bytes.Buffer
						for i := genEndSJ; i < genEndSJ+2; i++ {
							buffer3.WriteString(string(AllSeqs[f.Chrom].At(i).L))
						}
						sThreeSJ := strings.ToUpper(buffer3.String())

						// Look SJ nucs up in maps
						fiveSJ := All5SJ[sFiveSJ]
						threeSJ := All3SJ[sThreeSJ]

						fmt.Println("The 5' class:", All5SJ[sFiveSJ], "Proof:", sFiveSJ)
						fmt.Println("The 3' class:", All3SJ[sThreeSJ], "Proof:", sThreeSJ)

						fmt.Fprintf(out, "%v \t%v \t %v \t %v \t %v \t %v \t %v \t %v \t %v \t %v \t %v\n",
							r.Name,    // read name
							f.Chrom,   // chromosome of L1
							f.Start(), // L1 genomic start
							f.End(),   // L1 genomic end
							startGap,  // start position of gap relative to L1
							endGap,    // end position of gap relative to L1
							fiveSJ,    // Class of 5' SJ
							threeSJ,   // Class of 3' SJ
							gapLen,    // length of gap
							r.Cigar,   // cigar string of read
							r.Flags,   // flags relative to read
						)

						// Include only if there is need for an intron
						if seqOutName != "" {
							nucs := AllSeqs[f.Chrom].Slice()

							fmt.Fprintf(seqOut, "%v \t%v \t %v \t %v \t %v \t %v \n",
								f.Chrom,   // chromosome of L1
								f.Start(), // L1 genomic start
								f.End(),   // L1 genomic end
								startGap,  // start position of gap relative to L1
								endGap,    // end position of gap relative to L1
								nucs.Slice(genStartGap-3, genEndGap+3), //
							)
						}

						// splice logo fasta files
						if logo5Name != "" {
							fmt.Fprintf(logo5out, ">Logo-5'%v:%v-%v\n%v\n",
								f.Chrom,  // chromosome name
								startGap, // start position of gap relative to L1
								endGap,   // end position of gap relative to L1
								sFiveSJ,  // letters at begin of splice
							)
						}

						if logo3Name != "" {
							fmt.Fprintf(logo3out, ">Logo-3'%v:%v-%v\n%v\n",
								f.Chrom,  // chromosome name
								startGap, // start position of gap relative to L1
								endGap,   // end position of gap relative to L1
								sThreeSJ, // letters at begin of splice
							)
						}
					}
					extra = gapLen // adds to overlap
				}
			}

			if readName != "" {
				fmt.Fprintf(readOut, "%v \t %v \t %v \t %v \t %v \t %v \t %v \t %v \n",
					f.Chrom,   // Chromosome of L1
					f.Start(), // Start position of L1
					f.End(),   // End position of L1
					startInL1, // Start position of read relative to L1
					endInL1,   // End position of read relative to L1
					r.Start(), // Start position of read relative to chromosome
					r.End(),   // End position of read relative to chromosome
					r.Cigar,   // Cigar string
				)
			}
			// Count number of spliced reads
			if countSplice {
				cSplice++
			}
		}

		err = i.Close()
		if err != nil {
			log.Fatal(err)
		}
		if fSplice {
			numSplice++
		}

		// Only print summary file of L1s which have at least a read mapped to them.
		if numRead > 0 {
			var pSplice float64
			pSplice = (float64(cSplice) / float64(numRead)) * 100.00
			fmt.Fprintf(readSum, "%v \t %v \t %v \t %v \t %v \t %v \t %.2f \n",
				f.Chrom,         // Chromosome of L1
				f.Start(),       // Start position of L1
				f.End(),         // End position of L1
				numRead,         // number of reads for that L1
				cSplice,         // number of spliced reads
				numRead-cSplice, // number of non-ggspliced reads
				pSplice,         // proportion spliced
			)
		}
	}
	err = fsc.Error()
	if err != nil {
		log.Fatalf("bed scan failed: %v", err)
	}
	fmt.Fprintf(rep, "There were %v intervals with at least one spliced read\n", numSplice)
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
