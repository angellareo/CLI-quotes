package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"
)

type Quote struct {
	Author      string `json:"author"`
	SourceTitle string `json:"source"`
	Quote       string `json:"quote"`
	SourceURL   string `json:"link"`
}

func parseQuotes(blob []byte) []Quote {
	var quotes []Quote
	if err := json.Unmarshal(blob, &quotes); err != nil {
		log.Fatal(err)
	}

	return quotes
}

func chooseRandomQuote(quotes []Quote) Quote {
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(quotes))
	return quotes[randomIndex]
}

func main() {
	var (
		quotesFile  = []byte(quotesJSON)
		quotes      = parseQuotes(quotesFile)
		randomQuote = chooseRandomQuote(quotes)
	)
	fmt.Println(randomQuote.Quote)
	fmt.Printf("\n— %s\n\n", randomQuote.Author)
}
