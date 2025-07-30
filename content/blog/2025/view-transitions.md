---
title: View transitions
description: Adding view transitions to the theme switcher
date: 2025-07-30
author: Bill Horsman
tags: coding
---

Using the advice from Kevin Powell: [Spice Up Your Site with Quick & Easy Theme Transitions](https://youtu.be/f_aqzyIDudI?si=hx7KiF4egbACavFW) I've added some nicer view transitions to the theme switcher (bottom right corner of this page).

And update the javascript to wrap the change inside a `startViewTransition`:

```js
  document.startViewTransition(() => {
    toggleScheme();
    updateHint();
  })
```

A little CSS:

```css
::view-transition-old(root) {
  animation-delay: 0.5s;
}

::view-transition-new(root) {
  animation: circle-in 0.5s;
}

@keyframes circle-in {
  from {
    clip-path: circle(0% at 50% 50%);
  }

  to {
    clip-path: circle(100% at 50% 50%);
  }
}
```

Let's use a custom property so we get the duration right:

```css
:root {
  --view-transition-duration: 0.5s;
}

::view-transition-old(root) {
  animation-delay: var(--view-transition-duration);
}

::view-transition-new(root) {
  animation: circle-in var(--view-transition-duration);
}
```

Finally, let's check the user doesn't mind some motion:

```css
@media (prefers-reduced-motion: no-preference) {
  ::view-transition-old(root) {
    animation-delay: var(--view-transition-duration);
  }

  ::view-transition-new(root) {
    animation: circle-in var(--view-transition-duration);
  }
}
```

That's it. Nice and simple.
