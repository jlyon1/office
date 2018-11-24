package main

import (
	"fmt"

	"github.com/jlyon1/office/loader"
	"github.com/jlyon1/office/search"
)

func main() {

	l := loader.New("./script")
	x, err := search.Load(*l)
	x.Query("michael scarn")
	if err != nil {
		fmt.Println(err)
	}

}
