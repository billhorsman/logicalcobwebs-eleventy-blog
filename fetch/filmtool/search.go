package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"strconv"
	"strings"
)

func cmdSearch(root string, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("usage: filmtool search <title>")
	}
	query := strings.Join(args, " ")

	data, err := apiGetJSON("search/movie", url.Values{"query": {query}})
	if err != nil {
		return err
	}
	results, _ := data["results"].([]any)
	if len(results) == 0 {
		return fmt.Errorf("no TMDB matches for %q", query)
	}

	// A single match needs no picking
	if len(results) == 1 {
		m, _ := results[0].(map[string]any)
		fmt.Printf("One match: %s (%s)\n", getString(m, "title"), yearOf(getString(m, "release_date")))
		return addFromResult(root, m)
	}

	shown := min(10, len(results))
	for i := 0; i < shown; i++ {
		m, _ := results[i].(map[string]any)
		if m == nil {
			continue
		}
		year := yearOf(getString(m, "release_date"))
		if year == "" {
			year = "????"
		}
		fmt.Printf("%2d. %s (%s) — %s\n", i+1, getString(m, "title"), year, summarize(getString(m, "overview")))
	}

	fmt.Printf("Pick a film [1-%d], or q to quit: ", shown)
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return scanner.Err()
	}
	choice := strings.TrimSpace(scanner.Text())
	if choice == "" || choice == "q" {
		return nil
	}
	n, err := strconv.Atoi(choice)
	if err != nil || n < 1 || n > shown {
		return fmt.Errorf("invalid choice %q", choice)
	}

	m, _ := results[n-1].(map[string]any)
	return addFromResult(root, m)
}

func addFromResult(root string, result map[string]any) error {
	id, ok := result["id"].(json.Number)
	if !ok {
		return fmt.Errorf("unexpected result format")
	}
	movieID, err := strconv.Atoi(id.String())
	if err != nil {
		return err
	}
	return addFilm(root, movieID, false)
}

func summarize(overview string) string {
	overview = strings.ReplaceAll(overview, "\n", " ")
	if len(overview) > 90 {
		return overview[:90] + "…"
	}
	return overview
}
