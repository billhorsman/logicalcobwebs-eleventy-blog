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
    <a class="simple" href="../">45 / 100</a>
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
    <img class="poster" src="../films/posters/{{ slug }}.jpg" alt="" eleventy:ignore>
    <img class="backdrop" src="../films/backdrops/{{ slug }}.jpg" alt="" eleventy:ignore>
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

  <section class="cast-grid">
  <div class="cast-grid-cards">
    <div class="cast-card">
  <img src="../films/profiles/11108.jpg" alt="Simon Pegg" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Simon Pegg</span>
    <span class="cast-card-character">Nicholas Angel</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/11109.jpg" alt="Nick Frost" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Nick Frost</span>
    <span class="cast-card-character">PC Danny Butterman</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/388.jpg" alt="Jim Broadbent" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Jim Broadbent</span>
    <span class="cast-card-character">Inspector Frank Butterman</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/14887.jpg" alt="Paddy Considine" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Paddy Considine</span>
    <span class="cast-card-character">DS Andy Wainwright</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/28847.jpg" alt="Rafe Spall" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Rafe Spall</span>
    <span class="cast-card-character">DC Andy Cartwright</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/39185.jpg" alt="Kevin Eldon" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Kevin Eldon</span>
    <span class="cast-card-character">Sergeant Tony Fisher</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/39187.jpg" alt="Olivia Colman" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Olivia Colman</span>
    <span class="cast-card-character">PC Doris Thatcher</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/24265.jpg" alt="Bill Bailey" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Bill Bailey</span>
    <span class="cast-card-character">Sergeant Turner</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/39186.jpg" alt="Karl Johnson" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Karl Johnson</span>
    <span class="cast-card-character">PC Bob Walker</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/10669.jpg" alt="Timothy Dalton" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Timothy Dalton</span>
    <span class="cast-card-character">Simon Skinner</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/39188.jpg" alt="Edward Woodward" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Edward Woodward</span>
    <span class="cast-card-character">Tom Weaver</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/8226.jpg" alt="Billie Whitelaw" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Billie Whitelaw</span>
    <span class="cast-card-character">Joyce Cooper</span>
  </div>
</div>
  </div>
</section>

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
