---
title: "Interstellar"
layout: layouts/films.njk
slug: interstellar-2014
ogImage: content/bill/films/backdrops/interstellar-2014.jpg
description: "The adventures of a group of explorers who make use of a newly discovered wormhole to surpass the limitations on human space travel and conquer the vast distances involved in an interstellar voyage."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../dallas-buyers-club-2013"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">60 / 100</a>
  </div>
  <div class="next">
    <a href="../mr-turner-2014">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      Dallas Buyers Club
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      Mr. Turner
    </span>
  </div>
</nav>

<article class="film slug-interstellar-2014">
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
    <li><a href="../good-will-hunting-1997">Good Will Hunting</a> because of Matt Damon and Casey Affleck</li>
<li><a href="../the-talented-mr-ripley-1999">The Talented Mr. Ripley</a> and <a href="../the-bourne-identity-2002">The Bourne Identity</a> because of Matt Damon</li>
<li><a href="../dallas-buyers-club-2013">Dallas Buyers Club</a> because of Matthew McConaughey</li>
<li><a href="../killers-of-the-flower-moon-2023">Killers of the Flower Moon</a> because of John Lithgow</li>
<li><a href="../dune-2021">Dune</a> and <a href="../the-french-dispatch-2021">The French Dispatch</a> because of Timothée Chalamet</li>
  </ul>
</section>

</article>
