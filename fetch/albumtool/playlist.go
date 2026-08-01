package main

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os/exec"
	"strings"
)

// cmdPlaylist builds (or rebuilds) a Spotify playlist containing every
// top-100 album's tracks, in list order.
//
// One-time setup: create an app at https://developer.spotify.com/dashboard
// with redirect URI http://127.0.0.1:8123/callback, and put its client id
// in .env as SPOTIFY_CLIENT_ID. Authorisation happens in the browser via
// PKCE — no client secret is needed and no credentials touch this tool.
const spotifyRedirect = "http://127.0.0.1:8123/callback"

func cmdPlaylist(root string, args []string) error {
	clientID := envValue(root, "SPOTIFY_CLIENT_ID")
	if clientID == "" {
		return fmt.Errorf("SPOTIFY_CLIENT_ID is not set (in the environment or .env);\ncreate an app at https://developer.spotify.com/dashboard with redirect URI %s", spotifyRedirect)
	}

	fs := flag.NewFlagSet("playlist", flag.ExitOnError)
	into := fs.String("into", "", "add tracks to an existing playlist (id or URL) instead of creating one")
	fs.Parse(args)
	name := "Bill's Top 100 Albums"
	if fs.NArg() > 0 {
		name = strings.Join(fs.Args(), " ")
	}

	slugs, err := readSlugList(topAlbumsPath(root))
	if err != nil {
		return err
	}

	token, err := spotifyAuthorize(clientID)
	if err != nil {
		return err
	}

	me, err := spotifyGet(token, "https://api.spotify.com/v1/me")
	if err != nil {
		return err
	}
	fmt.Printf("Authorised as %s\n", getString(me, "id"))

	var playlist map[string]any
	var playlistID string
	if *into != "" {
		playlistID = *into
		if i := strings.LastIndex(playlistID, "/playlist/"); i >= 0 {
			playlistID = playlistID[i+len("/playlist/"):]
			playlistID, _, _ = strings.Cut(playlistID, "?")
		}
		playlist, err = spotifyGet(token, "https://api.spotify.com/v1/playlists/"+playlistID)
		if err != nil {
			return err
		}
		fmt.Printf("Adding to existing playlist %q\n", getString(playlist, "name"))
	} else {
		playlist, err = spotifyPost(token, "https://api.spotify.com/v1/me/playlists", map[string]any{
			"name":        name,
			"description": "The albums from logicalcobwebs.com/bill/albums, in release-date order.",
			"public":      true,
		})
		if err != nil {
			return err
		}
		playlistID = getString(playlist, "id")
	}

	var uris []string
	missing := 0
	for _, slug := range slugs {
		data, err := readAlbumData(root, slug)
		if err != nil {
			return err
		}
		spotify := getString(data, "spotify")
		if spotify == "" {
			// MusicBrainz had no Spotify relationship; search Spotify
			// itself and save the result for the Listen button too
			spotify = searchSpotifyAlbum(token, getString(data, "title"), getString(data, "artist"))
			if spotify == "" {
				fmt.Printf("No Spotify link for %s — skipped\n", slug)
				missing++
				continue
			}
			data["spotify"] = spotify
			if err := writeAlbumData(root, slug, data); err != nil {
				return err
			}
			fmt.Printf("%s: found via Spotify search -> %s\n", slug, spotify)
		}
		albumID := strings.TrimPrefix(spotify, "https://open.spotify.com/album/")
		tracks, err := spotifyGet(token, "https://api.spotify.com/v1/albums/"+albumID+"/tracks?limit=50")
		if err != nil {
			fmt.Printf("Tracks for %s: %v — skipped\n", slug, err)
			missing++
			continue
		}
		items, _ := tracks["items"].([]any)
		for _, item := range items {
			if m, ok := item.(map[string]any); ok {
				uris = append(uris, getString(m, "uri"))
			}
		}
		fmt.Printf("%s: %d tracks\n", slug, len(items))
	}

	// Only real track URIs; anything else 403s the whole batch
	valid := uris[:0]
	for _, u := range uris {
		if strings.HasPrefix(u, "spotify:track:") {
			valid = append(valid, u)
		} else {
			fmt.Printf("Skipping non-track URI %q\n", u)
		}
	}
	uris = valid

	if err := addTracks(token, playlistID, uris); err != nil {
		return err
	}

	fmt.Printf("\nCreated %q: %d tracks from %d albums (%d skipped)\n%s\n",
		name, len(uris), len(slugs)-missing, missing, getString(map[string]any{"u": externalURL(playlist)}, "u"))
	return nil
}

