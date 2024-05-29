package handlers

import (
	"fmt"
	"net/http"
)

func checkURLAccessibility(url string) error {
    // Perform an HTTP GET request
    resp, err := http.Get(url)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    // Check if the status code is 200 OK
    if resp.StatusCode != http.StatusOK {
        return fmt.Errorf("received non-OK HTTP status: %s", resp.Status)
    }

    return nil
}

func CheckURLEndpoint(w http.ResponseWriter, r *http.Request) {
    // URL to check
    url := "https://mau.se/"

    // Check URL accessibility
    err := checkURLAccessibility(url)
    if err != nil {
        http.Error(w, fmt.Sprintf("The website is not accessible: %v", err), http.StatusInternalServerError)
        return
    }

    fmt.Fprint(w, "The website is accessible.")
}
