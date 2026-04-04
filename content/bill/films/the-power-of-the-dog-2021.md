---
title: "The Power of the Dog"
layout: layouts/films.njk
slug: the-power-of-the-dog-2021
ogImage: content/bill/films/backdrops/the-power-of-the-dog-2021.jpg
description: "A domineering but charismatic rancher wages a war of intimidation on his brother's new wife and her teen son, until long-hidden secrets come to light."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../the-french-dispatch-2021"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">85 / 100</a>
  </div>
  <div class="next">
    <a href="../the-tragedy-of-macbeth-2021">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      The French Dispatch
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      The Tragedy of Macbeth
    </span>
  </div>
</nav>

<article class="film slug-the-power-of-the-dog-2021">
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

  <section class="cast-grid">
  <div class="cast-grid-cards">
    <div class="cast-card">
  <img src="../films/profiles/71580.jpg" alt="Benedict Cumberbatch" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Benedict Cumberbatch</span>
    <span class="cast-card-character">Phil Burbank</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/113505.jpg" alt="Kodi Smit-McPhee" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Kodi Smit-McPhee</span>
    <span class="cast-card-character">Peter Gordon</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/205.jpg" alt="Kirsten Dunst" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Kirsten Dunst</span>
    <span class="cast-card-character">Rose Gordon</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/88124.jpg" alt="Jesse Plemons" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Jesse Plemons</span>
    <span class="cast-card-character">George Burbank</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1356758.jpg" alt="Thomasin McKenzie" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Thomasin McKenzie</span>
    <span class="cast-card-character">Lola</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/10760.jpg" alt="Geneviève Lemon" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Geneviève Lemon</span>
    <span class="cast-card-character">Mrs. Lewis</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/30613.jpg" alt="Keith Carradine" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Keith Carradine</span>
    <span class="cast-card-character">The Governor</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/4432.jpg" alt="Frances Conroy" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Frances Conroy</span>
    <span class="cast-card-character">Old Lady</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/57904.jpg" alt="Kenneth Radley" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Kenneth Radley</span>
    <span class="cast-card-character">Barkeep</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/962109.jpg" alt="Sean Keenan" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Sean Keenan</span>
    <span class="cast-card-character">Sven</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1399811.jpg" alt="George Mason" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">George Mason</span>
    <span class="cast-card-character">Cricket</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">Ramontay McConnell</span>
    <span class="cast-card-character">Theo</span>
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
    <li><a href="../1917-2019">1917</a> because of Benedict Cumberbatch</li>
  </ul>
</section>

</article>
