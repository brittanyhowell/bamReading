package main

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/biogo/hts/sam"
)

func main() {

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

				fmt.Fprintf(out, "%v \t%v \t %v \t %v \t %v \t %v \t %v \t %v \t %v \t %v \t %v \t %v \t %v\n",
					r.Name,    // read name
					f.Chrom,   // chromosome of L1
					f.Start(), // L1 genomic start
					f.End(),   // L1 genomic end
					startGap,  // start position of gap relative to L1
					endGap,    // end position of gap relative to L1
					fiveSJ,    // Class of 5' SJ
					threeSJ,   // Class of 3' SJ
					sFiveSJ,   // 5' nucs
					sThreeSJ,  // 3' nucs
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

}
