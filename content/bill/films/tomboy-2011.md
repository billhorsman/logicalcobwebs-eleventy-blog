---
title: "Tomboy"
layout: layouts/home.njk
slug: tomboy-2011
ogImage: content/bill/films/backdrops/tomboy-2011.jpg
description: "A French family moves to a new neighborhood with during the summer holidays. The story follows a 10-year-old gender non-conforming child, Laure, who experiments with their gender presentation, adopting the name Mikäel."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../le-havre-2011"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a href="../">Film list</a>
  </div>
  <div class="next">
    <a href="../all-is-lost-2013">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
</nav>

<p>59 / 100</p>

<article class="film slug-tomboy-2011">
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
