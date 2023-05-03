package fuzz

import(
	tea "github.com/Pauloo27/tuner/search"
)

func Fuzz(fuzz_bytes []byte) int{
	tea.Search(string(fuzz_bytes),10	)
	return 0
}