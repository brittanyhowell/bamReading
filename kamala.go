// attempt at reading a bam.
package main

import (
	"fmt"
	"os"

	"github.com/biogo/hts/bam"
)

func main() {

	f, err := os.Open("sample100.bam")
	// if err != nil {
	// 	return err
	// }
	// defer f.Close()

	// var br interface {
	// 	Read() (*sam.Record, error)
	// }
	// br, err = bam.NewReader(f, 0)
	// if err != nil {
	// 	return err
	// }
	// defer br.Close()

	var br *bam.Reader
	br, err = bam.NewReader(f, 0)
	if err != nil {
		return err
	}
	defer br.Close()

	// for {
	// 	r, err := br.Read()
	// 	if err != nil {
	// 		if err != io.EOF {
	// 			return err
	// 		}
	// 		break
	// 	}
	// }
	fmt.Printf("I am complete \n")
}
