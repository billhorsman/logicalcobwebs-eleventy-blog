// Filters the top-100 list as you type. Matches titles/artists first;
// entries matching only a track, actor, or director stay visible with a
// hint showing what matched. Progressive: without JS the full list shows.
(function () {
	const input = document.querySelector("[data-list-search]");
	if (!input) return;
	const scope = document.querySelector(input.dataset.scope);
	if (!scope) return;

	let index = null;
	async function loadIndex() {
		if (!index) {
			const entries = await (await fetch(input.dataset.endpoint)).json();
			index = new Map(entries.map((e) => [e.slug, e]));
		}
		return index;
	}
	// Warm the index as soon as the field is touched
	input.addEventListener("focus", loadIndex, { once: true });

	const fold = (s) =>
		s.toLowerCase().normalize("NFD").replace(/[̀-ͯ]/g, "");

	input.addEventListener("input", async () => {
		const q = fold(input.value.trim());
		const idx = await loadIndex();
		for (const card of scope.querySelectorAll("[data-search-slug]")) {
			const entry = idx.get(card.dataset.searchSlug);
			let hint = card.querySelector(".search-match-hint");
			const show = (withHint) => {
				card.classList.remove("search-hidden");
				if (withHint) {
					if (!hint) {
						hint = document.createElement("span");
						hint.className = "search-match-hint";
						card.appendChild(hint);
					}
					hint.textContent = "“" + withHint + "”";
				} else {
					hint?.remove();
				}
			};
			if (!q || !entry) {
				show(null);
				continue;
			}
			if (fold(entry.primary).includes(q)) {
				show(null);
				continue;
			}
			const deepHit = entry.deep.find((d) => fold(d).includes(q));
			if (deepHit) {
				show(deepHit);
			} else {
				card.classList.add("search-hidden");
				hint?.remove();
			}
		}
	});
})();
