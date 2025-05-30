document.querySelectorAll('button[data-toggle-list]').forEach(button => {
  const section = button.closest('section');
  const view = button.dataset.toggleList;

  if (localStorage.getItem('toggle-list') === view) {
    section.classList.add(view);
  }

  button.addEventListener('click', () => {
    if (section.classList.contains(button.dataset.toggleList)) {
      section.classList.remove(button.dataset.toggleList);
      localStorage.removeItem('toggle-list');
    } else {
      section.classList.add(button.dataset.toggleList);
      localStorage.setItem('toggle-list', view);
    }
  });
});
