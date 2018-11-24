package loader

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/jlyon1/office/episode"
)

// Loader loads scripts into episodes and seasons
type Loader struct {
	directory string
}

// New sets a new directory to load the script from and returns the loader
func New(dir string) *Loader {

	return &Loader{
		directory: dir,
	}
}

// LoadEpisodes loads in every episode in the directory assuming s[n]e[n].txt format and sets the values accordingly
func (l *Loader) LoadEpisodes() ([]episode.Episode, error) {
	root := l.directory
	var episodes []episode.Episode
	var files []string

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".txt" {
			files = append(files, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			return nil, err
		}
		val := make([]byte, 20000)
		_, err = f.Read(val)
		if err != nil {
			continue
		}
		season, err := strconv.Atoi(string(strings.TrimLeft(strings.TrimRight(file, "e"), l.directory+"/s")[0]))
		if err != nil {
			return nil, err
		}
		episodeBase := file[len(l.directory+"/s"+strconv.Itoa(season)+"e")-2 : len(file)]
		eNum, err := strconv.Atoi(strings.TrimRight(episodeBase, ".txt"))
		if err != nil {
			return nil, err
		}

		e := episode.New(val, season, eNum)
		episodes = append(episodes, *e)
	}
	return episodes, nil
}
