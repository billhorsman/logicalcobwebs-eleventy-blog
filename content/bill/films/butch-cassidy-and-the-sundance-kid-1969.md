---
title: "Butch Cassidy and the Sundance Kid"
layout: layouts/films.njk
slug: butch-cassidy-and-the-sundance-kid-1969
ogImage: content/bill/films/backdrops/butch-cassidy-and-the-sundance-kid-1969.jpg
description: "As the west rapidly becomes civilized, a pair of outlaws in 1890s Wyoming find themselves pursued by a posse and decide to flee to South America in hopes of evading the law."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../once-upon-a-time-in-the-west-1968"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">12 / 100</a>
  </div>
  <div class="next">
    <a href="../the-sting-1973">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      Once Upon a Time in the West
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      The Sting
    </span>
  </div>
</nav>

<article class="film slug-butch-cassidy-and-the-sundance-kid-1969">
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
    <li><a href="../in-the-heat-of-the-night-1967">In the Heat of the Night</a> because of Timothy Scott</li>
<li><a href="../the-sting-1973">The Sting</a> because of Paul Newman, Robert Redford, Charles Dierkop and George Roy Hill</li>
<li><a href="../three-days-of-the-condor-1975">Three Days of the Condor</a> and <a href="../all-is-lost-2013">All Is Lost</a> because of Robert Redford</li>
<li><a href="../the-big-lebowski-1998">The Big Lebowski</a> because of Sam Elliott</li>
  </ul>
</section>

</article>
