package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type reads struct {
	chr    string
	gStart int
	// gEnd   int
	// lStart int
	// lEnd   int
	// rStart int
	// rEnd   int
	// Cigar  sam.Cigar
}

var (
	readFile string
)

func main() {

	flag.StringVar(&readFile, "readFile", "", "File with read data")
	flag.Parse()

	f, err := os.Open(readFile)
	if err != nil {
		log.Printf("error: could not open %s to read %v", f, err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	// sc.Split(bufio.ScanLines)
	i := 0

	// block := []string{}

	for sc.Scan() {

		line := sc.Text()
		line = strings.TrimSpace(line)
		words := strings.Split(line, " ")

		word, err := strconv.Atoi(words[1])
		if err != nil {
			os.Exit(1)

		}

		fmt.Printf("words: %v  (%T)\n word: %v (%T)\n", words[1], words[1], word, word)
		dat := reads{chr: words[0], gStart: word}
		// block = append(block, line)
		i++
		// fmt.Println("So far:", block, i)
		fmt.Println(dat)
	}

	// fmt.Println("a block: ", block)

	// scS := seqio.NewScanner(r)
	// for scS.Next() {
	// 	v = append(v, scS.Seq())
	// }
	// if sc.Error() != nil {
	// 	log.Fatalf("failed to read sequences: %v", sc.Error())
	// }

	// for _, co := range r.Cigar {

	// 	pos := r.Pos
	// 	t := co.Type()
	// 	con := t.Consumes()
	// 	gapLen = co.Len() * con.Reference

	// 	if con.Query == con.Reference {
	// 		o := min(pos+gapLen, r.End()) - max(pos, r.Start())
	// 		if o > 0 {
	// 			overlap += o
	// 		}
	// 	}
	// 	overlap += extra
	// 	extra = 0

	// 	startInL1 = r.Start() - f.Start()
	// 	endInL1 = r.End() - f.Start()
	// 	startGap = startInL1 + overlap
	// 	endGap = startInL1 + overlap + gapLen

	// 	switch co.Type() {
	// 	case sam.CigarSkipped, sam.CigarInsertion:

	// 		extra = gapLen // adds to overlap
	// 	}
	// }

}
