---
title: "The Sting"
layout: layouts/films.njk
slug: the-sting-1973
ogImage: content/bill/films/backdrops/the-sting-1973.jpg
description: "A novice con man teams up with an acknowledged master to avenge the murder of a mutual friend by pulling off the ultimate big con and swindling a fortune from a big-time mobster."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../butch-cassidy-and-the-sundance-kid-1969"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">14 / 100</a>
  </div>
  <div class="next">
    <a href="../dog-day-afternoon-1975">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      Butch Cassidy and the Sundance Kid
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      Dog Day Afternoon
    </span>
  </div>
</nav>

<article class="film slug-the-sting-1973">
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
  <li><a href="../north-by-northwest-1959">North by Northwest</a> by Alexander Lockwood and Arthur Tovey</li>
<li><a href="../barefoot-in-the-park-1967">Barefoot in the Park</a>, <a href="../three-days-of-the-condor-1975">Three Days of the Condor</a> and <a href="../all-is-lost-2013">All Is Lost</a> by Robert Redford</li>
<li><a href="../butch-cassidy-and-the-sundance-kid-1969">Butch Cassidy and the Sundance Kid</a> by Robert Redford, Paul Newman, Charles Dierkop and George Roy Hill</li>
<li><a href="../in-the-heat-of-the-night-1967">In the Heat of the Night</a> by Larry D. Mann</li>
<li><a href="../dog-day-afternoon-1975">Dog Day Afternoon</a> by Charles Durning</li>
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
