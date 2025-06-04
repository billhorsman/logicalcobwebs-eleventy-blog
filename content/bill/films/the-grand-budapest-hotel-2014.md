---
title: "The Grand Budapest Hotel"
layout: layouts/films.njk
slug: the-grand-budapest-hotel-2014
ogImage: content/bill/films/backdrops/the-grand-budapest-hotel-2014.jpg
description: "The Grand Budapest Hotel tells of a legendary concierge at a famous European hotel between the wars and his friendship with a young employee who becomes his trusted protégé. The story involves the theft and recovery of a priceless Renaissance painting, the battle for an enormous family fortune and the slow and then sudden upheavals that transformed Europe during the first half of the 20th century."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../mr-turner-2014"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">64 / 100</a>
  </div>
  <div class="next">
    <a href="../maudie-2016">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
</nav>

<article class="film slug-the-grand-budapest-hotel-2014">
  <div class="backdrop-and-poster">
    <img class="poster" src="../films/posters/{{ slug }}.jpg" alt="">
    <img class="backdrop" src="../films/backdrops/{{ slug }}.jpg" alt="">
  </div>

  <h1>{{ film.title }} ({{ film | filmYear }})</h1>

  <p>
    {%- if film.language -%}Language: {{ film.language }}.{% endif %}
    
  </p>

  <p class="director">
    Directed by <strong>{{ film | directors }}</strong>
  </p>

  {%- if films.reviews[slug] -%}
    <blockquote> 
      {{ films.reviews[slug] | safe }} <em>—&nbsp;<a href="/bill">Bill</a></em>
    </blockquote> 
  {%- endif -%}

  <details>
    <summary>
      Cast
    </summary>
    <ul>
      {%- for cast in film.credits.cast -%}
        <li>
          {{ cast.name }} as <em>{{ cast.character }}</em>
        </li>
      {%- endfor -%}
    </ul>
  </details>

  <details>
    <summary>
      Crew
    </summary>
    <ul>
      {%- for crew in film.credits.crew -%}
        <li>
          {{ crew.name }} &mdash; <em>{{ crew.job }}</em>
        </li>
      {%- endfor -%}
    </ul>
  </details>

</article>
<footer>
  <a href="../about">About this list</a>
</footer>
