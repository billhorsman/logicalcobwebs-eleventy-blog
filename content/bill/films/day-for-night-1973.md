---
title: "Day for Night"
layout: layouts/films.njk
slug: day-for-night-1973
ogImage: content/bill/films/backdrops/day-for-night-1973.jpg
description: "A committed film director struggles to complete his movie while coping with a myriad of crises, personal and professional, among the cast and crew."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../the-sting-1973"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">13 / 100</a>
  </div>
  <div class="next">
    <a href="../dog-day-afternoon-1975">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      The Sting
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      Dog Day Afternoon
    </span>
  </div>
</nav>

<article class="film slug-day-for-night-1973">
  <div class="backdrop-and-poster">
    <img class="poster" src="../films/posters/{{ slug }}.jpg" alt="">
    <img class="backdrop" src="../films/backdrops/{{ slug }}.jpg" alt="">
  </div>

  <h1>{{ film.title }} ({{ film | filmYear }})</h1>

  <p>
    {%- if film.language -%}Language: {{ film.language }}.{% endif %}
    Also known as La Nuit américaine.
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
    <li><a href="../bullitt-1968">Bullitt</a> because of Jacqueline Bisset</li>
<li><a href="../le-havre-2011">Le Havre</a> because of Jean-Pierre Léaud</li>
<li><a href="../man-on-the-train-2002">Man on the Train</a> because of Jean-François Stévenin</li>
  </ul>
</section>

</article>
