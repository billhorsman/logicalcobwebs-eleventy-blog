---
title: "Hot Fuzz"
layout: layouts/films.njk
slug: hot-fuzz-2007
ogImage: content/bill/films/backdrops/hot-fuzz-2007.jpg
description: "Former London constable Nicholas Angel finds it difficult to adapt to his new assignment in the sleepy British village of Sandford. Not only does he miss the excitement of the big city, but he also has a well-meaning oaf for a partner. However, when a series of grisly accidents rocks Sandford, Angel smells something rotten in the idyllic village."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../the-motorcycle-diaries-2004"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">48 / 100</a>
  </div>
  <div class="next">
    <a href="../no-country-for-old-men-2007">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      The Motorcycle Diaries
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      No Country for Old Men
    </span>
  </div>
</nav>

<article class="film slug-hot-fuzz-2007">
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
    <li><a href="../brazil-1985">Brazil</a> because of Jim Broadbent</li>
<li><a href="../the-talented-mr-ripley-1999">The Talented Mr. Ripley</a> because of Cate Blanchett</li>
<li><a href="../24-hour-party-people-2002">24 Hour Party People</a> because of Steve Coogan, Paddy Considine, Simon Pegg and Ron Cook</li>
<li><a href="../empire-of-light-2022">Empire of Light</a> because of Ron Cook and Olivia Colman</li>
<li><a href="../mr-turner-2014">Mr. Turner</a> because of Karl Johnson and Peter Wight</li>
<li><a href="../eternal-beauty-2020">Eternal Beauty</a> because of Alice Lowe</li>
<li><a href="../fantastic-mr-fox-2009">Fantastic Mr. Fox</a> because of Garth Jennings</li>
  </ul>
</section>

</article>
