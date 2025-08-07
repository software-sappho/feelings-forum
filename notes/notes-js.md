Gotcha! Let me break down what this JavaScript code actually **does** in plain language — no jargon:

---

### What is this JS code doing?

1. **Grab the toggle button element** from the page:

```js
const toggleBtn = document.getElementById('theme-toggle');
```

This lets the code *talk* to the button you click to switch themes.

---

2. **Grab the root HTML element**:

```js
const rootElement = document.documentElement;
```

This is the `<html>` tag — the top-level element — where the `data-theme` attribute gets set to `'light'` or `'dark'`.

---

3. **Function to apply the theme:**

```js
function setTheme(theme) {
  // Set the attribute 'data-theme' on <html> (used by CSS to switch colors)
  rootElement.setAttribute('data-theme', theme);

  // Update button's aria-pressed attribute for accessibility (true if dark)
  toggleBtn.setAttribute('aria-pressed', theme === 'dark');

  // Change button text so user knows what clicking it will do next
  toggleBtn.textContent = theme === 'dark' ? 'Switch to Light Mode' : 'Switch to Dark Mode';

  // Save the chosen theme in browser storage so it remembers your choice later
  localStorage.setItem('theme', theme);
}
```

This function basically **switches the theme**, updates the toggle button's text and accessibility info, and remembers the choice for next time.

---

4. **Listen for clicks on the toggle button:**

```js
toggleBtn.addEventListener('click', () => {
  // Check current theme, default to light if none set
  const currentTheme = rootElement.getAttribute('data-theme') || 'light';

  // Decide what the new theme should be (flip it)
  const newTheme = currentTheme === 'light' ? 'dark' : 'light';

  // Apply the new theme
  setTheme(newTheme);
});
```

When you click the button, it flips between light and dark, calling the function above.

---

5. **When the page loads, check if a theme was saved before:**

```js
const savedTheme = localStorage.getItem('theme');
if (savedTheme) {
  setTheme(savedTheme);
} else if (window.matchMedia('(prefers-color-scheme: dark)').matches) {
  setTheme('dark');
}
```

* If you’ve used the site before and picked a theme, it uses that.
* Otherwise, it checks if your system prefers dark mode, and sets dark mode automatically.
* If neither, it defaults to light mode.

---

### TL;DR

* The JS **makes your theme toggle button work**.
* It **remembers your choice** so next time you visit, it uses the same theme.
* It **starts with your system preference** if you never picked one.
* It updates the page **without reloading** by switching CSS variables via the `data-theme` attribute.

---

defer tells browser to load script after HTML parsing.

