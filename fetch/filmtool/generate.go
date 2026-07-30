package main

// Port of generate.rb. The markdown output is byte-for-byte identical to
// the Ruby version (including its slightly uneven indentation, which came
// from interpolating heredocs into heredocs), so regenerating the top-100
// pages produces no spurious diffs.

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func cmdGenerate(root string) error {
	slugs, err := readTopFilmSlugs(root)
	if err != nil {
		return err
	}
	films := make([]*Film, 0, len(slugs))
	for _, slug := range slugs {
		film, err := loadFilm(root, slug)
		if err != nil {
			return err
		}
		films = append(films, film)
	}

	if err := writeTopCast(root, films); err != nil {
		return err
	}
	if err := writeTopDirectors(root, films); err != nil {
		return err
	}

	related := buildRelated(films)

	for i, film := range films {
		var prev, next *Film
		if i > 0 {
			prev = films[i-1]
		}
		if i < len(films)-1 {
			next = films[i+1]
		}
		md := renderFilmPage(film, prev, next, i, len(films), related[film])
		fmt.Printf("Writing %s.md\n", film.Slug)
		path := filepath.Join(root, "content", "bill", "films", film.Slug+".md")
		if err := os.WriteFile(path, []byte(md), 0o644); err != nil {
			return err
		}
	}
	return nil
}

type nameCount struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

// tally counts occurrences preserving first-seen order, like Ruby's tally.
func tally(names []string) []nameCount {
	counts := map[string]int{}
	var order []string
	for _, name := range names {
		if counts[name] == 0 {
			order = append(order, name)
		}
		counts[name]++
	}
	out := make([]nameCount, len(order))
	for i, name := range order {
		out[i] = nameCount{name, counts[name]}
	}
	return out
}

func sortByCountThenName(list []nameCount) {
	sort.SliceStable(list, func(i, j int) bool {
		if list[i].Count != list[j].Count {
			return list[i].Count > list[j].Count
		}
		return strings.ToLower(list[i].Name) < strings.ToLower(list[j].Name)
	})
}

func writeCountsJSON(root, filename string, list []nameCount) error {
	out, err := marshalPretty(list)
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(root, "_data", filename), out, 0o644)
}

// writeTopCast tallies the top-3 billed cast across all films, then keeps
// raising the appearance threshold until fewer than ~10 names remain.
func writeTopCast(root string, films []*Film) error {
	var names []string
	for _, film := range films {
		for _, p := range film.Cast[:min(3, len(film.Cast))] {
			names = append(names, p.Name)
		}
	}
	topCast := tally(names)
	sortByCountThenName(topCast)
	for i := 1; i <= 5; i++ {
		if len(topCast) < 10 {
			break
		}
		var kept []nameCount
		for _, nc := range topCast {
			if nc.Count > i {
				kept = append(kept, nc)
			}
		}
		topCast = kept
	}
	return writeCountsJSON(root, "top_cast.json", topCast)
}

func writeTopDirectors(root string, films []*Film) error {
	var names []string
	for _, film := range films {
		for _, d := range film.Directors() {
			names = append(names, d.Name)
		}
	}
	var topDirectors []nameCount
	for _, nc := range tally(names) {
		if nc.Count > 1 {
			topDirectors = append(topDirectors, nc)
		}
	}
	sortByCountThenName(topDirectors)
	return writeCountsJSON(root, "top_directors.json", topDirectors)
}

// relEntry links a film to another film via the people they share.
type relEntry struct {
	other *Film
	names []string
}

// buildRelated maps each film to the other films it shares cast members or
// directors with. Iteration order mirrors the Ruby version's hash insertion
// order so grouping and output order match exactly.
func buildRelated(films []*Film) map[*Film][]*relEntry {
	var nameOrder []string
	filmsByName := map[string][]*Film{}
	for _, film := range films {
		credits := append(append([]Person(nil), film.Cast...), film.Directors()...)
		for _, p := range credits {
			existing, seen := filmsByName[p.Name]
			if !seen {
				nameOrder = append(nameOrder, p.Name)
			}
			if !containsFilm(existing, film) {
				filmsByName[p.Name] = append(existing, film)
			}
		}
	}

	related := map[*Film][]*relEntry{}
	for _, name := range nameOrder {
		shared := filmsByName[name]
		if len(shared) < 2 {
			continue
		}
		for _, film := range shared {
			for _, other := range shared {
				if other == film {
					continue
				}
				entry := findEntry(related[film], other)
				if entry == nil {
					entry = &relEntry{other: other}
					related[film] = append(related[film], entry)
				}
				if !containsString(entry.names, name) {
					entry.names = append(entry.names, name)
				}
			}
		}
	}
	return related
}

