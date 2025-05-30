---
title: "What's Eating Gilbert Grape"
layout: layouts/home.njk
slug: whats-eating-gilbert-grape-1993
---

{% set film = films[slug] %}

<nav class="films">
  <a class="prev" href="../the-fugitive-1993">Previous</a>
  <a href="../">Film list</a>
  <a class="next" href="../four-weddings-and-a-funeral-1994">Next</a>
</nav>

<p>33 / 100</p>

<article class="film">
  <div class="backdrop-and-poster">
    <img class="poster" src="../films/posters/{{ film.slug }}.jpg" alt="">
    <img class="backdrop" src="../films/backdrops/{{ film.slug }}.jpg" alt="">
  </div>

  <h1>{{ film.title }} ({{ film | filmYear }})</h1>

  

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
