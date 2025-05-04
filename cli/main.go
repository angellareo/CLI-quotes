package main

import (
	"fmt"
	"math/rand"
	"strings"
)

// Selects a random quote for display
func chooseRandomQuote(quotes []Quote) Quote {
	randomIndex := rand.Intn(len(quotes))
	return quotes[randomIndex]
}

// justifies  the  quotes  to a set number of
// characters per line for better display.
func justify(text string) []string {
	var (
		maxSize    = 60
		lines      = []string{}
		words      = strings.Fields(text)
		bufferLine = ""
	)

	for _, word := range words {
		if len(bufferLine)+len(word)+1 < maxSize {
			bufferLine += word + " "
		} else {
			lines = append(lines, bufferLine)
			bufferLine = word + " "
		}
	}
	lines = append(lines, bufferLine)
	return lines
}

// Prints  the  quotes  in  a  justified  and
// beautiful way.
func printQuoteText(quote Quote) {
	for _, line := range justify(quote.Quote) {
		fmt.Print("\033[95m ┃ \033[34m")
		fmt.Println(line)
		fmt.Print("\033[0m")
	}
}

func printQuoteAuthor(quote Quote) {
	fmt.Println()
	if quote.AuthorURL != "" {
		fmt.Printf("\033[34m %s (%s)\033[0m\n", quote.Author, quote.AuthorURL)
	} else {
		fmt.Printf("\033[34m %s\033[0m\n", quote.Author)
	}
	fmt.Println()
}

func main() {
	randomQuote := chooseRandomQuote(quotes)
	printQuoteAuthor(randomQuote)
	printQuoteText(randomQuote)
	printQuoteSource(randomQuote)
	fmt.Println()
}

func printQuoteSource(quote Quote) {
	fmt.Println()
	if quote.SourceURL != "" {
		fmt.Printf("\033[34m Source: \033[0m\033[96m%s\033[0m\n", quote.SourceURL)
	}
}
