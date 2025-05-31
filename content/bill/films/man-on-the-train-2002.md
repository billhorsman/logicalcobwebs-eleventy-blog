---
title: "Man on the Train"
layout: layouts/home.njk
slug: man-on-the-train-2002
ogImage: content/bill/films/backdrops/man-on-the-train-2002.jpg
description: "A man, Milan steps off a train, into a small French village. As he waits for the day when he will rob the town bank, he runs into an old retired poetry teacher named M. Manesquier. The two men strike up a strange friendship and explore the road not taken, each wanting to live the other's life."
---

{% set film = films[slug] %}

<nav class="films">
  <a class="prev" href="../the-bourne-identity-2002">Previous</a>
  <a href="../">Film list</a>
  <a class="next" href="../phone-booth-2003">Next</a>
</nav>

<p>49 / 100</p>

<article class="film">
  <div class="backdrop-and-poster">
    <img class="poster" src="../films/posters/{{ slug }}.jpg" alt="">
    <img class="backdrop" src="../films/backdrops/{{ slug }}.jpg" alt="">
  </div>

  <h1>{{ film.title }} ({{ film | filmYear }})</h1>

  <p>Also known as <strong>L'Homme du train</strong></p>

  <p class="director">
    Directed by <strong>{{ film | directors }}</strong>
  </p>

  {% if films.reviews[slug] %}
    <blockquote> 
      {{ films.reviews[slug] }} <em>— Bill</em>
    </blockquote> 
  {% endif %}

  <h2>
    Cast
  </h2>
  <ul>
    {%- for cast in film.credits.cast -%}
      <li>
        <strong>{{ cast.name }}</strong> as <em>{{ cast.character }}</em>
      </li>
    {%- endfor -%}
  </ul>
</article>
<footer>
  <a href="../about">About this list</a>
</footer>
