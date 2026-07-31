export default {
	tags: [
		"posts"
	],
	"layout": "layouts/post.njk",
	author: "Bill Horsman",
	// Film review posts (those with a filmSlug) derive their title, share
	// image, and description from the film data and star rating; explicit
	// frontmatter still wins.
	eleventyComputed: {
		title: (data) => data.title || (data.filmSlug ? data.films[data.filmSlug]?.title : undefined),
		ogImage: (data) => data.ogImage || (data.filmSlug ? `content/bill/films/backdrops/${data.filmSlug}.jpg` : undefined),
		description: (data) => {
			if (data.description) return data.description;
			if (!data.filmSlug) return undefined;
			const stars = typeof data.stars === "number" ? ` — ${"⭐".repeat(data.stars)}` : "";
			return `Bill's review of ${data.title}${stars}`;
		},
	},
};
