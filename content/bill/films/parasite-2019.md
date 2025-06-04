---
title: "Parasite"
layout: layouts/films.njk
slug: parasite-2019
ogImage: content/bill/films/backdrops/parasite-2019.jpg
description: "All unemployed, Ki-taek's family takes peculiar interest in the wealthy and glamorous Parks for their livelihood until they get entangled in an unexpected incident."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../woman-at-war-2018"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">73 / 100</a>
  </div>
  <div class="next">
    <a href="../portrait-of-a-lady-on-fire-2019">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
</nav>

<article class="film slug-parasite-2019">
  <div class="backdrop-and-poster">
    <img class="poster" src="../films/posters/{{ slug }}.jpg" alt="">
    <img class="backdrop" src="../films/backdrops/{{ slug }}.jpg" alt="">
  </div>

  <h1>{{ film.title }} ({{ film | filmYear }})</h1>

  <p>
    {%- if film.language -%}Language: {{ film.language }}.{% endif %}
    Also known as 기생충.
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