func containsFilm(films []*Film, f *Film) bool {
	for _, x := range films {
		if x == f {
			return true
		}
	}
	return false
}

func containsString(list []string, s string) bool {
	for _, x := range list {
		if x == s {
			return true
		}
	}
	return false
}

func findEntry(entries []*relEntry, other *Film) *relEntry {
	for _, e := range entries {
		if e.other == other {
			return e
		}
	}
	return nil
}

func toSentence(items []string) string {
	var out strings.Builder
	for i, x := range items {
		out.WriteString(x)
		if i < len(items)-2 {
			out.WriteString(", ")
		} else if i == len(items)-2 {
			out.WriteString(" and ")
		}
	}
	return out.String()
}

// relatedSentences groups related films that share the same set of names,
// producing lines like "A and B because of X".
func relatedSentences(entries []*relEntry) []string {
	type group struct {
		names []string
		links []string
	}
	var groups []*group
	for _, e := range entries {
		key := strings.Join(e.names, "\x00")
		var g *group
		for _, existing := range groups {
			if strings.Join(existing.names, "\x00") == key {
				g = existing
				break
			}
		}
		if g == nil {
			g = &group{names: e.names}
			groups = append(groups, g)
		}
		g.links = append(g.links, fmt.Sprintf("<a href=\"../%s\">%s</a>", e.other.Slug, e.other.Title))
	}
	var sentences []string
	for _, g := range groups {
		sentences = append(sentences, toSentence(g.links)+" because of "+toSentence(g.names))
	}
	return sentences
}

func castCard(p Person) string {
	imageTag := "<div class=\"cast-card-no-image\"><i class=\"fa-solid fa-user\"></i></div>"
	if p.ProfilePath != "" {
		imageTag = fmt.Sprintf("<img src=\"../films/profiles/%s.jpg\" alt=\"%s\" loading=\"lazy\" eleventy:ignore>", p.ID, p.Name)
	}
	return "<div class=\"cast-card\">\n" +
		"  " + imageTag + "\n" +
		"  <div class=\"cast-card-info\">\n" +
		"    <span class=\"cast-card-name\">" + p.Name + "</span>\n" +
		"    <span class=\"cast-card-character\">" + p.Character + "</span>\n" +
		"  </div>\n" +
		"</div>"
}

func castGridSection(film *Film) string {
	cards := make([]string, 0, 12)
	for _, p := range film.Cast[:min(12, len(film.Cast))] {
		cards = append(cards, castCard(p))
	}
	return "<section class=\"cast-grid\">\n" +
		"  <div class=\"cast-grid-cards\">\n" +
		"    " + strings.Join(cards, "\n    ") + "\n" +
		"  </div>\n" +
		"</section>\n"
}

func relatedSection(sentences []string) string {
	if len(sentences) == 0 {
		return ""
	}
	items := make([]string, len(sentences))
	for i, s := range sentences {
		items[i] = "<li>" + s + "</li>"
	}
	return "<section class=\"related-films\">\n" +
		"  <h2>Related films</h2>\n" +
		"  <ul>\n" +
		"    " + strings.Join(items, "\n") + "\n" +
		"  </ul>\n" +
		"</section>\n"
}

