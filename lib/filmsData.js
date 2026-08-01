import fs from "node:fs";
import { toSentence } from "./toSentence.mjs";

// Cross-film data for the top-100 pages (top cast/directors, related
// films), computed at build time from _data/films/*.json. Ported from
// the retired filmtool generate command.

export function loadTopFilms() {
	const slugs = JSON.parse(fs.readFileSync("_data/top_films.json", "utf8"));
	return slugs.map((slug) => {
		const data = JSON.parse(fs.readFileSync(`_data/films/${slug}.json`, "utf8"));
		return { slug, data };
	});
}

const cast = (film) => film.data.credits?.cast ?? [];
const directors = (film) =>
	(film.data.credits?.crew ?? []).filter((p) => p.job === "Director");

// Tally preserving first-seen order, like Ruby's tally / Go's port
function tally(names) {
	const counts = new Map();
	for (const name of names) {
		counts.set(name, (counts.get(name) ?? 0) + 1);
	}
	return [...counts].map(([name, count]) => ({ name, count }));
}

function byCountThenName(a, b) {
	return b.count - a.count || a.name.toLowerCase().localeCompare(b.name.toLowerCase());
}

// Top-3 billed cast across all films, threshold raised until the list
// drops under 10 names
export function topCast(films) {
	const names = films.flatMap((film) => cast(film).slice(0, 3).map((p) => p.name));
	let list = tally(names).sort(byCountThenName);
	for (let i = 1; i <= 5; i++) {
		if (list.length < 10) break;
		list = list.filter((nc) => nc.count > i);
	}
	return list;
}

export function topDirectors(films) {
	const names = films.flatMap((film) => directors(film).map((p) => p.name));
	return tally(names).filter((nc) => nc.count > 1).sort(byCountThenName);
}

// For each film, sentences like "<a>Paris, Texas</a> because of Wim
// Wenders", grouping other films by the shared people
export function relatedFilms(films) {
	const filmsByName = new Map(); // name -> [film], insertion-ordered
	for (const film of films) {
		for (const person of [...cast(film), ...directors(film)]) {
			const list = filmsByName.get(person.name) ?? [];
			if (!list.includes(film)) list.push(film);
			filmsByName.set(person.name, list);
		}
	}

	const related = new Map(); // film -> [{other, names}]
	for (const [name, shared] of filmsByName) {
		if (shared.length < 2) continue;
		for (const film of shared) {
			const entries = related.get(film) ?? [];
			for (const other of shared) {
				if (other === film) continue;
				let entry = entries.find((e) => e.other === other);
				if (!entry) {
					entry = { other, names: [] };
					entries.push(entry);
				}
				if (!entry.names.includes(name)) entry.names.push(name);
			}
			related.set(film, entries);
		}
	}

	const sentences = {};
	for (const film of films) {
		const groups = [];
		for (const entry of related.get(film) ?? []) {
			const key = entry.names.join("\x00");
			let group = groups.find((g) => g.key === key);
			if (!group) {
				group = { key, names: entry.names, links: [] };
				groups.push(group);
			}
			group.links.push(`<a href="../${entry.other.slug}">${entry.other.data.title}</a>`);
		}
		sentences[film.slug] = groups.map(
			(g) => `${toSentence(g.links)} because of ${toSentence(g.names)}`,
		);
	}
	return sentences;
}
