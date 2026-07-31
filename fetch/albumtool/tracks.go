package main

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

// Track listings live on releases, not release groups, and editions
// differ (bonus tracks, remasters). The earliest official release is
// used as the canonical tracklist; `tracks <slug> <release>` pins a
// different edition.

// canonicalRelease returns the id of the release group's earliest
// official release (dated releases first).
func canonicalRelease(mbid string) (string, error) {
	data, err := mbGet("release", url.Values{"release-group": {mbid}, "limit": {"100"}})
	if err != nil {
		return "", err
	}
	type candidate struct {
		id, date string
		official bool
	}
	var candidates []candidate
	releases, _ := data["releases"].([]any)
	for _, r := range releases {
		m, ok := r.(map[string]any)
		if !ok {
			continue
		}
		candidates = append(candidates, candidate{
			id:       getString(m, "id"),
			date:     getString(m, "date"),
			official: getString(m, "status") == "Official",
		})
	}
	if len(candidates) == 0 {
		return "", fmt.Errorf("release group %s has no releases", mbid)
	}
	sort.SliceStable(candidates, func(i, j int) bool {
		// official releases first, then dated before undated, then oldest
		if candidates[i].official != candidates[j].official {
			return candidates[i].official
		}
		di, dj := candidates[i].date, candidates[j].date
		if (di == "") != (dj == "") {
			return di != ""
		}
		return di < dj
	})
	return candidates[0].id, nil
}

// releaseTracks fetches a release's track listing as one entry per
// disc: {"title": ..., "tracks": [{"position", "title", "length"}]}.
func releaseTracks(releaseID string) ([]any, error) {
	data, err := mbGet("release/"+releaseID, url.Values{"inc": {"recordings"}})
	if err != nil {
		return nil, err
	}
	var discs []any
	media, _ := data["media"].([]any)
	for _, m := range media {
		medium, ok := m.(map[string]any)
		if !ok {
			continue
		}
		var tracks []any
		trackList, _ := medium["tracks"].([]any)
		for _, t := range trackList {
			track, ok := t.(map[string]any)
			if !ok {
				continue
			}
			length := trackLength(track)
			tracks = append(tracks, map[string]any{
				"position": track["position"],
				"title":    getString(track, "title"),
				"length":   length,
			})
		}
		if len(tracks) > 0 {
			discs = append(discs, map[string]any{
				"title":  getString(medium, "title"),
				"tracks": tracks,
			})
		}
	}
	return discs, nil
}

// trackLength formats a track's length (milliseconds, possibly on the
// recording instead of the track) as m:ss; empty when unknown.
func trackLength(track map[string]any) string {
	ms, ok := track["length"].(json.Number)
	if !ok {
		if recording, isMap := track["recording"].(map[string]any); isMap {
			ms, ok = recording["length"].(json.Number)
		}
	}
	if !ok {
		return ""
	}
	n, err := ms.Int64()
	if err != nil || n <= 0 {
		return ""
	}
	seconds := (n + 500) / 1000
	return fmt.Sprintf("%d:%02d", seconds/60, seconds%60)
}

// albumTracks resolves a release group to its canonical tracklist.
func albumTracks(mbid string) ([]any, error) {
	releaseID, err := canonicalRelease(mbid)
	if err != nil {
		return nil, err
	}
	return releaseTracks(releaseID)
}

// cmdTracks with no arguments backfills tracklists for albums that
// lack one; with <slug> <release-mbid-or-url> it pins a specific
// edition's tracklist.
func cmdTracks(root string, args []string) error {
	if len(args) == 2 {
		return pinTracks(root, args[0], args[1])
	}
	if len(args) != 0 {
		return fmt.Errorf("usage: albumtool tracks [<slug> <release-mbid-or-url>]")
	}

	entries, err := os.ReadDir(filepath.Join(root, "_data", "albums"))
	if err != nil {
		return err
	}
	missing := 0
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".json") {
			continue
		}
		slug := strings.TrimSuffix(entry.Name(), ".json")
		data, err := readAlbumData(root, slug)
		if err != nil {
			return err
		}
		if _, ok := data["tracks"]; ok {
			continue
		}

		discs, err := albumTracks(getString(data, "id"))
		// two MusicBrainz requests were just made; stay under 1/second
		time.Sleep(2200 * time.Millisecond)
		if err != nil || len(discs) == 0 {
			fmt.Printf("No tracklist for %s (%v)\n", slug, err)
			missing++
			continue
		}
		data["tracks"] = discs
		if err := writeAlbumData(root, slug, data); err != nil {
			return err
		}
		fmt.Printf("%s: %d disc(s)\n", slug, len(discs))
	}
	if missing > 0 {
		fmt.Printf("%d album(s) without tracklists\n", missing)
	}
	return nil
}

func pinTracks(root, slug, arg string) error {
	releaseID := mbidPattern.FindString(arg)
	if releaseID == "" {
		return fmt.Errorf("%q is not a MusicBrainz release id or URL", arg)
	}
	data, err := readAlbumData(root, slug)
	if err != nil {
		return err
	}
	discs, err := releaseTracks(releaseID)
	if err != nil {
		return err
	}
	if len(discs) == 0 {
		return fmt.Errorf("release %s has no tracks", releaseID)
	}
	data["tracks"] = discs
	if err := writeAlbumData(root, slug, data); err != nil {
		return err
	}
	fmt.Printf("%s: pinned tracklist from release %s (%d disc(s))\n", slug, releaseID, len(discs))
	return nil
}

func readAlbumData(root, slug string) (map[string]any, error) {
	f, err := os.Open(albumDataPath(root, slug))
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return decodeJSON(f)
}

func writeAlbumData(root, slug string, data map[string]any) error {
	out, err := marshalPretty(data)
	if err != nil {
		return err
	}
	return os.WriteFile(albumDataPath(root, slug), out, 0o644)
}
