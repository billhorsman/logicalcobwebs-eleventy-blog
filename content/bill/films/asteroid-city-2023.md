---
title: "Asteroid City"
layout: layouts/home.njk
slug: asteroid-city-2023
ogImage: content/bill/films/backdrops/asteroid-city-2023.jpg
description: "In an American desert town circa 1955, the itinerary of a Junior Stargazer/Space Cadet convention is spectacularly disrupted by world-changing events."
---

{% set film = films[slug] %}

<nav class="films">
  <a class="prev" href="../blue-jean-2023">Previous</a>
  <a href="../">Film list</a>
  <a class="next" href="../killers-of-the-flower-moon-2023">Next</a>
</nav>

<p>98 / 100</p>

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
