---
title: "Portrait of a Lady on Fire"
layout: layouts/films.njk
slug: portrait-of-a-lady-on-fire-2019
ogImage: content/bill/films/backdrops/portrait-of-a-lady-on-fire-2019.jpg
description: "On an isolated island in Brittany at the end of the eighteenth century, a female painter is obliged to paint a wedding portrait of a young woman."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../parasite-2019"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">74 / 100</a>
  </div>
  <div class="next">
    <a href="../eternal-beauty-2020">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      Parasite
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      Eternal Beauty
    </span>
  </div>
</nav>

<article class="film slug-portrait-of-a-lady-on-fire-2019">
  <div class="backdrop-and-poster">
    <img class="poster" src="../films/posters/{{ slug }}.jpg" alt="">
    <img class="backdrop" src="../films/backdrops/{{ slug }}.jpg" alt="">
  </div>

  <h1>{{ film.title }} ({{ film | filmYear }})</h1>

  <p>
    {%- if film.language -%}Language: {{ film.language }}.{% endif %}
    Also known as Portrait de la jeune fille en feu.
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
