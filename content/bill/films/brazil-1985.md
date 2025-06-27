---
title: "Brazil"
layout: layouts/films.njk
slug: brazil-1985
ogImage: content/bill/films/backdrops/brazil-1985.jpg
description: "Low-level bureaucrat Sam Lowry escapes the monotony of his day-to-day life through a recurring daydream of himself as a virtuous hero saving a beautiful damsel. Investigating a case that led to the wrongful arrest and eventual death of an innocent man instead of wanted terrorist Harry Tuttle, he meets the woman from his daydream, and in trying to help her gets caught in a web of mistaken identities, mindless bureaucracy and lies."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../paris-texas-1984"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">25 / 100</a>
  </div>
  <div class="next">
    <a href="../withnail--i-1987">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      Paris, Texas
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      Withnail & I
    </span>
  </div>
</nav>

<article class="film slug-brazil-1985">
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
  <li><a href="../the-deer-hunter-1978">The Deer Hunter</a> (1978) by Robert De Niro</li>
<li><a href="../killers-of-the-flower-moon-2023">Killers of the Flower Moon</a> (2023) by Robert De Niro</li>
<li><a href="../hot-fuzz-2007">Hot Fuzz</a> (2007) by Jim Broadbent</li>
<li><a href="../mr-turner-2014">Mr. Turner</a> (2014) by Roger Ashton-Griffiths</li>
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
