---
title: "The Lighthouse"
layout: layouts/films.njk
slug: the-lighthouse-2019
ogImage: content/bill/films/backdrops/the-lighthouse-2019.jpg
description: "Two lighthouse keepers try to maintain their sanity while living on a remote and mysterious New England island in the 1890s."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../portrait-of-a-lady-on-fire-2019"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">70 / 100</a>
  </div>
  <div class="next">
    <a href="../uncut-gems-2019">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      Portrait of a Lady on Fire
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      Uncut Gems
    </span>
  </div>
</nav>

<article class="film slug-the-lighthouse-2019">
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
  <img src="../films/profiles/11288.jpg" alt="Robert Pattinson" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Robert Pattinson</span>
    <span class="cast-card-character">Thomas Howard</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/5293.jpg" alt="Willem Dafoe" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Willem Dafoe</span>
    <span class="cast-card-character">Thomas Wake</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/2309944.jpg" alt="Valeriia Karaman" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Valeriia Karaman</span>
    <span class="cast-card-character">Mermaid</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/2507696.jpg" alt="Logan Hawkes" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Logan Hawkes</span>
    <span class="cast-card-character">Ephraim Winslow</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/2507697.jpg" alt="Kyla Nicolle" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Kyla Nicolle</span>
    <span class="cast-card-character">Woman on the Rocks</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">Shaun Clarke</span>
    <span class="cast-card-character">Departing Wickie</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">Pierre Richard</span>
    <span class="cast-card-character">Departing Assistant Wickie</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">Preston Hudson</span>
    <span class="cast-card-character">Tender Mate</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">Jeff Cruts</span>
    <span class="cast-card-character">Tender Mate</span>
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
    <li><a href="../fantastic-mr-fox-2009">Fantastic Mr. Fox</a>, <a href="../the-grand-budapest-hotel-2014">The Grand Budapest Hotel</a>, <a href="../the-french-dispatch-2021">The French Dispatch</a> and <a href="../asteroid-city-2023">Asteroid City</a> because of Willem Dafoe</li>
  </ul>
</section>

</article>
