package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()

	//what we need to do is find the table from the html source that contains 'eventText', which is:
	//table class=omnipong > tbody > tr > th > contains:eventText (this is jQuery-esque notation)

	//text in the <th> to specify which event to listen for new slots in
	eventText := "Open Singles RR"

	//text that we're searching for, the start of the info about how many remaining entries exist
	entriesRemainingText := "Remaining slots: "

	//define what happens on HTML visitation of the URL (c.Visit, below)
	c.OnHTML("table[class=omnipong]:contains(\"" + eventText +"\")", func(e *colly.HTMLElement) {
		//e *colly.HTMLElement now refers to the Open Singles (or whichever event) <table>
		//from this, we need the <th>
		allTHText := string(e.ChildText("th"));
		
		// for debugging purposes:
		/* fmt.Println(allTHText); */

		entriesRemainingIndex := strings.Index(allTHText, entriesRemainingText) + len(entriesRemainingText);
		
		//a string
		entriesRemaining := string(allTHText[entriesRemainingIndex]);

		//strconv to int
		entriesRemainingInt, err := strconv.Atoi(entriesRemaining);
		if err != nil {
			panic(err)
		}

		if entriesRemainingInt != 0 {
			//send email. you only want to do this once, so probably want some kind of global boolean to check as well
			fmt.Println("found an opening for " + eventText + "!");
		} else {
			fmt.Println("still no openings yet :(");
		}
	})

	//visit the url
	c.Visit("https://omnipong.com/T-tourney.asp?t=102&r=3504&h=")//MBCTT Lob Palace December Open, 2023
}