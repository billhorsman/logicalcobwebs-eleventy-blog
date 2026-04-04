---
title: "Nomadland"
layout: layouts/films.njk
slug: nomadland-2021
ogImage: content/bill/films/backdrops/nomadland-2021.jpg
description: "A woman in her sixties embarks on a journey through the western United States after losing everything in the Great Recession, living as a van-dwelling modern-day nomad."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../licorice-pizza-2021"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">81 / 100</a>
  </div>
  <div class="next">
    <a href="../petite-maman-2021">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      Licorice Pizza
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      Petite Maman
    </span>
  </div>
</nav>

<article class="film slug-nomadland-2021">
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
  <img src="../films/profiles/3910.jpg" alt="Frances McDormand" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Frances McDormand</span>
    <span class="cast-card-character">Fern</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/11064.jpg" alt="David Strathairn" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">David Strathairn</span>
    <span class="cast-card-character">Dave</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/2241214.jpg" alt="Linda May" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Linda May</span>
    <span class="cast-card-character">Linda</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/2776341.jpg" alt="Swankie" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Swankie</span>
    <span class="cast-card-character">Swankie</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">Gay DeForest</span>
    <span class="cast-card-character">Gay</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">Patricia Grier</span>
    <span class="cast-card-character">Patty</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">Angela Reyes</span>
    <span class="cast-card-character">Angela</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">Carl R. Hughes</span>
    <span class="cast-card-character">Carl</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">Douglas G. Soul</span>
    <span class="cast-card-character">Doug</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">Ryan Aquino</span>
    <span class="cast-card-character">Ryan</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">Teresa Buchanan</span>
    <span class="cast-card-character">Teresa</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">Karie Lynn McDermott Wilder</span>
    <span class="cast-card-character">Karie</span>
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
    <li><a href="../fargo-1996">Fargo</a> because of Frances McDormand and Warren Keith</li>
<li><a href="../the-french-dispatch-2021">The French Dispatch</a> and <a href="../the-tragedy-of-macbeth-2021">The Tragedy of Macbeth</a> because of Frances McDormand</li>
<li><a href="../the-big-lebowski-1998">The Big Lebowski</a> because of Warren Keith</li>
  </ul>
</section>

</article>
