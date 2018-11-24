package episode

import (
	"strconv"
	"strings"
)

// Episode represents all of the lines in a single episode of the office. Takes a byte array of
// the script of the episode and separates it by the punctuation.
type Episode struct {
	Number int
	Season int
	lines  []string
}

// New parses the byte array and returns a new episode
func New(lines []byte, season int, episode int) *Episode {
	lstr := strings.FieldsFunc(string(lines), split)
	ret := &Episode{
		Number: episode,
		Season: season,
		lines:  lstr,
	}

	return ret
}

// GetReadableName returns the name of the episode in a readable format
func (e *Episode) GetReadableName() string {
	return "Season " + strconv.Itoa(e.Season) + " Episode " + strconv.Itoa(e.Number)
}

// GetNumLines returns the number of lines loaded from the episode
func (e *Episode) GetNumLines() int {
	return len(e.lines)
}

func (e *Episode) GetLines() []string {
	var ret []string
	for i := 0; i < len(e.lines); i++ {
		if len(e.lines[i]) > 0 && len(e.lines[i]) < 2000 {
			ret = append(ret, e.lines[i])
		}
	}
	return ret
}

func split(r rune) bool {
	return r == '.' || r == '!' || r == '?' || r == '"' || r == ')' || r == '('
}
