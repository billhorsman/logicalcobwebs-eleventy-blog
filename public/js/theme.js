const colorScheme = document.querySelector('meta[name=color-scheme]');
let currentScheme = localStorage.getItem('color-scheme') || colorScheme.content;
colorScheme.content = currentScheme;
const switchButtons = document.querySelectorAll('button');

switchButtons.forEach(button => {
  button.setAttribute(
    'aria-pressed', button.value === currentScheme
  )

  button.addEventListener('click', () => {
    const currentButton = button;

    switchButtons.forEach(
      button => button.setAttribute(
        'aria-pressed', button === currentButton
      )
    );

    colorScheme.content = button.value;
    localStorage.setItem('color-scheme', button.value);
  });
});
