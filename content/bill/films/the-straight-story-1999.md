---
title: "The Straight Story"
layout: layouts/home.njk
slug: the-straight-story-1999
ogImage: content/bill/films/backdrops/the-straight-story-1999.jpg
description: "A retired farmer and widower in his 70s, Alvin Straight learns one day that his distant brother Lyle has suffered a stroke and may not recover. Alvin is determined to make things right with Lyle while he still can, but his brother lives in Wisconsin, while Alvin is stuck in Iowa with no car and no driver's license. Then he hits on the idea of making the trip on his old lawnmower, thus beginning a picturesque and at times deeply spiritual odyssey."
---

{% set film = films[slug] %}

<nav class="films">
  <a class="prev" href="../fight-club-1999">Previous</a>
  <a href="../">Film list</a>
  <a class="next" href="../magnolia-1999">Next</a>
</nav>

<p>40 / 100</p>

<article class="film slug-the-straight-story-1999">
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
