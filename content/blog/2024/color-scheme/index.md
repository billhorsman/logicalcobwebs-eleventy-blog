---
title: Light and dark
description: Adding light and dark color controls to this site
excerpt: This site has had a light and dark mode for a while but now you can toggle it on the site.
ogImage: content/blog/2024/color-scheme/og.png
date: 2024-12-11
author: Bill Horsman
tags:
  - coding
---

Following advice in the article [Native HTML light and dark color scheme switching](https://htmhell.dev/adventcalendar/2024/9/) from the wonderful HTMLHell website, I've added light and dark color controls to this site.

I've been wanting to do this for a while and finally got around to it. It uses the [light-dark](https://developer.mozilla.org/en-US/docs/Web/CSS/color_value/light-dark) CSS function which is now widely supported. 

The controls are in the bottom right of the page. If you can't see them then you're using a browser that doesn't support the `light-dark` CSS function. The fallback is the browser's default color scheme.

## How it works

Some HTML:

```html
<html>
  <head>
    <meta name="color-scheme" content="auto">
  </head>
  <body>
    <section class="scheme-switcher" aria-label="Color scheme switcher">
      <button aria-pressed="false" value="light">
        Light
      </button>
      <button aria-pressed="true" value="light dark">
        Auto
      </button>
      <button aria-pressed="false" value="dark">
        Dark
      </button>
    </section>			
  </body>
</html>
```

A little CSS (as an example):

```css
body {
  background-color: light-dark(white, black);
}
```

A some JavaScript to wire it up (simplified):

```js
button.addEventListener("click", () => {
  colorScheme.content = button.value;
  button.setAttribute("aria-pressed", true);
  localStorage.setItem("color-scheme", button.value);
});
```

As a bonus, make the controls only show up if the browser supports the `light-dark` CSS function:

```css
.scheme-switcher {
  display: none;
}

@supports (color: light-dark(black, white)) {
  .scheme-switcher {
    display: flex;
  }
}
```

## Conclusion

Some things I like about this implementation:

- No flash of incorrect color scheme on page load ([FOUC](https://en.wikipedia.org/wiki/Flash_of_unstyled_content))
- Responds to the system changing the color scheme without page reload
- Graceful fallback to the browser's default color scheme

Supported on most modern browsers and doesn't screw things up for those that don't.

### Update 4 May 2025

By making use of CSS properties and the `@supports` syntax, we can more gracefully fallback to, say, light mode if `light-dark` is not supported.

```css
--background-color: #ffe;
--color: #333;

@supports (color: light-dark(black, white)) {
  --background-color: light-dark(#ffe, #220);
  --color: light-dark(#333, #ddd);
}
```

My proposition is that literal colour values in CSS are only used as properties and these are all defined like the ones above. Try and limit the range of colours you are using. 

That's a little bit nicer than just using the browser default colours if `light-dark` isn't supported.
