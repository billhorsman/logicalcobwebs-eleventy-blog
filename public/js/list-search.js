// Filters the top-100 list as you type. Matches titles/artists first;
// entries matching only a track, actor, or director stay visible with a
// hint showing what matched. The query is mirrored into ?q= so a search
// can be shared or bookmarked. Progressive: without JS the full list shows.
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

	async function applyFilter() {
		const q = fold(input.value.trim());
		// Match only at word starts: "order" finds New Order, not recorder
		const matcher = q
			? new RegExp("(^|[^a-z0-9])" + q.replace(/[.*+?^${}()|[\]\\]/g, "\\$&"))
			: null;
		const matches = (text) => matcher.test(fold(text));
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
			if (matches(entry.primary)) {
				show(null);
				continue;
			}
			const deepHit = entry.deep.find(matches);
			if (deepHit) {
				show(deepHit);
			} else {
				card.classList.add("search-hidden");
				hint?.remove();
			}
		}
	}

	function syncURL() {
		const url = new URL(location);
		if (input.value.trim()) {
			url.searchParams.set("q", input.value.trim());
		} else {
			url.searchParams.delete("q");
		}
		history.replaceState(null, "", url);
	}

	input.addEventListener("input", () => {
		syncURL();
		applyFilter();
	});

	// Apply a shared/bookmarked search on load
	const initial = new URLSearchParams(location.search).get("q");
	if (initial) {
		input.value = initial;
		applyFilter();
	}
})();
