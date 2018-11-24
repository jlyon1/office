package search

import (
	"fmt"

	"github.com/jlyon1/office/episode"
	"github.com/jlyon1/office/loader"
	"github.com/sahilm/fuzzy"
)

// Episodes contains all of the episodes and allows querying by episode name
type Episodes struct {
	episodes []episode.Episode
}

// New returns a new episodes object
func New(episodes []episode.Episode) *Episodes {
	return &Episodes{
		episodes: episodes,
	}
}

// Load creates the episodes struct from a loader
func Load(l loader.Loader) (*Episodes, error) {
	r, err := l.LoadEpisodes()
	return &Episodes{r}, err
}

// Query searches every Episode for a string
func (e *Episodes) Query(pattern string) []Result {
	var ret []Result
	var data []string

	for i := range e.episodes {
		data = e.episodes[i].GetLines()

		matches := fuzzy.Find(pattern, data)
		if len(matches) > 0 {
			var r Result
			r.Episode = e.episodes[i]
			for _, match := range matches {
				if match.Score < 5000 {
					continue
				}
				fmt.Println(match.Str)
				r.Lines = append(r.Lines, match.Str)
			}
			fmt.Println()
			ret = append(ret, r)
		}
	}
	return ret
}

func contains(needle int, haystack []int) bool {
	for _, i := range haystack {
		if needle == i {
			return true
		}
	}
	return false
}
