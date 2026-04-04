---
title: "Lady Bird"
layout: layouts/films.njk
slug: lady-bird-2017
ogImage: content/bill/films/backdrops/lady-bird-2017.jpg
description: "Lady Bird McPherson, a strong willed, deeply opinionated, artistic 17 year old comes of age in Sacramento. Her relationship with her mother and her upbringing are questioned and tested as she plans to head off to college."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../cest-la-vie-2017"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">61 / 100</a>
  </div>
  <div class="next">
    <a href="../lucky-2017">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      C'est la vie!
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      Lucky
    </span>
  </div>
</nav>

<article class="film slug-lady-bird-2017">
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
  <img src="../films/profiles/36592.jpg" alt="Saoirse Ronan" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Saoirse Ronan</span>
    <span class="cast-card-character">Lady Bird McPherson</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/12133.jpg" alt="Laurie Metcalf" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Laurie Metcalf</span>
    <span class="cast-card-character">Marion McPherson</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/72638.jpg" alt="Tracy Letts" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Tracy Letts</span>
    <span class="cast-card-character">Larry McPherson</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1105079.jpg" alt="Lucas Hedges" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Lucas Hedges</span>
    <span class="cast-card-character">Danny O'Neill</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1190668.jpg" alt="Timothée Chalamet" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Timothée Chalamet</span>
    <span class="cast-card-character">Kyle Sheible</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1496393.jpg" alt="Beanie Feldstein" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Beanie Feldstein</span>
    <span class="cast-card-character">Julie Steffans</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/2207.jpg" alt="Lois Smith" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Lois Smith</span>
    <span class="cast-card-character">Sister Sarah Joan</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/196179.jpg" alt="Stephen McKinley Henderson" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Stephen McKinley Henderson</span>
    <span class="cast-card-character">Father Leviatch</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1121786.jpg" alt="Odeya Rush" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Odeya Rush</span>
    <span class="cast-card-character">Jenna Walton</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1231115.jpg" alt="Jordan Rodrigues" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Jordan Rodrigues</span>
    <span class="cast-card-character">Miguel McPherson</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1702762.jpg" alt="Marielle Scott" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Marielle Scott</span>
    <span class="cast-card-character">Shelly Yuhan</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1061132.jpg" alt="John Karna" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">John Karna</span>
    <span class="cast-card-character">Greg Anrue</span>
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
    <li><a href="../fight-club-1999">Fight Club</a> because of Bob Stephenson</li>
<li><a href="../the-grand-budapest-hotel-2014">The Grand Budapest Hotel</a> because of Saoirse Ronan and Lucas Hedges</li>
<li><a href="../little-women-2019">Little Women</a> because of Saoirse Ronan, Tracy Letts, Timothée Chalamet and Greta Gerwig</li>
<li><a href="../the-french-dispatch-2021">The French Dispatch</a> because of Saoirse Ronan, Timothée Chalamet and Lois Smith</li>
<li><a href="../dune-2021">Dune</a> because of Timothée Chalamet and Stephen McKinley Henderson</li>
  </ul>
</section>

</article>
