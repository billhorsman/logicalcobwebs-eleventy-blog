---
title: "Barefoot in the Park"
layout: layouts/home.njk
slug: barefoot-in-the-park-1967
ogImage: content/bill/films/backdrops/barefoot-in-the-park-1967.jpg
description: "In this film based on a Neil Simon play, newlyweds Corie, a free spirit, and Paul Bratter, an uptight lawyer, share a sixth-floor apartment in Greenwich Village. Soon after their marriage, Corie tries to find a companion for mother, Ethel, who is now alone, and sets up Ethel with neighbor Victor. Inappropriate behavior on a double date causes conflict, and the young couple considers divorce."
---

{% set film = films[slug] %}

<nav class="films">
  <a class="prev" href="../breathless-1960">Previous</a>
  <a href="../">Film list</a>
  <a class="next" href="../in-the-heat-of-the-night-1967">Next</a>
</nav>

<p>9 / 100</p>

<article class="film">
  <div class="backdrop-and-poster">
    <img class="poster" src="../films/posters/{{ slug }}.jpg" alt="">
    <img class="backdrop" src="../films/backdrops/{{ slug }}.jpg" alt="">
  </div>

  <h1>{{ film.title }} ({{ film | filmYear }})</h1>

  

  <p class="director">
    Directed by <strong>{{ film | directors }}</strong>
  </p>

  {% if films.reviews[slug] %}
    <blockquote> 
      {{ films.reviews[slug] | safe }} <em>— Bill</em>
    </blockquote> 
  {% endif %}

  <h2>
    Cast
  </h2>
  <ul>
    {%- for cast in film.credits.cast -%}
      <li>
        {{ cast.name }} as <em>{{ cast.character }}</em>
      </li>
    {%- endfor -%}
  </ul>

  <h2>
    Crew
  </h2>
  <ul>
    {%- for crew in film.credits.crew -%}
      <li>
        {{ crew.name }} &mdash; <em>{{ crew.job }}</em>
      </li>
    {%- endfor -%}
  </ul>
</article>
<footer>
  <a href="../about">About this list</a>
</footer>
