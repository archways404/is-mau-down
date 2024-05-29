package handlers

import (
	"fmt"
	"net/http"

	"github.com/go-rod/rod"
)

func TestEndpoint(w http.ResponseWriter, r *http.Request) {
    // URL to check
    url := "https://mau.se/"

    // Start a new browser
    browser := rod.New().MustConnect()
    defer browser.MustClose()

    // Navigate to the URL
    page := browser.MustPage(url)

    // Check if the page contains a specific content or title to confirm accessibility
    title := page.MustEval("() => document.title").String()
    
    if title != "" {
        fmt.Fprintf(w, "The website is accessible. Title: %s", title)
    } else {
        http.Error(w, "The website is not accessible", http.StatusInternalServerError)
    }
}
