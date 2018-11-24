package main

import (
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/gocolly/colly"
)

func fileName(season int, episode int) string {
	return "s" + strconv.Itoa(season) + "e" + strconv.Itoa(episode) + ".txt"
}

func padEp(episode int) string {
	if episode < 10 {
		return "0" + strconv.Itoa(episode)
	} else {
		return strconv.Itoa(episode)
	}
}

func main() {
	c := colly.NewCollector()
	season := 3
	episode := 11
	c.OnHTML("html", func(e *colly.HTMLElement) {
		res := e.ChildText("div.scrolling-script-container")
		fmt.Println(len(res))
		if len(res) == 0 {
			episode = 1
			season++
			c.Visit("https://www.springfieldspringfield.co.uk/view_episode_scripts.php?tv-show=the-office-us&episode=s0" + strconv.Itoa(season) + "e" + padEp(episode))
			return
		}

		err := ioutil.WriteFile(fileName(season, episode), []byte(res), 0644)

		if err != nil {
			fmt.Println(err)
		}

		episode++
		c.Visit("https://www.springfieldspringfield.co.uk/view_episode_scripts.php?tv-show=the-office-us&episode=s0" + strconv.Itoa(season) + "e" + padEp(episode))

	})

	c.OnError(func(e *colly.Response, err error) {
		episode = 1
		season++
		c.Visit("https://www.springfieldspringfield.co.uk/view_episode_scripts.php?tv-show=the-office-us&episode=s0" + strconv.Itoa(season) + "e" + padEp(episode))
		return
	})

	c.OnRequest(func(r *colly.Request) {

		fmt.Println("Visiting", r.URL)
	})

	c.Visit("https://www.springfieldspringfield.co.uk/view_episode_scripts.php?tv-show=the-office-us&episode=s01e01")
}
