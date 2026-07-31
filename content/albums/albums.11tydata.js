export default {
	// Work in progress: keep /albums out of collections (sitemap, feeds,
	// tag pages) and out of search engines until it's ready.
	eleventyExcludeFromCollections: true,
	noindex: true,
	layout: "layouts/base.njk",
};
