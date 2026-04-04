---
title: "Fantastic Mr. Fox"
layout: layouts/films.njk
slug: fantastic-mr-fox-2009
ogImage: content/bill/films/backdrops/fantastic-mr-fox-2009.jpg
description: "The Fantastic Mr. Fox, bored with his current life, plans a heist against the three local farmers. The farmers, tired of sharing their chickens with the sly fox, seek revenge against him and his family."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../in-bruges-2008"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">48 / 100</a>
  </div>
  <div class="next">
    <a href="../micmacs-2009">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      In Bruges
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      Micmacs
    </span>
  </div>
</nav>

<article class="film slug-fantastic-mr-fox-2009">
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
  <img src="../films/profiles/1461.jpg" alt="George Clooney" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">George Clooney</span>
    <span class="cast-card-character">Mr. Fox (voice)</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/5064.jpg" alt="Meryl Streep" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Meryl Streep</span>
    <span class="cast-card-character">Felicity Fox (voice)</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/17881.jpg" alt="Jason Schwartzman" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Jason Schwartzman</span>
    <span class="cast-card-character">Ash Fox (voice)</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/486.jpg" alt="Wallace Wolodarsky" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Wallace Wolodarsky</span>
    <span class="cast-card-character">Kylie (voice)</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1332415.jpg" alt="Eric Chase Anderson" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Eric Chase Anderson</span>
    <span class="cast-card-character">Kristofferson Silverfox (voice)</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/5293.jpg" alt="Willem Dafoe" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Willem Dafoe</span>
    <span class="cast-card-character">Rat (voice)</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1532.jpg" alt="Bill Murray" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Bill Murray</span>
    <span class="cast-card-character">Clive Badger (voice)</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1332416.jpg" alt="Robin Hurlstone" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Robin Hurlstone</span>
    <span class="cast-card-character">Walter Boggis (voice)</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/992097.jpg" alt="Hugo Guinness" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Hugo Guinness</span>
    <span class="cast-card-character">Nathan Bunce (voice)</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/5658.jpg" alt="Michael Gambon" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Michael Gambon</span>
    <span class="cast-card-character">Franklin Bean (voice)</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/15737.jpg" alt="Helen McCrory" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Helen McCrory</span>
    <span class="cast-card-character">Mrs. Bean (voice)</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/5655.jpg" alt="Wes Anderson" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Wes Anderson</span>
    <span class="cast-card-character">Stan Weasel (voice)</span>
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
    <li><a href="../the-deer-hunter-1978">The Deer Hunter</a> and <a href="../little-women-2019">Little Women</a> because of Meryl Streep</li>
<li><a href="../the-bourne-identity-2002">The Bourne Identity</a> because of Brian Cox</li>
<li><a href="../hot-fuzz-2007">Hot Fuzz</a> because of Garth Jennings</li>
<li><a href="../the-grand-budapest-hotel-2014">The Grand Budapest Hotel</a> because of Jason Schwartzman, Wallace Wolodarsky, Willem Dafoe, Bill Murray, Robin Hurlstone, Wes Anderson, Owen Wilson and Adrien Brody</li>
<li><a href="../the-french-dispatch-2021">The French Dispatch</a> because of Jason Schwartzman, Wallace Wolodarsky, Willem Dafoe, Bill Murray, Wes Anderson, Jarvis Cocker, Owen Wilson and Adrien Brody</li>
<li><a href="../asteroid-city-2023">Asteroid City</a> because of Jason Schwartzman, Willem Dafoe, Wes Anderson, Jarvis Cocker and Adrien Brody</li>
<li><a href="../the-lighthouse-2019">The Lighthouse</a> because of Willem Dafoe</li>
  </ul>
</section>

</article>
