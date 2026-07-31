package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"

	xdraw "golang.org/x/image/draw"
)

// The montage is the OpenGraph image for the albums index: a grid of
// covers sampled evenly across the top 100 (so it spans the eras), at
// 1200x600 — close to the OG-preferred 1.91:1.
const (
	montageCols = 6
	montageRows = 3
	montageCell = 200
)

func montagePath(root string) string {
	return coverPath(root, "../montage")
}

func cmdMontage(root string) error {
	slugs, err := readSlugList(topAlbumsPath(root))
	if err != nil {
		return err
	}

	// Sample evenly, skipping albums without covers
	var withCovers []string
	for _, slug := range slugs {
		if _, err := os.Stat(coverPath(root, slug)); err == nil {
			withCovers = append(withCovers, slug)
		}
	}
	want := montageCols * montageRows
	if len(withCovers) < want {
		return fmt.Errorf("only %d albums have covers; need %d", len(withCovers), want)
	}
	var picks []string
	for i := 0; i < want; i++ {
		picks = append(picks, withCovers[i*len(withCovers)/want])
	}

	canvas := image.NewRGBA(image.Rect(0, 0, montageCols*montageCell, montageRows*montageCell))
	for i, slug := range picks {
		f, err := os.Open(coverPath(root, slug))
		if err != nil {
			return err
		}
		cover, _, err := image.Decode(f)
		f.Close()
		if err != nil {
			return fmt.Errorf("decoding %s: %w", slug, err)
		}
		x := (i % montageCols) * montageCell
		y := (i / montageCols) * montageCell
		cell := image.Rect(x, y, x+montageCell, y+montageCell)
		xdraw.ApproxBiLinear.Scale(canvas, cell, cover, squareCrop(cover.Bounds()), xdraw.Over, nil)
	}

	out, err := os.Create(montagePath(root))
	if err != nil {
		return err
	}
	defer out.Close()
	if err := jpeg.Encode(out, canvas, &jpeg.Options{Quality: 85}); err != nil {
		return err
	}
	fmt.Printf("Wrote %s (%d covers)\n", montagePath(root), want)
	return nil
}

// squareCrop returns the largest centred square within bounds, so
// slightly non-square covers fill their cell without distortion.
func squareCrop(b image.Rectangle) image.Rectangle {
	w, h := b.Dx(), b.Dy()
	size := min(w, h)
	x := b.Min.X + (w-size)/2
	y := b.Min.Y + (h-size)/2
	return image.Rect(x, y, x+size, y+size)
}
