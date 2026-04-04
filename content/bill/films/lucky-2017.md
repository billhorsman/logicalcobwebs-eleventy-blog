---
title: "Lucky"
layout: layouts/films.njk
slug: lucky-2017
ogImage: content/bill/films/backdrops/lucky-2017.jpg
description: "Follows the journey of a 90-year-old atheist and the quirky characters that inhabit his off-the-map desert town. He finds himself at the precipice of life, thrust into a journey of self-exploration."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../lady-bird-2017"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">62 / 100</a>
  </div>
  <div class="next">
    <a href="../the-party-2017">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      Lady Bird
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      The Party
    </span>
  </div>
</nav>

<article class="film slug-lucky-2017">
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
  <img src="../films/profiles/5048.jpg" alt="Harry Dean Stanton" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Harry Dean Stanton</span>
    <span class="cast-card-character">Lucky</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/5602.jpg" alt="David Lynch" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">David Lynch</span>
    <span class="cast-card-character">Howard</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/17402.jpg" alt="Ron Livingston" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Ron Livingston</span>
    <span class="cast-card-character">Bobby Lawrence</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/42157.jpg" alt="Ed Begley Jr." loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Ed Begley Jr.</span>
    <span class="cast-card-character">Dr. Kneedler</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/4139.jpg" alt="Tom Skerritt" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Tom Skerritt</span>
    <span class="cast-card-character">Fred</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/8689.jpg" alt="Barry Shabaka Henley" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Barry Shabaka Henley</span>
    <span class="cast-card-character">Joe</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/100318.jpg" alt="James Darren" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">James Darren</span>
    <span class="cast-card-character">Paulie</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/5151.jpg" alt="Beth Grant" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Beth Grant</span>
    <span class="cast-card-character">Elaine</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1753771.jpg" alt="Yvonne Huff Lee" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Yvonne Huff Lee</span>
    <span class="cast-card-character">Loretta</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/87095.jpg" alt="Hugo Armstrong" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Hugo Armstrong</span>
    <span class="cast-card-character">Vincent</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/143171.jpg" alt="Bertila Damas" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Bertila Damas</span>
    <span class="cast-card-character">Bibi</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">Pamela Sparks</span>
    <span class="cast-card-character">Pam</span>
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
    <li><a href="../paris-texas-1984">Paris, Texas</a> because of Harry Dean Stanton</li>
<li><a href="../the-straight-story-1999">The Straight Story</a> because of Harry Dean Stanton and David Lynch</li>
<li><a href="../fargo-1996">Fargo</a> because of John Carroll Lynch</li>
<li><a href="../the-fabelmans-2022">The Fabelmans</a> because of David Lynch</li>
<li><a href="../no-country-for-old-men-2007">No Country for Old Men</a> because of Beth Grant</li>
  </ul>
</section>

</article>
