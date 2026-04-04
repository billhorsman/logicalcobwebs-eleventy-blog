---
title: "Paris, Texas"
layout: layouts/films.njk
slug: paris-texas-1984
ogImage: content/bill/films/backdrops/paris-texas-1984.jpg
description: "A man wanders out of the desert not knowing who he is. His brother finds him, and helps to pull his memory back of the life he led before he walked out on his family and disappeared four years earlier."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../local-hero-1983"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">21 / 100</a>
  </div>
  <div class="next">
    <a href="../brazil-1985">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      Local Hero
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      Brazil
    </span>
  </div>
</nav>

<article class="film slug-paris-texas-1984">
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
  <img src="../films/profiles/5048.jpg" alt="Harry Dean Stanton" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Harry Dean Stanton</span>
    <span class="cast-card-character">Travis</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/2630.jpg" alt="Nastassja Kinski" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Nastassja Kinski</span>
    <span class="cast-card-character">Jane</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/923.jpg" alt="Dean Stockwell" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Dean Stockwell</span>
    <span class="cast-card-character">Walt</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/9892.jpg" alt="Hunter Carson" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Hunter Carson</span>
    <span class="cast-card-character">Hunter</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/9893.jpg" alt="Aurore Clément" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Aurore Clément</span>
    <span class="cast-card-character">Anne</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/9894.jpg" alt="Bernhard Wicki" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Bernhard Wicki</span>
    <span class="cast-card-character">Doctor Ulmer</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">Sam Berry</span>
    <span class="cast-card-character">Gas Station Attendant</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">Claresie Mobley</span>
    <span class="cast-card-character">Car Rental Clerk</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/54754.jpg" alt="Viva" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Viva</span>
    <span class="cast-card-character">Woman on TV</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">Socorro Valdez</span>
    <span class="cast-card-character">Carmelita</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">Edward Fayton</span>
    <span class="cast-card-character">Hunter's Friend</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">Justin Hogg</span>
    <span class="cast-card-character">Hunter (Age 3)</span>
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
    <li><a href="../apocalypse-now-1979">Apocalypse Now</a> because of Aurore Clément</li>
<li><a href="../the-straight-story-1999">The Straight Story</a> and <a href="../lucky-2017">Lucky</a> because of Harry Dean Stanton</li>
<li><a href="../perfect-days-2023">Perfect Days</a> because of Wim Wenders</li>
  </ul>
</section>

</article>
