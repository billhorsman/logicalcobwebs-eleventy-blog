import QRCode from "qrcode-svg";

// Renders a DCA cinema ticket as inline SVG from `dca` frontmatter:
//
//   dca:
//     date: "2026-03-12 19:45"
//     cinema: Cinema Two
//     seat: B5
//     rating: "15"
//     price: £12.00
//
// Only date is required. The QR code points at the post's canonical URL
// and the barcode is random, seeded by the URL so builds are stable.

const SITE_URL = "https://logicalcobwebs.com";

// Matches the aspect ratio of the original scanned tickets (~2.48:1)
const WIDTH = 350;
const HEIGHT = 141;

// FNV-1a, to seed the barcode from the page URL
function hash(str) {
	let h = 0x811c9dc5;
	for (let i = 0; i < str.length; i++) {
		h ^= str.charCodeAt(i);
		h = Math.imul(h, 0x01000193);
	}
	return h >>> 0;
}

function mulberry32(seed) {
	return function () {
		seed |= 0;
		seed = (seed + 0x6d2b79f5) | 0;
		let t = Math.imul(seed ^ (seed >>> 15), 1 | seed);
		t = (t + Math.imul(t ^ (t >>> 7), 61 | t)) ^ t;
		return ((t ^ (t >>> 14)) >>> 0) / 4294967296;
	};
}

// "2026-03-12 19:45" → { day: "Thu 12 Mar 2026", time: "7:45PM" }.
// YAML may parse an unquoted date as a Date (treated as UTC), so both
// forms are accepted; the wall-clock time is used as written.
function ticketDate(value) {
	let y, m, d, hh, mm;
	if (value instanceof Date) {
		[y, m, d, hh, mm] = [
			value.getUTCFullYear(), value.getUTCMonth(), value.getUTCDate(),
			value.getUTCHours(), value.getUTCMinutes(),
		];
	} else {
		const match = String(value).match(/^(\d{4})-(\d{2})-(\d{2})[T ](\d{1,2}):(\d{2})/);
		if (!match) {
			throw new Error(`dca.date must be "YYYY-MM-DD HH:MM", got ${JSON.stringify(value)}`);
		}
		[y, m, d, hh, mm] = [+match[1], +match[2] - 1, +match[3], +match[4], +match[5]];
	}
	const weekdays = ["Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"];
	const months = ["Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"];
	const weekday = weekdays[new Date(Date.UTC(y, m, d)).getUTCDay()];
	const hour12 = hh % 12 === 0 ? 12 : hh % 12;
	return {
		day: `${weekday} ${d} ${months[m]} ${y}`,
		time: `${hour12}:${String(mm).padStart(2, "0")}${hh < 12 ? "AM" : "PM"}`,
		long: `${d} ${months[m]} ${y}`,
	};
}

function escapeXML(s) {
	return String(s)
		.replaceAll("&", "&amp;")
		.replaceAll("<", "&lt;")
		.replaceAll(">", "&gt;")
		.replaceAll('"', "&quot;");
}

function barcode(rand, x, y, width, height) {
	const bars = [];
	let at = x;
	while (at < x + width) {
		const bar = Math.ceil(rand() * 3);
		const gap = Math.ceil(rand() * 3);
		bars.push(`<rect x="${at}" y="${y}" width="${bar}" height="${height}"/>`);
		at += bar + gap;
	}
	return bars.join("");
}

function qrCode(content, x, y, size) {
	const qr = new QRCode({ content, padding: 0, ecl: "M" });
	const modules = qr.qrcode.modules;
	const cell = size / modules.length;
	const rects = [];
	for (let row = 0; row < modules.length; row++) {
		for (let col = 0; col < modules.length; col++) {
			if (modules[col][row]) {
				const cx = (x + col * cell).toFixed(2);
				const cy = (y + row * cell).toFixed(2);
				rects.push(`<rect x="${cx}" y="${cy}" width="${cell.toFixed(2)}" height="${cell.toFixed(2)}"/>`);
			}
		}
	}
	return rects.join("");
}

export default function (eleventyConfig) {
	eleventyConfig.addShortcode("dcaTicket", function () {
		const dca = this.ctx.dca;
		if (!dca || !dca.date) {
			throw new Error(`dcaTicket needs dca.date frontmatter (page: ${this.page.inputPath})`);
		}
		const title = this.ctx.title;
		const seat = dca.seat || "A1";
		const price = dca.price || "£10.00";
		const cinema = dca.cinema || "Cinema One";
		const when = ticketDate(dca.date);
		const canonicalUrl = `${SITE_URL}${this.page.url}`;

		// The film's TMDB id doubles as the ticket's reference number and
		// the barcode seed; without a filmSlug, fall back to the URL.
		const film = this.ctx.filmSlug && this.ctx.films?.[this.ctx.filmSlug];
		const reference = film?.id ? String(film.id) : "";
		const rand = mulberry32(hash(reference || this.page.url));

		const label = `Ticket stub for seat ${seat} in the DCA's ${cinema.toLowerCase()} on ${when.long}`;

		const rating = dca.rating
			? `<text x="60" y="64" font-size="10.5">${escapeXML(dca.rating)}</text>`
			: "";
		const referenceText = reference
			? `<text transform="translate(243 118) rotate(-90)" font-size="7.5" letter-spacing="1">${escapeXML(reference)}</text>`
			: "";

		return `<svg viewBox="0 0 ${WIDTH} ${HEIGHT}" xmlns="http://www.w3.org/2000/svg" role="img" aria-label="${escapeXML(label)}">
  <rect width="${WIDTH}" height="${HEIGHT}" rx="5" fill="white"/>
  <g fill="black" font-family="Arial, Helvetica, sans-serif">
    <text x="60" y="30" font-size="13.5" font-weight="bold">${escapeXML(title)}</text>
    <text x="60" y="48" font-size="12">${when.day} ${when.time}</text>
    ${rating}
    <text x="60" y="90" font-size="11">${escapeXML(cinema)}&#160;&#160;&#160;Seat:&#160;&#160;${escapeXML(seat)}</text>
    <text x="60" y="106" font-size="11">Standard - ${escapeXML(price)}</text>
    ${barcode(rand, 60, 116, 130, 15)}
    ${referenceText}
    ${qrCode(canonicalUrl, 252, 26, 84)}
  </g>
</svg>`;
	});
}
