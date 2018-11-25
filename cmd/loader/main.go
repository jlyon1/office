package main

import (
	"fmt"
	"os"

	"github.com/jlyon1/office/loader"
	"github.com/jlyon1/office/search"
)

func main() {

	l := loader.New("./script")

	args := os.Args[1:]
	query := "Dwight you ignorant slut"
	if len(args) > 0 {
		query = args[0]
	}
	x, err := search.Load(*l)
	resp := x.Query(query)
	if err != nil {
		fmt.Println(err)
	}
	for _, r := range resp {
		fmt.Println(r.Episode.GetReadableName(), r.Rank)
		for _, line := range r.Lines {
			fmt.Println(line)
		}
	}
}
