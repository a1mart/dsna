package main

import (
	"fmt"
	"sort"
)

type Suffix struct {
	index int
	suffix string
}

type SuffixArray []Suffix

func NewSuffixArray(text string) SuffixArray {
	suffixes := make([]Suffix, len(text))
	for i := range text {
		suffixes[i] = Suffix{i, text[i:]}
	}
	sort.Slice(suffixes, func(i, j int) bool {
		return suffixes[i].suffix < suffixes[j].suffix
	})
	return suffixes
}

func (sa SuffixArray) Index(i int) int {
	return sa[i].index
}

func (sa SuffixArray) Display() {
	for _, suffix := range sa {
		fmt.Printf("%d: %s\n", suffix.index, suffix.suffix)
	}
}

func main() {
	text := "banana"
	suffixArray := NewSuffixArray(text)
	fmt.Println("Suffix Array:")
	suffixArray.Display()
}
