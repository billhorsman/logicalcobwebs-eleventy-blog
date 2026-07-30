package main

import (
	"encoding/json"
	"os"
	"sort"
	"strings"
)

type Person struct {
	ID          json.Number
	Name        string
	Character   string
	ProfilePath string
	Job         string
	Order       int64
}

type Film struct {
	Slug          string
	Title         string
	OriginalTitle string // empty unless it differs from Title
	Year          string
	Overview      string
	Cast          []Person
	Crew          []Person
}

func getString(m map[string]any, key string) string {
	s, _ := m[key].(string)
	return s
}

func parsePerson(m map[string]any) Person {
	p := Person{
		Name:        getString(m, "name"),
		Character:   getString(m, "character"),
		ProfilePath: getString(m, "profile_path"),
		Job:         getString(m, "job"),
		Order:       999,
	}
	if id, ok := m["id"].(json.Number); ok {
		p.ID = id
	}
	if order, ok := m["order"].(json.Number); ok {
		if n, err := order.Int64(); err == nil {
			p.Order = n
		}
	}
	return p
}

func parsePeople(v any) []Person {
	list, _ := v.([]any)
	people := make([]Person, 0, len(list))
	for _, item := range list {
		if m, ok := item.(map[string]any); ok {
			people = append(people, parsePerson(m))
		}
	}
	return people
}

func loadFilm(root, slug string) (*Film, error) {
	f, err := os.Open(filmDataPath(root, slug))
	if err != nil {
		return nil, err
	}
	defer f.Close()
	data, err := decodeJSON(f)
	if err != nil {
		return nil, err
	}

	film := &Film{
		Slug:     slug,
		Title:    getString(data, "title"),
		Year:     getString(data, "year"),
		Overview: getString(data, "overview"),
	}
	if credits, ok := data["credits"].(map[string]any); ok {
		film.Cast = parsePeople(credits["cast"])
		film.Crew = parsePeople(credits["crew"])
	}
	// original_title is only interesting when it differs from the title
	if orig := getString(data, "original_title"); orig != "" &&
		strings.ToLower(orig) != strings.ToLower(film.Title) {
		film.OriginalTitle = orig
	}
	return film, nil
}

func (f *Film) Directors() []Person {
	var directors []Person
	for _, p := range f.Crew {
		if p.Job == "Director" {
			directors = append(directors, p)
		}
	}
	return directors
}

// TopCast returns the first 12 billed cast members, in credit order.
func (f *Film) TopCast() []Person {
	cast := append([]Person(nil), f.Cast...)
	sort.SliceStable(cast, func(i, j int) bool { return cast[i].Order < cast[j].Order })
	return cast[:min(12, len(cast))]
}
