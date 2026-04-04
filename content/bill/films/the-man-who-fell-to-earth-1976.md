---
title: "The Man Who Fell to Earth"
layout: layouts/films.njk
slug: the-man-who-fell-to-earth-1976
ogImage: content/bill/films/backdrops/the-man-who-fell-to-earth-1976.jpg
description: "Thomas Jerome Newton is an alien who has come to Earth in search of water to save his home planet. Aided by lawyer Oliver Farnsworth, Thomas uses his knowledge of advanced technology to create profitable inventions. While developing a method to transport water, Thomas meets Mary-Lou, a quiet hotel clerk, and begins to fall in love with her. Just as he is ready to leave Earth, Thomas is intercepted by the U.S. government, and his entire plan is threatened."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../three-days-of-the-condor-1975"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">15 / 100</a>
  </div>
  <div class="next">
    <a href="../the-deer-hunter-1978">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      Three Days of the Condor
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      The Deer Hunter
    </span>
  </div>
</nav>

<article class="film slug-the-man-who-fell-to-earth-1976">
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
  <img src="../films/profiles/7487.jpg" alt="David Bowie" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">David Bowie</span>
    <span class="cast-card-character">Thomas Jerome Newton</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/9626.jpg" alt="Rip Torn" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Rip Torn</span>
    <span class="cast-card-character">Nathan Bryce</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/12407.jpg" alt="Candy Clark" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Candy Clark</span>
    <span class="cast-card-character">Mary-Lou</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/14904.jpg" alt="Tony Mascia" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Tony Mascia</span>
    <span class="cast-card-character">Arthur</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/7795.jpg" alt="Buck Henry" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Buck Henry</span>
    <span class="cast-card-character">Oliver Farnsworth</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/10939.jpg" alt="Bernie Casey" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Bernie Casey</span>
    <span class="cast-card-character">Peters</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/105335.jpg" alt="Adrienne La Russa" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Adrienne La Russa</span>
    <span class="cast-card-character">Helen</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/45469.jpg" alt="Claudia Jennings" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Claudia Jennings</span>
    <span class="cast-card-character">Peters' Wife</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/14908.jpg" alt="Rick Riccardo" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Rick Riccardo</span>
    <span class="cast-card-character">Trevor</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/7852.jpg" alt="Jim Lovell" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Jim Lovell</span>
    <span class="cast-card-character">James Lovell</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">Hilary Holland</span>
    <span class="cast-card-character">Jill</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">Linda Hutton</span>
    <span class="cast-card-character">Elaine</span>
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

  
</article>
