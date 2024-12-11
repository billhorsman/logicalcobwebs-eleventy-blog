---
title: Light and dark
description: Adding light and dark color controls to this site
excerpt: This site has had a light and dark mode for a while but now you can toggle it on the site.
canonical: https://logicalcobwebs.com/blog/2024/color-scheme
ogImage: content/blog/2024/color-scheme/og.png
date: 2024-12-11
author: Bill Horsman
tags:
  - coding
---

Following advice in the article [Native HTML light and dark color scheme switching](https://htmhell.dev/adventcalendar/2024/9/) from the wonderful HTMLHell website, I've added light and dark color controls to this site.

I've been wanting to do this for a while and finally got around to it. It uses the [light-dark](https://developer.mozilla.org/en-US/docs/Web/CSS/color_value/light-dark) CSS function which is now widely supported. 

The controls are in the bottom right of the page. If you can't see them then you're using a browser that doesn't support the `light-dark` CSS function. The fallback is the browser's default color scheme.
