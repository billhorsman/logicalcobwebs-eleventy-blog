---
title: "Perfect Days"
layout: layouts/films.njk
slug: perfect-days-2023
ogImage: content/bill/films/backdrops/perfect-days-2023.jpg
description: "Hirayama is content with his life as a toilet cleaner in Tokyo. Outside of his structured routine, he cherishes music on cassette tapes, books, and taking photos of trees. Through unexpected encounters, he reflects on finding beauty in the world."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../all-of-us-strangers-2023"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">95 / 100</a>
  </div>
  <div class="next">
    <a href="../asteroid-city-2023">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      All of Us Strangers
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      Asteroid City
    </span>
  </div>
</nav>

<article class="film slug-perfect-days-2023">
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
  <img src="../films/profiles/18056.jpg" alt="Koji Yakusho" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Koji Yakusho</span>
    <span class="cast-card-character">Hirayama</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1015822.jpg" alt="Tokio Emoto" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Tokio Emoto</span>
    <span class="cast-card-character">Takashi</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/4011271.jpg" alt="Arisa Nakano" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Arisa Nakano</span>
    <span class="cast-card-character">Niko</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/3086230.jpg" alt="Aoi Yamada" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Aoi Yamada</span>
    <span class="cast-card-character">Aya</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/79457.jpg" alt="Yumi Aso" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Yumi Aso</span>
    <span class="cast-card-character">Keiko</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1630770.jpg" alt="Sayuri Ishikawa" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Sayuri Ishikawa</span>
    <span class="cast-card-character">Mama</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/33134.jpg" alt="Tomokazu Miura" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Tomokazu Miura</span>
    <span class="cast-card-character">Tomoyama</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/136191.jpg" alt="Min Tanaka" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Min Tanaka</span>
    <span class="cast-card-character">Homeless</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">Miyako Tanaka</span>
    <span class="cast-card-character">Old Lady with Brush</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/2212955.jpg" alt="Ron Mizuma" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Ron Mizuma</span>
    <span class="cast-card-character">Businessman</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/4500898.jpg" alt="Soraji Shibuya" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Soraji Shibuya</span>
    <span class="cast-card-character">Kid</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">Aoi Iwasaki</span>
    <span class="cast-card-character">Kid</span>
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
    <li><a href="../paris-texas-1984">Paris, Texas</a> because of Wim Wenders</li>
  </ul>
</section>

</article>
