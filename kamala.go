// attempt at reading a bam.
package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"reflect"

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
	index       string
	bamFile     string
	bedFile     string
	outPath     string
	outName     string
	genome      string
	seqOutName  string
	logo5Name   string
	logo3Name   string
	readName    string
	readSumName string
	numSplice   int
	cSplice     int
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
	flag.StringVar(&readSumName, "readSumName", "", "read summary file")
	flag.Parse()

	fmt.Println("Begin")

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

	seqFile := fmt.Sprintf("%v%v", outPath, seqOutName)
	seqOut, err := os.Create(seqFile)
	if err != nil {
		log.Fatalf("failed to create seqOut %s: %v", seqFile, err)
	}
	defer seqOut.Close()

	threeFile := fmt.Sprintf("%v%v", outPath, logo3Name)
	logo3out, err := os.Create(threeFile)
	if err != nil {
		log.Fatalf("failed to create logo3out %s: %v", threeFile, err)
	}
	defer out.Close()

	fiveFile := fmt.Sprintf("%v%v", outPath, logo5Name)
	logo5out, err := os.Create(fiveFile)
	if err != nil {
		log.Fatalf("failed to create fiveFile %s: %v", fiveFile, err)
	}
	defer out.Close()

	readFile := fmt.Sprintf("%v%v", outPath, readName)
	readOut, err := os.Create(readFile)
	if err != nil {
		log.Fatalf("failed to create readOut %s: %v", readFile, err)
	}
	defer out.Close()

	readSumFile := fmt.Sprintf("%v%v", outPath, readSumName)
	readSum, err := os.Create(readSumFile)
	if err != nil {
		log.Fatalf("failed to create readSum %s: %v", readSumFile, err)
	}
	defer out.Close()

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

						genStartSJ := genStartGap - 2

						var buffer bytes.Buffer
						for i := genStartSJ; i < genStartSJ+4; i++ {

							buffer.WriteString(string(AllSeqs[f.Chrom].At(i).L))
						}
						sFiveSJ := buffer.String()
						fmt.Println(sFiveSJ)

						sjscore(sFiveSJ)
						// fmt.Println()
						// fmt.Println("Whatever this is", genStartGap)
						// // newNucs := string(AllSeqs[f.Chrom].At(genStartGap).L)
						// newNucs := string(AllSeqs[f.Chrom].At(2).L)
						// fmt.Printf("Whatever this extra thing is: %v, type: %v \n", newNucs, reflect.TypeOf(newNucs))
						// fmt.Println()

						nucs := AllSeqs[f.Chrom].Slice()
						fiveSJ := nucs.Slice(genStartGap-2, genStartGap+2)
						threeSJ := nucs.Slice(genEndGap-2, genEndGap+1)

						sFSJ := fiveSJ
						fmt.Println(sFSJ)

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

						// sjscore(fiveSJ)
						fmt.Printf("the fiveSJ: %v\n", fiveSJ)

						fmt.Println(reflect.TypeOf(fiveSJ))

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
			// fmt.Printf("Read information: L1 relative: %v\t%v\t genome relative: %v\t%v\t%v\n", startInL1, endInL1, f.Chrom, r.Start(), r.End())

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

		fmt.Printf("There were %v reads for that L1 (%v - %v)(%v were spliced)\n", numRead, f.Start(), f.End(), cSplice)
		var pSplice float64
		pSplice = (float64(cSplice) / float64(numRead)) * 100.00
		fmt.Fprintf(readSum, "%v \t %v \t %v \t %v \t %v \t %v \t %.2f \n", // FIX ME
			f.Chrom,         // Chromosome of L1
			f.Start(),       // Start position of L1
			f.End(),         // End position of L1
			numRead,         // number of reads for that L1
			cSplice,         // number of spliced reads
			numRead-cSplice, // number of non-spliced reads
			pSplice,         // proportion spliced
		)
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

