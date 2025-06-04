---
title: "Lucky"
layout: layouts/films.njk
slug: lucky-2017
ogImage: content/bill/films/backdrops/lucky-2017.jpg
description: "Follows the journey of a 90-year-old atheist and the quirky characters that inhabit his off-the-map desert town. He finds himself at the precipice of life, thrust into a journey of self-exploration."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../cest-la-vie-2017"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">68 / 100</a>
  </div>
  <div class="next">
    <a href="../the-party-2017">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
</nav>

<article class="film slug-lucky-2017">
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
