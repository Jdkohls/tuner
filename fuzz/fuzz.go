package fuzz

import(
	search "github.com/Pauloo27/tuner/search"
)

func Fuzz(fuzz_bytes []byte) int{
	sources := []search.SearchSource{search.SourceYouTube}
	sources = append(sources, &search.SourceSoundCloud)
	result := search.Search(string(fuzz_bytes), 10, sources...)
	if result == nil{
		panic("Result is nil")
	}

	return 0

}