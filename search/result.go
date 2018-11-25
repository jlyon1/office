package search

import (
	"github.com/jlyon1/office/episode"
)

// Result contains the results for a search, a single episode and the matching lines
type Result struct {
	Episode episode.Episode
	Lines   []string
	Rank    int
}
