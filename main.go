package main

import (
	"fmt"
	"github.com/toukii/jsnm"
)

func main() {
	jm := jsnm.FileNameFmt("city.md")
	provinces := jm.Get("p").Arr()
	cities := jm.Get("c").Arr()
	for i, it := range provinces {
		fmt.Printf("***** %dth: %v\n", i, it.RawData().String())
		ccs := cities[i].Arr()
		for i, it := range ccs {
			fmt.Printf("%dth: %v\t", i, it.RawData().String())
		}
		fmt.Println()
	}

	code_jm := jsnm.FileNameFmt("code.md")
	bjc := code_jm.Get("北京").RawData().String()
	fmt.Printf("Beijing's code:%s\n", bjc)
}
