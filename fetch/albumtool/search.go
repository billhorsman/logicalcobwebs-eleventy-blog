package main

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"strconv"
	"strings"
)

func cmdSearch(root string, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("usage: albumtool search <title>")
	}
	query := strings.Join(args, " ")

	data, err := mbGet("release-group", url.Values{
		"query": {fmt.Sprintf("releasegroup:%q AND primarytype:album", query)},
		"limit": {"10"},
	})
	if err != nil {
		return err
	}
	results, _ := data["release-groups"].([]any)
	if len(results) == 0 {
		return fmt.Errorf("no MusicBrainz matches for %q", query)
	}

	describe := func(m map[string]any) string {
		credits, _ := m["artist-credit"].([]any)
		year := yearOf(getString(m, "first-release-date"))
		if year == "" {
			year = "????"
		}
		return fmt.Sprintf("%s — %s (%s)", getString(m, "title"), artistName(credits), year)
	}

	// A single match needs no picking
	if len(results) == 1 {
		m, _ := results[0].(map[string]any)
		fmt.Printf("One match: %s\n", describe(m))
		return addAlbum(root, getString(m, "id"), false)
	}

	for i, r := range results {
		if m, ok := r.(map[string]any); ok {
			fmt.Printf("%2d. %s\n", i+1, describe(m))
		}
	}

	fmt.Printf("Pick an album [1-%d], or q to quit: ", len(results))
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return scanner.Err()
	}
	choice := strings.TrimSpace(scanner.Text())
	if choice == "" || choice == "q" {
		return nil
	}
	n, err := strconv.Atoi(choice)
	if err != nil || n < 1 || n > len(results) {
		return fmt.Errorf("invalid choice %q", choice)
	}
	m, _ := results[n-1].(map[string]any)
	return addAlbum(root, getString(m, "id"), false)
}
