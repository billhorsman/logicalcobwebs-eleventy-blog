---
title: "Limbo"
layout: layouts/films.njk
slug: limbo-2020
ogImage: content/bill/films/backdrops/limbo-2020.jpg
description: "An offbeat observation of refugees waiting to be granted asylum on a fictional remote Scottish island. It focuses on Omar, a young Syrian musician who is burdened by the weight of his grandfather’s oud, which he has carried all the way from his homeland."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../eternal-beauty-2020"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">74 / 100</a>
  </div>
  <div class="next">
    <a href="../the-truffle-hunters-2020">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      Eternal Beauty
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      The Truffle Hunters
    </span>
  </div>
</nav>

<article class="film slug-limbo-2020">
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
  <img src="../films/profiles/1739403.jpg" alt="Amir El-Masry" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Amir El-Masry</span>
    <span class="cast-card-character">Omar</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/2151678.jpg" alt="Vikash Bhai" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Vikash Bhai</span>
    <span class="cast-card-character">Farhad</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/2151679.jpg" alt="Ola Orebiyi" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Ola Orebiyi</span>
    <span class="cast-card-character">Wasef</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">Kwabena Ansah</span>
    <span class="cast-card-character">Abedi</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/32683.jpg" alt="Sidse Babett Knudsen" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Sidse Babett Knudsen</span>
    <span class="cast-card-character">Helga</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/25087.jpg" alt="Qais Nashif" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Qais Nashif</span>
    <span class="cast-card-character"></span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/997569.jpg" alt="Kenneth Collard" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Kenneth Collard</span>
    <span class="cast-card-character">Boris</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1216533.jpg" alt="Sanjeev Kohli" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Sanjeev Kohli</span>
    <span class="cast-card-character">Vikram</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/2180052.jpg" alt="Cameron Fulton" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Cameron Fulton</span>
    <span class="cast-card-character">Plug</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/2240655.jpg" alt="Lewis Gribben" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Lewis Gribben</span>
    <span class="cast-card-character">Stevie</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1636378.jpg" alt="Grace Chilton" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Grace Chilton</span>
    <span class="cast-card-character">Margaret</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/16909.jpg" alt="Raymond Mearns" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Raymond Mearns</span>
    <span class="cast-card-character">Mike</span>
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
