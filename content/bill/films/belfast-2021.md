---
title: "Belfast"
layout: layouts/films.njk
slug: belfast-2021
ogImage: content/bill/films/backdrops/belfast-2021.jpg
description: "Buddy is a young boy on the cusp of adolescence, whose life is filled with familial love, childhood hijinks, and a blossoming romance. Yet, with his beloved hometown caught up in increasing turmoil, his family faces a momentous choice: hope the conflict will pass or leave everything they know behind for a new life."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../the-truffle-hunters-2020"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">76 / 100</a>
  </div>
  <div class="next">
    <a href="../coda-2021">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      The Truffle Hunters
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      CODA
    </span>
  </div>
</nav>

<article class="film slug-belfast-2021">
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
  <img src="../films/profiles/3232669.jpg" alt="Jude Hill" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Jude Hill</span>
    <span class="cast-card-character">Buddy</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1254583.jpg" alt="Jamie Dornan" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Jamie Dornan</span>
    <span class="cast-card-character">Pa</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/147056.jpg" alt="Caitríona Balfe" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Caitríona Balfe</span>
    <span class="cast-card-character">Ma</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/3065081.jpg" alt="Lewis McAskie" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Lewis McAskie</span>
    <span class="cast-card-character">Will</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/5309.jpg" alt="Judi Dench" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Judi Dench</span>
    <span class="cast-card-character">Granny</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/8785.jpg" alt="Ciarán Hinds" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Ciarán Hinds</span>
    <span class="cast-card-character">Pop</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1726976.jpg" alt="Lara McDonnell" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Lara McDonnell</span>
    <span class="cast-card-character">Moira</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/228866.jpg" alt="Colin Morgan" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Colin Morgan</span>
    <span class="cast-card-character">Billy Clanton</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/56101.jpg" alt="Gerard Horan" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Gerard Horan</span>
    <span class="cast-card-character">Mackie</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1709739.jpg" alt="Josie Walker" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Josie Walker</span>
    <span class="cast-card-character">Auntie Violet</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/2893134.jpg" alt="Olive Tennant" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Olive Tennant</span>
    <span class="cast-card-character">Catherine</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/17483.jpg" alt="Michael Maloney" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Michael Maloney</span>
    <span class="cast-card-character">Frankie West</span>
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
    <li><a href="../in-bruges-2008">In Bruges</a> because of Ciarán Hinds</li>
  </ul>
</section>

</article>
