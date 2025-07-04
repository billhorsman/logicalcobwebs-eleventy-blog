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
    <a href="../little-women-2019"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">70 / 100</a>
  </div>
  <div class="next">
    <a href="../portrait-of-a-lady-on-fire-2019">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      Little Women
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

  <section class="related-films">
  <h2>Related films</h2>
  <ul>
    <li><a href="../rear-window-1954">Rear Window</a> and <a href="../north-by-northwest-1959">North by Northwest</a> because of Alfred Hitchcock</li>
<li><a href="../the-handmaiden-2016">The Handmaiden</a> because of Lee Ji-hye and Ahn Seong-bong</li>
  </ul>
</section>

</article>
