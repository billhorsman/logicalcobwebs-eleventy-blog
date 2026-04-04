---
title: "Uncut Gems"
layout: layouts/films.njk
slug: uncut-gems-2019
ogImage: content/bill/films/backdrops/uncut-gems-2019.jpg
description: "A charismatic New York City jeweler always on the lookout for the next big score makes a series of high-stakes bets that could lead to the windfall of a lifetime. Howard must perform a precarious high-wire act, balancing business, family, and encroaching adversaries on all sides in his relentless pursuit of the ultimate win."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../the-lighthouse-2019"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">71 / 100</a>
  </div>
  <div class="next">
    <a href="../1917-2019">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      The Lighthouse
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      1917
    </span>
  </div>
</nav>

<article class="film slug-uncut-gems-2019">
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
  <img src="../films/profiles/19292.jpg" alt="Adam Sandler" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Adam Sandler</span>
    <span class="cast-card-character">Howard Ratner</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1200864.jpg" alt="LaKeith Stanfield" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">LaKeith Stanfield</span>
    <span class="cast-card-character">Demany</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/2371446.jpg" alt="Julia Fox" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Julia Fox</span>
    <span class="cast-card-character">Julia De Flore</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1177652.jpg" alt="Kevin Garnett" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Kevin Garnett</span>
    <span class="cast-card-character">Kevin Garnett</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/19394.jpg" alt="Idina Menzel" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Idina Menzel</span>
    <span class="cast-card-character">Dinah Ratner</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/10866.jpg" alt="Eric Bogosian" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Eric Bogosian</span>
    <span class="cast-card-character">Arno Moradian</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/6167.jpg" alt="Judd Hirsch" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Judd Hirsch</span>
    <span class="cast-card-character">Gooey</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/2410045.jpg" alt="Keith William Richards" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Keith William Richards</span>
    <span class="cast-card-character">Phil</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1315479.jpg" alt="Mike Francesa" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Mike Francesa</span>
    <span class="cast-card-character">Gary</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/2371447.jpg" alt="Jonathan Aranbayev" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Jonathan Aranbayev</span>
    <span class="cast-card-character">Eddie Ratner</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/2371449.jpg" alt="Noa Fisher" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Noa Fisher</span>
    <span class="cast-card-character">Marcel Ratner</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1549008.jpg" alt="The Weeknd" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">The Weeknd</span>
    <span class="cast-card-character">The Weeknd</span>
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
    <li><a href="../the-grand-budapest-hotel-2014">The Grand Budapest Hotel</a> and <a href="../the-french-dispatch-2021">The French Dispatch</a> because of Tilda Swinton</li>
<li><a href="../asteroid-city-2023">Asteroid City</a> because of Tilda Swinton and Jake Ryan</li>
<li><a href="../the-fabelmans-2022">The Fabelmans</a> because of Judd Hirsch</li>
<li><a href="../licorice-pizza-2021">Licorice Pizza</a> because of Benny Safdie</li>
  </ul>
</section>

</article>
