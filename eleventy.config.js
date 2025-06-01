import { IdAttributePlugin, InputPathToUrlTransformPlugin, HtmlBasePlugin } from "@11ty/eleventy";
import { feedPlugin } from "@11ty/eleventy-plugin-rss";
import pluginSyntaxHighlight from "@11ty/eleventy-plugin-syntaxhighlight";
import pluginNavigation from "@11ty/eleventy-navigation";
import Image from "@11ty/eleventy-img";
import { eleventyImageTransformPlugin } from "@11ty/eleventy-img";
import { toSentence } from "./lib/toSentence.mjs";

import pluginFilters from "./_config/filters.js";

/** @param {import("@11ty/eleventy").UserConfig} eleventyConfig */
export default async function(eleventyConfig) {
	// Drafts, see also _data/eleventyDataSchema.js
	eleventyConfig.addPreprocessor("drafts", "*", (data, content) => {
		if(data.draft && process.env.ELEVENTY_RUN_MODE === "build") {
			return false;
		}
	});

	// Copy the contents of the `public` folder to the output folder
	// For example, `./public/css/` ends up in `_site/css/`
	eleventyConfig
		.addPassthroughCopy({
			"./public/": "/"
		})
		.addPassthroughCopy("content/**/*.pdf", {
			mode: "html-relative"
		})
		.addPassthroughCopy("./content/feed/pretty-atom-feed.xsl");

	// Run Eleventy when these files change:
	// https://www.11ty.dev/docs/watch-serve/#add-your-own-watch-targets

	// Watch content images for the image pipeline.
	eleventyConfig.addWatchTarget("content/**/*.{svg,webp,png,jpeg}");

	// Per-page bundles, see https://github.com/11ty/eleventy-plugin-bundle
	// Adds the {% css %} paired shortcode
	eleventyConfig.addBundle("css", {
		toFileDirectory: "dist",
	});
	// Adds the {% js %} paired shortcode
	eleventyConfig.addBundle("js", {
		toFileDirectory: "dist",
	});

	// Official plugins
	eleventyConfig.addPlugin(pluginSyntaxHighlight, {
		preAttributes: { tabindex: 0 }
	});
	eleventyConfig.addPlugin(pluginNavigation);
	eleventyConfig.addPlugin(HtmlBasePlugin);
	eleventyConfig.addPlugin(InputPathToUrlTransformPlugin);

	eleventyConfig.addPlugin(feedPlugin, {
		type: "atom", // or "rss", "json"
		outputPath: "/feed/feed.xml",
		stylesheet: "pretty-atom-feed.xsl",
		collection: {
			name: "posts",
			limit: 10,
		},
		metadata: {
			language: "en",
			title: "A Blog from Logical Cobwebs",
			subtitle: "Topics including coding, cycling, running, and the Hebrides.",
			base: "https://logicalcobwebs.com/blog/",
			author: {
				name: "Bill Horsman"
			}
		}
	});

	// Image optimization: https://www.11ty.dev/docs/plugins/image/#eleventy-transform
	eleventyConfig.addPlugin(eleventyImageTransformPlugin, {
		// File extensions to process in _site folder
		extensions: "html",

		// Output formats for each image.
		formats: ["avif", "webp", "auto"],

		// widths: ["auto"],

		defaultAttributes: {
			// e.g. <img loading decoding> assigned on the HTML tag will override these values.
			loading: "lazy",
			decoding: "async",
		}
	});

	// Filters
	eleventyConfig.addPlugin(pluginFilters);

	eleventyConfig.addPlugin(IdAttributePlugin, {
		// by default we use Eleventy’s built-in `slugify` filter:
		// slugify: eleventyConfig.getFilter("slugify"),
		// selector: "h1,h2,h3,h4,h5,h6", // default
	});

	eleventyConfig.addShortcode("currentBuildDate", () => {
		return (new Date()).toISOString();
	});

	eleventyConfig.addShortcode("currentYear", function () {
		return new Date().getFullYear();
	});

	eleventyConfig.addFilter("yearsSince", (year) => {
		const thisYear = new Date().getFullYear();
		return thisYear - year;
	});

	eleventyConfig.addFilter("daysToGo", (date) => {
		const event = new Date(date);
		const now = new Date();
		const diffTime = event - now;
		const diffDays = Math.ceil(diffTime / (1000 * 60 * 60 * 24));
		return Math.max(diffDays, 0);
	});

	eleventyConfig.addFilter("filmYear", (film) => {
		return new Date(film.release_date).getFullYear();
	});

	eleventyConfig.addFilter("directors", (film) => {
		return toSentence(film.credits.crew.filter(crew => crew.job === "Director").map(crew => crew.name));
	});

	eleventyConfig.addFilter("directorsOfPhotography", (film) => {
		return toSentence(film.credits.crew.filter(crew => crew.job === "Director of Photography").map(crew => crew.name));
	});

	eleventyConfig.addFilter("cast", (film) => {
		return film.credits.cast.map(cast => cast.name).join(", ");
	});

	eleventyConfig.addFilter("topCast", (film) => {
		return toSentence(film.credits.cast.filter(cast => cast.popularity > 2).map(cast => cast.name));
	});

	eleventyConfig.addFilter("nameSentence", (list) => {
		return toSentence(list.map(o => o.name));
	});

	eleventyConfig.addFilter("toSentence", (list) => {
		return toSentence(list);
	});

	eleventyConfig.addAsyncShortcode(
		"shareImageUri",
		async function shareImageUri(src, format) {

			// Full list of formats here: https://www.11ty.dev/docs/plugins/image/#output-formats
			// Warning: Avif can be resource-intensive so take care!

			format = format || "png";
			let formats = [format];
			let metadata = await Image(src, {
				widths: ["600px"],
				formats,
				outputDir: "./_site/img/", // Advanced usage note: `eleventyConfig.dir` works here because we’re using addPlugin.
			});

			const data = metadata[format][0];
			// data.url might be /blog/hello-world/xfO_genLg4-600.jpeg
			// note the filename is a content hash-width combination
			return data.url;
		}
	);
	
	eleventyConfig.addAsyncShortcode(
		"canonicalUrlToUse",
		async function (canonicalUrl) {
			let url = canonicalUrl
			if (!url) {
				url = this.page.url;
			}	
			if (!url.startsWith("http")) {
				url = `https://logicalcobwebs.com${url}`;
			}
			return url;
		}
	);

	eleventyConfig.addAsyncShortcode(
		"summariseFilms", 
		async function(list) {
			return toSentence(list.map(slug => {
				const film = this.ctx.films[slug];
				return `<a href="/bill/films/${slug}">${film.title}</a>`;
			}));
		}
	);

	eleventyConfig.addAsyncShortcode(
		"filmLink", 
		async function(slug) {
			const film = this.ctx.films[slug];
			if (film) {
				return `<a href="/bill/films/${slug}">${film.title}</a>`;
			} 
			return `Missing film: ${slug}`;
		}
	);



	// Features to make your build faster (when you need them)

	// If your passthrough copy gets heavy and cumbersome, add this line
	// to emulate the file copy on the dev server. Learn more:
	// https://www.11ty.dev/docs/copy/#emulate-passthrough-copy-during-serve

	// eleventyConfig.setServerPassthroughCopyBehavior("passthrough");
};

export const config = {
	// Control which files Eleventy will process
	// e.g.: *.md, *.njk, *.html, *.liquid
	templateFormats: [
		"md",
		"njk",
		"html",
		"liquid",
		"11ty.js",
	],

	// Pre-process *.md files with: (default: `liquid`)
	markdownTemplateEngine: "njk",

	// Pre-process *.html files with: (default: `liquid`)
	htmlTemplateEngine: "njk",

	// These are all optional:
	dir: {
		input: "content",          // default: "."
		includes: "../_includes",  // default: "_includes" (`input` relative)
		data: "../_data",          // default: "_data" (`input` relative)
		output: "_site"
	},

	// -----------------------------------------------------------------
	// Optional items:
	// -----------------------------------------------------------------

	// If your site deploys to a subdirectory, change `pathPrefix`.
	// Read more: https://www.11ty.dev/docs/config/#deploy-to-a-subdirectory-with-a-path-prefix

	// When paired with the HTML <base> plugin https://www.11ty.dev/docs/plugins/html-base/
	// it will transform any absolute URLs in your HTML to include this
	// folder name and does **not** affect where things go in the output folder.

	// pathPrefix: "/",
};
