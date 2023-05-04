package fuzz

import(
	search "github.com/Pauloo27/tuner/search"
	_ "fmt"
)

func Fuzz(fuzz_bytes []byte) int{
	sources := []search.SearchSource{search.SourceYouTube}
	//sources := []search.SearchSource{search.SourceSoundCloud}
	_ = search.Search(string(fuzz_bytes), 10, sources...)
	//fmt.Println("done")

	return 0

}