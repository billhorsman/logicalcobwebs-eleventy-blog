---
title: "I'm All Right Jack"
layout: layouts/films.njk
slug: im-all-right-jack-1959
ogImage: content/bill/films/backdrops/im-all-right-jack-1959.jpg
description: "Naive Stanley Windrush returns from the war, his mind set on a successful career in business. Much to his own dismay, he soon finds he has to start from the bottom and work his way up, and also that the management as well as the trade union use him as a tool in their fight for power."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../la-strada-1954"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">4 / 100</a>
  </div>
  <div class="next">
    <a href="../north-by-northwest-1959">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      La Strada
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      North by Northwest
    </span>
  </div>
</nav>

<article class="film slug-im-all-right-jack-1959">
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
