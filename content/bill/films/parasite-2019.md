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
    <a class="simple" href="../">71 / 100</a>
  </div>
  <div class="next">
    <a href="../portrait-of-a-lady-on-fire-2019">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      Woman at War
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      Portrait of a Lady on Fire
    </span>
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

  <p class="related-films">Related to:</p>
  <ul class="related-films">
  <li><a href="../north-by-northwest-1959">North by Northwest</a> (1959) by Alfred Hitchcock</li>
<li><a href="../the-handmaiden-2016">The Handmaiden</a> (2016) by Lee Ji-hye and Ahn Seong-bong</li>
  </ul>

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
