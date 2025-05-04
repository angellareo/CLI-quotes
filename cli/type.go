package main

// Quote  object  that  captures a quote from
// the data set.
type Quote struct {
	Author      string `json:"author"`
	AuthorURL   string `json:"authorLink,omitempty"`
	SourceTitle string `json:"source"`
	Quote       string `json:"quote"`
	SourceURL   string `json:"link"`
}
