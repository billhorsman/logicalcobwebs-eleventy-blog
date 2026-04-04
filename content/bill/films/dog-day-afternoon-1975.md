---
title: "Dog Day Afternoon"
layout: layouts/films.njk
slug: dog-day-afternoon-1975
ogImage: content/bill/films/backdrops/dog-day-afternoon-1975.jpg
description: "Based on the true story of would-be Brooklyn bank robbers John Wojtowicz and Salvatore Naturile. Sonny and Sal attempt a bank heist which quickly turns sour and escalates into a hostage situation and stand-off with the police. As Sonny's motives for the robbery are slowly revealed and things become more complicated, the heist turns into a media circus."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../day-for-night-1973"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">13 / 100</a>
  </div>
  <div class="next">
    <a href="../three-days-of-the-condor-1975">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      Day for Night
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      Three Days of the Condor
    </span>
  </div>
</nav>

<article class="film slug-dog-day-afternoon-1975">
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
  <img src="../films/profiles/1158.jpg" alt="Al Pacino" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Al Pacino</span>
    <span class="cast-card-character">Sonny</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/3096.jpg" alt="John Cazale" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">John Cazale</span>
    <span class="cast-card-character">Sal</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1466.jpg" alt="Charles Durning" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Charles Durning</span>
    <span class="cast-card-character">Moretti</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/14541.jpg" alt="Chris Sarandon" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Chris Sarandon</span>
    <span class="cast-card-character">Leon</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/14542.jpg" alt="James Broderick" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">James Broderick</span>
    <span class="cast-card-character">Sheldon</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/14544.jpg" alt="Penelope Allen" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Penelope Allen</span>
    <span class="cast-card-character">Sylvia</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/14545.jpg" alt="Sully Boyar" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Sully Boyar</span>
    <span class="cast-card-character">Mulvaney</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">Beulah Garrick</span>
    <span class="cast-card-character">Margaret</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/10556.jpg" alt="Carol Kane" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Carol Kane</span>
    <span class="cast-card-character">Jenny</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">Sandra Kazan</span>
    <span class="cast-card-character">Deborah</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/14548.jpg" alt="Marcia Jean Kurtz" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Marcia Jean Kurtz</span>
    <span class="cast-card-character">Miriam</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">Amy Levitt</span>
    <span class="cast-card-character">Maria</span>
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
    <li><a href="../the-sting-1973">The Sting</a> because of Charles Durning</li>
<li><a href="../house-of-gucci-2021">House of Gucci</a> because of Al Pacino</li>
<li><a href="../the-deer-hunter-1978">The Deer Hunter</a> because of John Cazale</li>
  </ul>
</section>

</article>
