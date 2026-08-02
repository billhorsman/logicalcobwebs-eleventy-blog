import fs from "node:fs";

// Total running time of the top-100 Spotify playlist, in hours, from
// the per-album durations albumtool caches (spotify-duration-ms).
export default () => {
	const slugs = JSON.parse(fs.readFileSync("_data/top_albums.json", "utf8"));
	let ms = 0;
	for (const slug of slugs) {
		const album = JSON.parse(fs.readFileSync(`_data/albums/${slug}.json`, "utf8"));
		ms += album["spotify-duration-ms"] ?? 0;
	}
	return Math.round(ms / 3600000);
};
