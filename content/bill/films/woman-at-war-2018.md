---
title: "Woman at War"
layout: layouts/films.njk
slug: woman-at-war-2018
ogImage: content/bill/films/backdrops/woman-at-war-2018.jpg
description: "Halla declares a one-woman-war on the local aluminium industry. She is prepared to risk everything to protect the pristine Icelandic Highlands she loves… Until an orphan unexpectedly enters her life."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../sink-or-swim-2018"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">66 / 100</a>
  </div>
  <div class="next">
    <a href="../little-women-2019">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      Sink or Swim
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      Little Women
    </span>
  </div>
</nav>

<article class="film slug-woman-at-war-2018">
  <div class="backdrop-and-poster">
    <img class="poster" src="../films/posters/{{ slug }}.jpg" alt="">
    <img class="backdrop" src="../films/backdrops/{{ slug }}.jpg" alt="">
  </div>

  <h1>{{ film.title }} ({{ film | filmYear }})</h1>

  <p>
    {%- if film.language -%}Language: {{ film.language }}.{% endif %}
    Also known as Kona fer í stríð.
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
  <img src="../films/profiles/142874.jpg" alt="Halldóra Geirharðsdóttir" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Halldóra Geirharðsdóttir</span>
    <span class="cast-card-character">Halla / Ása</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/110901.jpg" alt="Jóhann Sigurðarson" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Jóhann Sigurðarson</span>
    <span class="cast-card-character">Sveinbjörn</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/544972.jpg" alt="Davíð Þór Jónsson" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Davíð Þór Jónsson</span>
    <span class="cast-card-character">Pianist / Accordion Player</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/2153591.jpg" alt="Magnús Trygvason Eliassen" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Magnús Trygvason Eliassen</span>
    <span class="cast-card-character">Drummer</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">Ómar Guðjónsson</span>
    <span class="cast-card-character">Sousaphone Player</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">Iryna Danyleiko</span>
    <span class="cast-card-character">Ukrainian Choir Singer</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">Galyna Goncharenko</span>
    <span class="cast-card-character">Ukrainian Choir Singer</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">Susanna Kurpenko</span>
    <span class="cast-card-character">Ukrainian Choir Singer</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/118345.jpg" alt="Jörundur Ragnarsson" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Jörundur Ragnarsson</span>
    <span class="cast-card-character">Baldvin</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1290765.jpg" alt="Juan Camillo Roman Estrada" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Juan Camillo Roman Estrada</span>
    <span class="cast-card-character">Juan Camillo</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1183138.jpg" alt="Charlotte Bøving" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Charlotte Bøving</span>
    <span class="cast-card-character">Adoption agency lady</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1093246.jpg" alt="Björn Thors" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Björn Thors</span>
    <span class="cast-card-character">The Prime Minister</span>
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
