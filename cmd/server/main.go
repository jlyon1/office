package main

import (
	"fmt"

	"github.com/jlyon1/office/api"
	"github.com/jlyon1/office/loader"
	"github.com/jlyon1/office/search"
)

func main() {
	l := loader.New("./script")

	// args := os.Args[1:]
	// query := "Dwight you ignorant slut"
	// if len(args) > 0 {
	// 	query = args[0]
	// }
	x, err := search.Load(*l)
	if err != nil {
		fmt.Println(err)
	}
	api := api.New(*x)
	fmt.Println("Starting server")

	api.Run()

}
