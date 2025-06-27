---
title: "North by Northwest"
layout: layouts/films.njk
slug: north-by-northwest-1959
ogImage: content/bill/films/backdrops/north-by-northwest-1959.jpg
description: "Advertising man Roger Thornhill is mistaken for a spy, triggering a deadly cross-country chase."
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
    <a href="../breathless-1960">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      La Strada
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      Breathless
    </span>
  </div>
</nav>

<article class="film slug-north-by-northwest-1959">
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

  <p class="related-films">Related to:</p>
  <ul class="related-films">
  <li><a href="../its-a-wonderful-life-1946">It's a Wonderful Life</a> by Finn Zirzow</li>
<li><a href="../bullitt-1968">Bullitt</a> by Paul Genge</li>
<li><a href="../parasite-2019">Parasite</a> by Alfred Hitchcock</li>
<li><a href="../the-sting-1973">The Sting</a> by Alexander Lockwood and Arthur Tovey</li>
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
