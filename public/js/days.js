document.querySelectorAll("time[days-to-go]").forEach((time) => {
	const string = time.getAttribute("date");
	const event = new Date(string);
	const now = new Date();
	const diffTime = event - now;
	const diffDays = Math.ceil(diffTime / (1000 * 60 * 60 * 24));
	const weeks = Math.floor(diffDays / 7);
	const days = diffDays - weeks * 7;
	const parts = [];
	if (weeks > 1) {
		parts.push(`${weeks} weeks`);
	} else if (weeks == 1) {
		parts.push(`1 week`);
	}
	if (days > 1) {
		parts.push(`${days} days`);
	} else if (days == 1) {
		parts.push(`1 day`);
	}
	if (diffDays > 0) {
		time.innerText = `${parts.join(" and ")} to go`;
	} else if (diffDays == 0) {
		time.innerText = "today";
	}
});
