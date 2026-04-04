---
title: "Maudie"
layout: layouts/films.njk
slug: maudie-2016
ogImage: content/bill/films/backdrops/maudie-2016.jpg
description: "Canadian folk artist Maud Lewis falls in love with a fishmonger while working for him as a live-in housekeeper."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../taxi-2015"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">58 / 100</a>
  </div>
  <div class="next">
    <a href="../the-handmaiden-2016">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      Taxi
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      The Handmaiden
    </span>
  </div>
</nav>

<article class="film slug-maudie-2016">
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
    <span class="cast-card-character">Maud Lewis</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/569.jpg" alt="Ethan Hawke" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Ethan Hawke</span>
    <span class="cast-card-character">Everett Lewis</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/58804.jpg" alt="Gabrielle Rose" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Gabrielle Rose</span>
    <span class="cast-card-character">Aunt Ida</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/63814.jpg" alt="Billy MacLellan" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Billy MacLellan</span>
    <span class="cast-card-character">Frank</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/5918.jpg" alt="Zachary Bennett" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Zachary Bennett</span>
    <span class="cast-card-character">Charles Dowley</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/5937.jpg" alt="Kari Matchett" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Kari Matchett</span>
    <span class="cast-card-character">Sandra</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1843639.jpg" alt="David Feehan" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">David Feehan</span>
    <span class="cast-card-character">Paul</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1670055.jpg" alt="Lawrence Barry" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Lawrence Barry</span>
    <span class="cast-card-character">Mr. Davis (Shopkeeper)</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1374432.jpg" alt="Marthe Bernard" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Marthe Bernard</span>
    <span class="cast-card-character">Kay</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1278477.jpg" alt="Greg Malone" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Greg Malone</span>
    <span class="cast-card-character">Mr. Hill</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1196541.jpg" alt="Nik Sexton" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Nik Sexton</span>
    <span class="cast-card-character">Steven (CBC Reporter)</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1877614.jpg" alt="Brian Marler" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Brian Marler</span>
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
    <li><a href="../eternal-beauty-2020">Eternal Beauty</a> because of Sally Hawkins</li>
  </ul>
</section>

</article>
