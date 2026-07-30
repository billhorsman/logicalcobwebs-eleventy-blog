package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

const (
	apiBase   = "https://api.themoviedb.org/3/"
	imageBase = "https://image.tmdb.org/t/p/w500"
)

func apiToken() (string, error) {
	token := os.Getenv("TMDB_API_TOKEN")
	if token == "" {
		return "", fmt.Errorf("TMDB_API_TOKEN is not set; create a read access token at https://www.themoviedb.org/settings/api")
	}
	return token, nil
}

// apiGet fetches a TMDB API endpoint (e.g. "movie/603") and returns the
// response body. language=en-US is always sent, matching the Ruby scripts.
func apiGet(endpoint string, params url.Values) ([]byte, error) {
	token, err := apiToken()
	if err != nil {
		return nil, err
	}
	if params == nil {
		params = url.Values{}
	}
	params.Set("language", "en-US")

	req, err := http.NewRequest("GET", apiBase+endpoint+"?"+params.Encode(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GET %s: %s: %s", endpoint, resp.Status, body)
	}
	return body, nil
}

// apiGetJSON decodes an API response into a map, preserving number
// precision (ids stay integers, ratings keep their exact form).
func apiGetJSON(endpoint string, params url.Values) (map[string]any, error) {
	body, err := apiGet(endpoint, params)
	if err != nil {
		return nil, err
	}
	return decodeJSON(bytes.NewReader(body))
}

func decodeJSON(r io.Reader) (map[string]any, error) {
	dec := json.NewDecoder(r)
	dec.UseNumber()
	var data map[string]any
	if err := dec.Decode(&data); err != nil {
		return nil, err
	}
	return data, nil
}

// marshalPretty renders JSON the way Ruby's JSON.pretty_generate does:
// two-space indent, no HTML escaping, no trailing newline.
func marshalPretty(v any) ([]byte, error) {
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "  ")
	if err := enc.Encode(v); err != nil {
		return nil, err
	}
	return bytes.TrimRight(buf.Bytes(), "\n"), nil
}

// downloadImage saves a TMDB image to dest unless it already exists.
func downloadImage(tmdbPath, dest string) error {
	if _, err := os.Stat(dest); err == nil {
		return nil
	}
	resp, err := http.Get(imageBase + tmdbPath)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("GET %s: %s", tmdbPath, resp.Status)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Printf("Writing %d bytes to %s\n", len(body), dest)
	return os.WriteFile(dest, body, 0o644)
}
