---
title: "Woman at War"
layout: layouts/films.njk
slug: woman-at-war-2018
ogImage: content/bill/films/backdrops/woman-at-war-2018.jpg
description: "Halla declares a one-woman-war on the local aluminium industry. She is prepared to risk everything to protect the pristine Icelandic Highlands she loves… Until an orphan unexpectedly enters her life."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../sink-or-swim-2018"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">71 / 100</a>
  </div>
  <div class="next">
    <a href="../parasite-2019">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      Sink or Swim
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      Parasite
    </span>
  </div>
</nav>

<article class="film slug-woman-at-war-2018">
  <div class="backdrop-and-poster">
    <img class="poster" src="../films/posters/{{ slug }}.jpg" alt="">
    <img class="backdrop" src="../films/backdrops/{{ slug }}.jpg" alt="">
  </div>

  <h1>{{ film.title }} ({{ film | filmYear }})</h1>

  <p>
    {%- if film.language -%}Language: {{ film.language }}.{% endif %}
    Also known as Kona fer í stríð.
  </p>

  <p class="director">
    Directed by <strong>{{ film | directors }}</strong>
  </p>

  {%- if films.reviews[slug] -%}
    <blockquote> 
      {{ films.reviews[slug] | safe }} <em>—&nbsp;<a href="/bill">Bill</a></em>
    </blockquote> 
  {%- endif -%}

  <section class="film-detail">
    <div>
      <details>
        <summary>
          <i class="fa-solid fa-masks-theater"></i>
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
          <i class="fa-solid fa-clapperboard"></i>
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
    </div>
  </section>
</article>
