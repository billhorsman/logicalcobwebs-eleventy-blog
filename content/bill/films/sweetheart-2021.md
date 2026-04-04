---
title: "Sweetheart"
layout: layouts/films.njk
slug: sweetheart-2021
ogImage: content/bill/films/backdrops/sweetheart-2021.jpg
description: "A socially awkward, environmentally-conscious teenager named AJ is dragged to a coastal holiday park by her painfully 'normal' family, where she becomes unexpectedly captivated by a chlorine smelling, sun-loving lifeguard named Isla."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../petite-maman-2021"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">83 / 100</a>
  </div>
  <div class="next">
    <a href="../the-french-dispatch-2021">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      Petite Maman
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      The French Dispatch
    </span>
  </div>
</nav>

<article class="film slug-sweetheart-2021">
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
  <img src="../films/profiles/1865843.jpg" alt="Nell Barlow" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Nell Barlow</span>
    <span class="cast-card-character">A.J.</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1770528.jpg" alt="Ella-Rae Smith" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Ella-Rae Smith</span>
    <span class="cast-card-character">Isla</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/70518.jpg" alt="Jo Hartley" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Jo Hartley</span>
    <span class="cast-card-character">Tina</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1252533.jpg" alt="Sophia Di Martino" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Sophia Di Martino</span>
    <span class="cast-card-character">Lucy</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/55465.jpg" alt="Samuel Anderson" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Samuel Anderson</span>
    <span class="cast-card-character">Steve</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/2798731.jpg" alt="Tabitha Byron" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Tabitha Byron</span>
    <span class="cast-card-character">Dayna</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/2108331.jpg" alt="Steffan Cennydd" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Steffan Cennydd</span>
    <span class="cast-card-character">Nathan</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/2304140.jpg" alt="Spike Fearn" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Spike Fearn</span>
    <span class="cast-card-character">Elvis</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1586290.jpg" alt="William Andrews" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">William Andrews</span>
    <span class="cast-card-character">Phil the Magician</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1945730.jpg" alt="Anna Antoniades" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Anna Antoniades</span>
    <span class="cast-card-character">Gemma G</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/2936317.jpg" alt="Celeste De Veazey" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Celeste De Veazey</span>
    <span class="cast-card-character">Bendy Wendy</span>
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
