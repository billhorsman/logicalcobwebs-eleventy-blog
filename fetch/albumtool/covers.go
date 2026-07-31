package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// cmdCovers sweeps every album whose cover file is missing and retries
// the full cover chain (Cover Art Archive, then Discogs).
func cmdCovers(root string) error {
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
		if _, err := os.Stat(coverPath(root, slug)); err == nil {
			continue
		}

		f, err := os.Open(albumDataPath(root, slug))
		if err != nil {
			return err
		}
		data, err := decodeJSON(f)
		f.Close()
		if err != nil {
			return err
		}
		title, artist := getString(data, "title"), getString(data, "artist")
		fmt.Printf("Missing cover: %s — %s\n", title, artist)
		if err := downloadCover(root, getString(data, "id"), title, artist, coverPath(root, slug)); err != nil {
			return err
		}
		if _, err := os.Stat(coverPath(root, slug)); err != nil {
			missing++
		}
	}
	if missing > 0 {
		fmt.Printf("%d album(s) still without covers\n", missing)
	} else {
		fmt.Println("All albums have covers")
	}
	return nil
}