func sjscore(five string) (int, error) {
	// str := alphabet.LettersToBytes(five)

	var sjClass = map[string]int{

		"AAAA": 1,
		"AAAC": 2,
		"AAAT": 3,
		"AAAG": 4,
		"AACA": 5,
		"AACC": 6,
		"AACT": 7,
		"AACG": 8,
		"AATA": 9,
		"AATC": 10,
		"AATT": 11,
		"AATG": 12,
		"AAGA": 13,
		"AAGC": 14,
		"AAGT": 15,
		"AAGG": 16,
		"ACAA": 17,
		"ACAC": 18,
		"ACAT": 19,
		"ACAG": 20,
		"ACCA": 21,
		"ACCC": 22,
		"ACCT": 23,
		"ACCG": 24,
		"ACTA": 25,
		"ACTC": 26,
		"ACTT": 27,
		"ACTG": 28,
		"ACGA": 29,
		"ACGC": 30,
		"ACGT": 31,
		"ACGG": 32,
		"ATAA": 33,
		"ATAC": 34,
		"ATAT": 35,
		"ATAG": 36,
		"ATCA": 37,
		"ATCC": 38,
		"ATCT": 39,
		"ATCG": 40,
		"ATTA": 41,
		"ATTC": 42,
		"ATTT": 43,
		"ATTG": 44,
		"ATGA": 45,
		"ATGC": 46,
		"ATGT": 47,
		"ATGG": 48,
		"AGAA": 49,
		"AGAC": 50,
		"AGAT": 51,
		"AGAG": 52,
		"AGCA": 53,
		"AGCC": 54,
		"AGCT": 55,
		"AGCG": 56,
		"AGTA": 57,
		"AGTC": 58,
		"AGTT": 59,
		"AGTG": 60,
		"AGGA": 61,
		"AGGC": 62,
		"AGGT": 63,
		"AGGG": 64,
		"CAAA": 65,
		"CAAC": 66,
		"CAAT": 67,
		"CAAG": 68,
		"CACA": 69,
		"CACC": 70,
		"CACT": 71,
		"CACG": 72,
		"CATA": 73,
		"CATC": 74,
		"CATT": 75,
		"CATG": 76,
		"CAGA": 77,
		"CAGC": 78,
		"CAGT": 79,
		"CAGG": 80,
		"CCAA": 81,
		"CCAC": 82,
		"CCAT": 83,
		"CCAG": 84,
		"CCCA": 85,
		"CCCC": 86,
		"CCCT": 87,
		"CCCG": 88,
		"CCTA": 89,
		"CCTC": 90,
		"CCTT": 91,
		"CCTG": 92,
		"CCGA": 93,
		"CCGC": 94,
		"CCGT": 95,
		"CCGG": 96,
		"CTAA": 97,
		"CTAC": 98,
		"CTAT": 99,
		"CTAG": 100,
		"CTCA": 101,
		"CTCC": 102,
		"CTCT": 103,
		"CTCG": 104,
		"CTTA": 105,
		"CTTC": 106,
		"CTTT": 107,
		"CTTG": 108,
		"CTGA": 109,
		"CTGC": 110,
		"CTGT": 111,
		"CTGG": 112,
		"CGAA": 113,
		"CGAC": 114,
		"CGAT": 115,
		"CGAG": 116,
		"CGCA": 117,
		"CGCC": 118,
		"CGCT": 119,
		"CGCG": 120,
		"CGTA": 121,
		"CGTC": 122,
		"CGTT": 123,
		"CGTG": 124,
		"CGGA": 125,
		"CGGC": 126,
		"CGGT": 127,
		"CGGG": 128,
		"TAAA": 129,
		"TAAC": 130,
		"TAAT": 131,
		"TAAG": 132,
		"TACA": 133,
		"TACC": 134,
		"TACT": 135,
		"TACG": 136,
		"TATA": 137,
		"TATC": 138,
		"TATT": 139,
		"TATG": 140,
		"TAGA": 141,
		"TAGC": 142,
		"TAGT": 143,
		"TAGG": 144,
		"TCAA": 145,
		"TCAC": 146,
		"TCAT": 147,
		"TCAG": 148,
		"TCCA": 149,
		"TCCC": 150,
		"TCCT": 151,
		"TCCG": 152,
		"TCTA": 153,
		"TCTC": 154,
		"TCTT": 155,
		"TCTG": 156,
		"TCGA": 157,
		"TCGC": 158,
		"TCGT": 159,
		"TCGG": 160,
		"TTAA": 161,
		"TTAC": 162,
		"TTAT": 163,
		"TTAG": 164,
		"TTCA": 165,
		"TTCC": 166,
		"TTCT": 167,
		"TTCG": 168,
		"TTTA": 169,
		"TTTC": 170,
		"TTTT": 171,
		"TTTG": 172,
		"TTGA": 173,
		"TTGC": 174,
		"TTGT": 175,
		"TTGG": 176,
		"TGAA": 177,
		"TGAC": 178,
		"TGAT": 179,
		"TGAG": 180,
		"TGCA": 181,
		"TGCC": 182,
		"TGCT": 183,
		"TGCG": 184,
		"TGTA": 185,
		"TGTC": 186,
		"TGTT": 187,
		"TGTG": 188,
		"TGGA": 189,
		"TGGC": 190,
		"TGGT": 191,
		"TGGG": 192,
		"GAAA": 193,
		"GAAC": 194,
		"GAAT": 195,
		"GAAG": 196,
		"GACA": 197,
		"GACC": 198,
		"GACT": 199,
		"GACG": 200,
		"GATA": 201,
		"GATC": 202,
		"GATT": 203,
		"GATG": 204,
		"GAGA": 205,
		"GAGC": 206,
		"GAGT": 207,
		"GAGG": 208,
		"GCAA": 209,
		"GCAC": 210,
		"GCAT": 211,
		"GCAG": 212,
		"GCCA": 213,
		"GCCC": 214,
		"GCCT": 215,
		"GCCG": 216,
		"GCTA": 217,
		"GCTC": 218,
		"GCTT": 219,
		"GCTG": 220,
		"GCGA": 221,
		"GCGC": 222,
		"GCGT": 223,
		"GCGG": 224,
		"GTAA": 225,
		"GTAC": 226,
		"GTAT": 227,
		"GTAG": 228,
		"GTCA": 229,
		"GTCC": 230,
		"GTCT": 231,
		"GTCG": 232,
		"GTTA": 233,
		"GTTC": 234,
		"GTTT": 235,
		"GTTG": 236,
		"GTGA": 237,
		"GTGC": 238,
		"GTGT": 239,
		"GTGG": 240,
		"GGAA": 241,
		"GGAC": 242,
		"GGAT": 243,
		"GGAG": 244,
		"GGCA": 245,
		"GGCC": 246,
		"GGCT": 247,
		"GGCG": 248,
		"GGTA": 249,
		"GGTC": 250,
		"GGTT": 251,
		"GGTG": 252,
		"GGGA": 253,
		"GGGC": 254,
		"GGGT": 255,
		"GGGG": 256,
	}
	return fmt.Println("Key: ", sjClass[five])

}
