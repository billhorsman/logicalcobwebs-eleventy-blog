---
title: "Dallas Buyers Club"
layout: layouts/home.njk
slug: dallas-buyers-club-2013
ogImage: content/bill/films/backdrops/dallas-buyers-club-2013.jpg
description: "Loosely based on the true-life tale of Ron Woodroof, a drug-taking, women-loving, homophobic man who in 1986 was diagnosed with HIV/AIDS and given thirty days to live."
---

{% set film = films[slug] %}

<nav class="films">
  <a class="prev" href="../all-is-lost-2013">Previous</a>
  <a href="../">Film list</a>
  <a class="next" href="../the-grand-budapest-hotel-2014">Next</a>
</nav>

<p>62 / 100</p>

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
        <strong>{{ cast.name }}</strong> as <em>{{ cast.character }}</em>
      </li>
    {%- endfor -%}
  </ul>
</article>
<footer>
  <a href="../about">About this list</a>
</footer>
