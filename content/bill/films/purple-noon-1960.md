---
title: "Purple Noon"
layout: layouts/films.njk
slug: purple-noon-1960
ogImage: content/bill/films/backdrops/purple-noon-1960.jpg
description: "Tom Ripley is a talented mimic, moocher, forger and all-around criminal improviser; but there's more to Tom Ripley than even he can guess."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../la-dolce-vita-1960"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">8 / 100</a>
  </div>
  <div class="next">
    <a href="../barefoot-in-the-park-1967">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      La Dolce Vita
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      Barefoot in the Park
    </span>
  </div>
</nav>

<article class="film slug-purple-noon-1960">
  <div class="backdrop-and-poster">
    <img class="poster" src="../films/posters/{{ slug }}.jpg" alt="">
    <img class="backdrop" src="../films/backdrops/{{ slug }}.jpg" alt="">
  </div>

  <h1>{{ film.title }} ({{ film | filmYear }})</h1>

  <p>
    {%- if film.language -%}Language: {{ film.language }}.{% endif %}
    Also known as Plein soleil.
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