func renderFilmPage(film, prev, next *Film, index, total int, related []*relEntry) string {
	prevLink := "<span><i class=\"fa-solid fa-chevron-left fa-xs\"></i> Previous</span>"
	prevHint := "Start of list"
	if prev != nil {
		prevLink = fmt.Sprintf("<a href=\"../%s\"><i class=\"fa-solid fa-chevron-left fa-xs\"></i> Previous</a>", prev.Slug)
		prevHint = prev.Title
	}
	nextLink := "<span>Next <i class=\"fa-solid fa-chevron-right fa-xs\"></i></span>"
	nextHint := "End of list"
	if next != nil {
		nextLink = fmt.Sprintf("<a href=\"../%s\">Next <i class=\"fa-solid fa-chevron-right fa-xs\"></i></a>", next.Slug)
		nextHint = next.Title
	}

	alsoKnown := ""
	if film.OriginalTitle != "" {
		alsoKnown = "Also known as " + film.OriginalTitle + "."
	}

	var b strings.Builder
	b.WriteString("---\n")
	b.WriteString("title: \"" + film.Title + "\"\n")
	b.WriteString("layout: layouts/films.njk\n")
	b.WriteString("slug: " + film.Slug + "\n")
	b.WriteString("ogImage: content/bill/films/backdrops/" + film.Slug + ".jpg\n")
	b.WriteString("description: \"" + strings.ReplaceAll(film.Overview, "\"", "\\\"") + "\"\n")
	b.WriteString("---\n")
	b.WriteString("\n")
	b.WriteString("{% set film = films[slug] %}\n")
	b.WriteString("\n")
	b.WriteString("<nav class=\"films\">\n")
	b.WriteString("  <div class=\"prev\">\n")
	b.WriteString("    " + prevLink + "\n")
	b.WriteString("  </div>\n")
	b.WriteString("  <div>\n")
	b.WriteString(fmt.Sprintf("    <a class=\"simple\" href=\"../\">%d / %d</a>\n", index+1, total))
	b.WriteString("  </div>\n")
	b.WriteString("  <div class=\"next\">\n")
	b.WriteString("    " + nextLink + "\n")
	b.WriteString("  </div>\n")
	b.WriteString("  <div class=\"hint\">\n")
	b.WriteString("    <span class=\"prev-hint\">\n")
	b.WriteString("      <span class=\"sr-only\">Previous film:</span>\n")
	b.WriteString("      " + prevHint + "\n")
	b.WriteString("    </span>\n")
	b.WriteString("    <span class=\"next-hint\">\n")
	b.WriteString("      <span class=\"sr-only\">Next film:</span>\n")
	b.WriteString("      " + nextHint + "\n")
	b.WriteString("    </span>\n")
	b.WriteString("  </div>\n")
	b.WriteString("</nav>\n")
	b.WriteString("\n")
	b.WriteString("<article class=\"film slug-" + film.Slug + "\">\n")
	b.WriteString("  <div class=\"backdrop-and-poster\">\n")
	b.WriteString("    <img class=\"poster\" src=\"../films/posters/{{ slug }}.jpg\" alt=\"\" eleventy:ignore>\n")
	b.WriteString("    <img class=\"backdrop\" src=\"../films/backdrops/{{ slug }}.jpg\" alt=\"\" eleventy:ignore>\n")
	b.WriteString("  </div>\n")
	b.WriteString("\n")
	b.WriteString("  <h1>{{ film.title }} ({{ film | filmYear }})</h1>\n")
	b.WriteString("\n")
	b.WriteString("  <p>\n")
	b.WriteString("    {%- if film.language -%}Language: {{ film.language }}.{% endif %}\n")
	b.WriteString("    " + alsoKnown + "\n")
	b.WriteString("  </p>\n")
	b.WriteString("\n")
	b.WriteString("  <p class=\"director\">\n")
	b.WriteString("    Directed by <strong>{{ film | directors }}</strong>\n")
	b.WriteString("  </p>\n")
	b.WriteString("\n")
	b.WriteString("  {%- if films.reviews[slug] -%}\n")
	b.WriteString("    <blockquote>\n")
	b.WriteString("      {{ films.reviews[slug] | safe }} <em>—&nbsp;<a href=\"/bill\">Bill</a></em>\n")
	b.WriteString("    </blockquote>\n")
	b.WriteString("  {%- endif -%}\n")
	b.WriteString("\n")
	b.WriteString("  " + castGridSection(film) + "\n")
	b.WriteString("  <section class=\"film-detail\">\n")
	b.WriteString("    <div>\n")
	b.WriteString("      <details>\n")
	b.WriteString("        <summary>\n")
	b.WriteString("          <i class=\"fa-solid fa-masks-theater\"></i>\n")
	b.WriteString("          Cast\n")
	b.WriteString("        </summary>\n")
	b.WriteString("        <ul>\n")
	b.WriteString("          {%- for cast in film.credits.cast -%}\n")
	b.WriteString("            <li>\n")
	b.WriteString("              {{ cast.name }} as <em>{{ cast.character }}</em>\n")
	b.WriteString("            </li>\n")
	b.WriteString("          {%- endfor -%}\n")
	b.WriteString("        </ul>\n")
	b.WriteString("      </details>\n")
	b.WriteString("      <details>\n")
	b.WriteString("        <summary>\n")
	b.WriteString("          <i class=\"fa-solid fa-clapperboard\"></i>\n")
	b.WriteString("          Crew\n")
	b.WriteString("        </summary>\n")
	b.WriteString("        <ul>\n")
	b.WriteString("          {%- for crew in film.credits.crew -%}\n")
	b.WriteString("            <li>\n")
	b.WriteString("              {{ crew.name }} &mdash; <em>{{ crew.job }}</em>\n")
	b.WriteString("            </li>\n")
	b.WriteString("          {%- endfor -%}\n")
	b.WriteString("        </ul>\n")
	b.WriteString("      </details>\n")
	b.WriteString("    </div>\n")
	b.WriteString("  </section>\n")
	b.WriteString("\n")
	b.WriteString("  " + relatedSection(relatedSentences(related)) + "\n")
	b.WriteString("</article>\n")
	return b.String()
}
