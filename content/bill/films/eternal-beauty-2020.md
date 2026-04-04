---
title: "Eternal Beauty"
layout: layouts/films.njk
slug: eternal-beauty-2020
ogImage: content/bill/films/backdrops/eternal-beauty-2020.jpg
description: "When Jane is rejected by life, she spirals into a chaotic, schizophrenic world where love and normality collide with humorous consequences."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../1917-2019"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">73 / 100</a>
  </div>
  <div class="next">
    <a href="../limbo-2020">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      1917
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      Limbo
    </span>
  </div>
</nav>

<article class="film slug-eternal-beauty-2020">
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
  <img src="../films/profiles/39658.jpg" alt="Sally Hawkins" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Sally Hawkins</span>
    <span class="cast-card-character">Jane</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/11207.jpg" alt="David Thewlis" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">David Thewlis</span>
    <span class="cast-card-character">Mike</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/182327.jpg" alt="Alice Lowe" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Alice Lowe</span>
    <span class="cast-card-character">Alice</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/26076.jpg" alt="Billie Piper" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Billie Piper</span>
    <span class="cast-card-character">Nicola</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1249.jpg" alt="Penelope Wilton" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Penelope Wilton</span>
    <span class="cast-card-character">Vivian</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/47643.jpg" alt="Robert Pugh" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Robert Pugh</span>
    <span class="cast-card-character">Dennis</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1954399.jpg" alt="Banita Sandhu" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Banita Sandhu</span>
    <span class="cast-card-character">Alex</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/134466.jpg" alt="Paul Hilton" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Paul Hilton</span>
    <span class="cast-card-character">Tony</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1321628.jpg" alt="Morfydd Clark" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Morfydd Clark</span>
    <span class="cast-card-character">Young Jane</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/2758853.jpg" alt="Natalie O'Neill" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Natalie O'Neill</span>
    <span class="cast-card-character">Young Nicola</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">Elysia Welch</span>
    <span class="cast-card-character">Young Alice</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">Boyd Clack</span>
    <span class="cast-card-character">Doctor</span>
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
    <li><a href="../the-big-lebowski-1998">The Big Lebowski</a> because of David Thewlis</li>
<li><a href="../hot-fuzz-2007">Hot Fuzz</a> because of Alice Lowe</li>
<li><a href="../maudie-2016">Maudie</a> because of Sally Hawkins</li>
  </ul>
</section>

</article>