// addTracks adds URIs in batches, bisecting on failure so one bad
// entry (or a mid-run error) is pinpointed instead of failing the lot.
func addTracks(token, playlistID string, uris []string) error {
	if len(uris) == 0 {
		return nil
	}
	_, err := spotifyPost(token, "https://api.spotify.com/v1/playlists/"+playlistID+"/tracks", map[string]any{"uris": uris})
	if err == nil {
		return nil
	}
	if len(uris) > 100 {
		mid := len(uris) / 2
		if err := addTracks(token, playlistID, uris[:mid]); err != nil {
			return err
		}
		return addTracks(token, playlistID, uris[mid:])
	}
	if len(uris) == 1 {
		fmt.Printf("Could not add %s: %v\n", uris[0], err)
		return nil
	}
	mid := len(uris) / 2
	if err := addTracks(token, playlistID, uris[:mid]); err != nil {
		return err
	}
	return addTracks(token, playlistID, uris[mid:])
}

// searchSpotifyAlbum finds an album by title and artist, returning its
// open.spotify.com URL or "".
func searchSpotifyAlbum(token, title, artist string) string {
	query := url.Values{
		"q":     {fmt.Sprintf("album:%s artist:%s", title, artist)},
		"type":  {"album"},
		"limit": {"1"},
	}
	result, err := spotifyGet(token, "https://api.spotify.com/v1/search?"+query.Encode())
	if err != nil {
		return ""
	}
	albums, _ := result["albums"].(map[string]any)
	items, _ := albums["items"].([]any)
	if len(items) == 0 {
		return ""
	}
	first, _ := items[0].(map[string]any)
	return externalURL(first)
}

func externalURL(obj map[string]any) string {
	if ext, ok := obj["external_urls"].(map[string]any); ok {
		return getString(ext, "spotify")
	}
	return ""
}

// spotifyAuthorize runs the PKCE authorisation-code flow: open the
// browser, catch the redirect on localhost, swap the code for a token.
func spotifyAuthorize(clientID string) (string, error) {
	verifierBytes := make([]byte, 32)
	if _, err := rand.Read(verifierBytes); err != nil {
		return "", err
	}
	verifier := base64.RawURLEncoding.EncodeToString(verifierBytes)
	sum := sha256.Sum256([]byte(verifier))
	challenge := base64.RawURLEncoding.EncodeToString(sum[:])

	authURL := "https://accounts.spotify.com/authorize?" + url.Values{
		"client_id":             {clientID},
		"response_type":         {"code"},
		"redirect_uri":          {spotifyRedirect},
		"scope":                 {"playlist-modify-public playlist-modify-private"},
		"code_challenge_method": {"S256"},
		"code_challenge":        {challenge},
	}.Encode()

	codeCh := make(chan string, 1)
	server := &http.Server{Addr: "127.0.0.1:8123"}
	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Authorised — you can close this tab and return to the terminal.")
		codeCh <- r.URL.Query().Get("code")
	})
	go server.ListenAndServe()
	defer server.Close()

	fmt.Printf("Opening browser to authorise (or visit yourself):\n%s\n", authURL)
	exec.Command("open", authURL).Start()
	code := <-codeCh
	if code == "" {
		return "", fmt.Errorf("authorisation was denied")
	}

	resp, err := http.PostForm("https://accounts.spotify.com/api/token", url.Values{
		"grant_type":    {"authorization_code"},
		"code":          {code},
		"redirect_uri":  {spotifyRedirect},
		"client_id":     {clientID},
		"code_verifier": {verifier},
	})
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("token exchange: %s: %s", resp.Status, body)
	}
	var tokenData struct {
		AccessToken string `json:"access_token"`
		Scope       string `json:"scope"`
	}
	if err := json.Unmarshal(body, &tokenData); err != nil {
		return "", err
	}
	fmt.Printf("Granted scopes: %s\n", tokenData.Scope)
	return tokenData.AccessToken, nil
}

func spotifyGet(token, apiURL string) (map[string]any, error) {
	return spotifyRequest(token, "GET", apiURL, nil)
}

func spotifyPost(token, apiURL string, payload map[string]any) (map[string]any, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	return spotifyRequest(token, "POST", apiURL, body)
}

func spotifyRequest(token, method, apiURL string, body []byte) (map[string]any, error) {
	req, err := http.NewRequest(method, apiURL, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+token)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= 300 {
		return nil, fmt.Errorf("%s %s: %s: %s", method, apiURL, resp.Status, respBody)
	}
	return decodeJSON(bytes.NewReader(respBody))
}
