---
title: "Hot Fuzz"
layout: layouts/home.njk
slug: hot-fuzz-2007
ogImage: content/bill/films/backdrops/hot-fuzz-2007.jpg
description: "Former London constable Nicholas Angel finds it difficult to adapt to his new assignment in the sleepy British village of Sandford. Not only does he miss the excitement of the big city, but he also has a well-meaning oaf for a partner. However, when a series of grisly accidents rocks Sandford, Angel smells something rotten in the idyllic village."
---

{% set film = films[slug] %}

<nav class="films">
  <a class="prev" href="../the-motorcycle-diaries-2004">Previous</a>
  <a href="../">Film list</a>
  <a class="next" href="../no-country-for-old-men-2007">Next</a>
</nav>

<p>52 / 100</p>

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
