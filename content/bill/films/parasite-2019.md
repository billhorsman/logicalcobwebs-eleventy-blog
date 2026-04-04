---
title: "Parasite"
layout: layouts/films.njk
slug: parasite-2019
ogImage: content/bill/films/backdrops/parasite-2019.jpg
description: "All unemployed, Ki-taek's family takes peculiar interest in the wealthy and glamorous Parks for their livelihood until they get entangled in an unexpected incident."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../little-women-2019"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">68 / 100</a>
  </div>
  <div class="next">
    <a href="../portrait-of-a-lady-on-fire-2019">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      Little Women
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      Portrait of a Lady on Fire
    </span>
  </div>
</nav>

<article class="film slug-parasite-2019">
  <div class="backdrop-and-poster">
    <img class="poster" src="../films/posters/{{ slug }}.jpg" alt="" eleventy:ignore>
    <img class="backdrop" src="../films/backdrops/{{ slug }}.jpg" alt="" eleventy:ignore>
  </div>

  <h1>{{ film.title }} ({{ film | filmYear }})</h1>

  <p>
    {%- if film.language -%}Language: {{ film.language }}.{% endif %}
    Also known as 기생충.
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
  <img src="../films/profiles/20738.jpg" alt="Song Kang-ho" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Song Kang-ho</span>
    <span class="cast-card-character">Kim Ki-taek</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/115290.jpg" alt="Lee Sun-kyun" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Lee Sun-kyun</span>
    <span class="cast-card-character">Park Dong-ik</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/556435.jpg" alt="Cho Yeo-jeong" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Cho Yeo-jeong</span>
    <span class="cast-card-character">Yeon-kyo</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1255881.jpg" alt="Choi Woo-shik" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Choi Woo-shik</span>
    <span class="cast-card-character">Ki-woo</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1442583.jpg" alt="Park So-dam" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Park So-dam</span>
    <span class="cast-card-character">Ki-jung</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1572354.jpg" alt="Lee Jung-eun" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Lee Jung-eun</span>
    <span class="cast-card-character">Moon-gwang</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/2158882.jpg" alt="Jang Hye-jin" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Jang Hye-jin</span>
    <span class="cast-card-character">Chung-sook</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1694435.jpg" alt="Park Myung-hoon" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Park Myung-hoon</span>
    <span class="cast-card-character">Geun-se</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1418762.jpg" alt="Jung Ji-so" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Jung Ji-so</span>
    <span class="cast-card-character">Da-hye</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/2306987.jpg" alt="Jung Hyeon-jun" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Jung Hyeon-jun</span>
    <span class="cast-card-character">Da-song</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1709230.jpg" alt="Park Keun-rok" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Park Keun-rok</span>
    <span class="cast-card-character">Driver Yoon</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/2326809.jpg" alt="Jung Yi-seo" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Jung Yi-seo</span>
    <span class="cast-card-character">Pizza Manager</span>
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
    <li><a href="../rear-window-1954">Rear Window</a> because of Alfred Hitchcock</li>
<li><a href="../the-handmaiden-2016">The Handmaiden</a> because of Lee Ji-hye and Ahn Seong-bong</li>
  </ul>
</section>

</article>
