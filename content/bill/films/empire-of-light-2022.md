---
title: "Empire of Light"
layout: layouts/films.njk
slug: empire-of-light-2022
ogImage: content/bill/films/backdrops/empire-of-light-2022.jpg
description: "The duty manager of a seaside cinema, who is struggling with her mental health, forms a relationship with a new employee on the south coast of England in the 1980s."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../eo-2022"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">89 / 100</a>
  </div>
  <div class="next">
    <a href="../one-fine-morning-2022">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      EO
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      One Fine Morning
    </span>
  </div>
</nav>

<article class="film slug-empire-of-light-2022">
  <div class="backdrop-and-poster">
    <img class="poster" src="../films/posters/{{ slug }}.jpg" alt="" eleventy:ignore>
    <img class="backdrop" src="../films/backdrops/{{ slug }}.jpg" alt="" eleventy:ignore>
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
  <img src="../films/profiles/39187.jpg" alt="Olivia Colman" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Olivia Colman</span>
    <span class="cast-card-character">Hilary</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/2380397.jpg" alt="Micheal Ward" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Micheal Ward</span>
    <span class="cast-card-character">Stephen</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/13014.jpg" alt="Toby Jones" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Toby Jones</span>
    <span class="cast-card-character">Norman</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/5472.jpg" alt="Colin Firth" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Colin Firth</span>
    <span class="cast-card-character">Donald Ellis</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/75066.jpg" alt="Tom Brooke" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Tom Brooke</span>
    <span class="cast-card-character">Neil</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/31427.jpg" alt="Tanya Moodie" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Tanya Moodie</span>
    <span class="cast-card-character">Delia</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/3260026.jpg" alt="Hannah Onslow" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Hannah Onslow</span>
    <span class="cast-card-character">Janine</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1371518.jpg" alt="Crystal Clarke" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Crystal Clarke</span>
    <span class="cast-card-character">Ruby</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/229606.jpg" alt="Monica Dolan" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Monica Dolan</span>
    <span class="cast-card-character">Rosemary Bates</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/5319.jpg" alt="Ron Cook" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Ron Cook</span>
    <span class="cast-card-character">Mr. Cooper</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/3902.jpg" alt="Sara Stewart" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Sara Stewart</span>
    <span class="cast-card-character">Brenda Ellis</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/206724.jpg" alt="Justin Edwards" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Justin Edwards</span>
    <span class="cast-card-character">Jim Booth</span>
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
    <li><a href="../24-hour-party-people-2002">24 Hour Party People</a> because of Ron Cook</li>
<li><a href="../hot-fuzz-2007">Hot Fuzz</a> because of Ron Cook and Olivia Colman</li>
<li><a href="../1917-2019">1917</a> because of Colin Firth, Justin Edwards, Spike Leighton and Sam Mendes</li>
  </ul>
</section>

</article>
