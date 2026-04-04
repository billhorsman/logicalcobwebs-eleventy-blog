---
title: "House of Gucci"
layout: layouts/films.njk
slug: house-of-gucci-2021
ogImage: content/bill/films/backdrops/house-of-gucci-2021.jpg
description: "When Patrizia Reggiani, an outsider from humble beginnings, marries into the Gucci family, her unbridled ambition begins to unravel the family legacy and triggers a reckless spiral of betrayal, decadence, revenge, and ultimately… murder."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../dune-2021"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">79 / 100</a>
  </div>
  <div class="next">
    <a href="../licorice-pizza-2021">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      Dune
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      Licorice Pizza
    </span>
  </div>
</nav>

<article class="film slug-house-of-gucci-2021">
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
  <img src="../films/profiles/237405.jpg" alt="Lady Gaga" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Lady Gaga</span>
    <span class="cast-card-character">Patrizia Reggiani</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1023139.jpg" alt="Adam Driver" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Adam Driver</span>
    <span class="cast-card-character">Maurizio Gucci</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1158.jpg" alt="Al Pacino" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Al Pacino</span>
    <span class="cast-card-character">Aldo Gucci</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/16940.jpg" alt="Jeremy Irons" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Jeremy Irons</span>
    <span class="cast-card-character">Rodolfo Gucci</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/7499.jpg" alt="Jared Leto" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Jared Leto</span>
    <span class="cast-card-character">Paolo Gucci</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/54738.jpg" alt="Jack Huston" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Jack Huston</span>
    <span class="cast-card-character">Domenico De Sole</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/3136.jpg" alt="Salma Hayek Pinault" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Salma Hayek Pinault</span>
    <span class="cast-card-character">Pina Auriemma</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/124682.jpg" alt="Alexia Murray" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Alexia Murray</span>
    <span class="cast-card-character">Silvana Reggiani</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/26669.jpg" alt="Vincent Riotta" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Vincent Riotta</span>
    <span class="cast-card-character">Fernando Reggiani</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/35103.jpg" alt="Gaetano Bruno" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Gaetano Bruno</span>
    <span class="cast-card-character">Franco</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1348474.jpg" alt="Camille Cottin" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Camille Cottin</span>
    <span class="cast-card-character">Paola Franchi</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1264979.jpg" alt="Youssef Kerkour" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Youssef Kerkour</span>
    <span class="cast-card-character">Nemir Kirdar</span>
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
    <li><a href="../dog-day-afternoon-1975">Dog Day Afternoon</a> because of Al Pacino</li>
<li><a href="../blade-runner-1982">Blade Runner</a> and <a href="../black-hawk-down-2001">Black Hawk Down</a> because of Ridley Scott</li>
<li><a href="../fight-club-1999">Fight Club</a>, <a href="../phone-booth-2003">Phone Booth</a> and <a href="../dallas-buyers-club-2013">Dallas Buyers Club</a> because of Jared Leto</li>
<li><a href="../the-talented-mr-ripley-1999">The Talented Mr. Ripley</a> because of Pietro Ragusa</li>
  </ul>
</section>

</article>
