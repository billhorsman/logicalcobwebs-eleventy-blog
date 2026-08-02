import fs from "node:fs";

// Artists with more than one album in the top 100, most first —
// the albums sibling of top_directors.
export default () => {
	const slugs = JSON.parse(fs.readFileSync("_data/top_albums.json", "utf8"));
	const counts = new Map();
	for (const slug of slugs) {
		const artist = JSON.parse(fs.readFileSync(`_data/albums/${slug}.json`, "utf8")).artist;
		counts.set(artist, (counts.get(artist) ?? 0) + 1);
	}
	return [...counts]
		.filter(([, count]) => count > 1)
		.map(([name, count]) => ({ name, count }))
		.sort((a, b) => b.count - a.count || a.name.toLowerCase().localeCompare(b.name.toLowerCase()));
};
