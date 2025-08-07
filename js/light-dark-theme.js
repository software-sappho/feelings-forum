const toggleBtn = document.getElementById('theme-toggle');
const rootElement = document.documentElement; // <html>

function setTheme(theme) {
  rootElement.setAttribute('data-theme', theme);
  toggleBtn.setAttribute('aria-pressed', theme === 'dark');
  toggleBtn.textContent = theme === 'dark' ? 'Switch to Light Mode' : 'Switch to Dark Mode';
  localStorage.setItem('theme', theme); // remember preference
}

toggleBtn.addEventListener('click', () => {
  const currentTheme = rootElement.getAttribute('data-theme') || 'light';
  const newTheme = currentTheme === 'light' ? 'dark' : 'light';
  setTheme(newTheme);
});

// On page load, check saved theme or system preference
const savedTheme = localStorage.getItem('theme');
if (savedTheme) {
  setTheme(savedTheme);
} else if (window.matchMedia('(prefers-color-scheme: dark)').matches) {
  setTheme('dark');
}
