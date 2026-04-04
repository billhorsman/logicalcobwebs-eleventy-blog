---
title: "Blade Runner"
layout: layouts/films.njk
slug: blade-runner-1982
ogImage: content/bill/films/backdrops/blade-runner-1982.jpg
description: "In the smog-choked dystopian Los Angeles of 2019, blade runner Rick Deckard is called out of retirement to terminate a quartet of replicants who have escaped to Earth seeking their creator for a way to extend their short life spans."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../diva-1981"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">19 / 100</a>
  </div>
  <div class="next">
    <a href="../local-hero-1983">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      Diva
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      Local Hero
    </span>
  </div>
</nav>

<article class="film slug-blade-runner-1982">
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
  <img src="../films/profiles/3.jpg" alt="Harrison Ford" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Harrison Ford</span>
    <span class="cast-card-character">Deckard</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/585.jpg" alt="Rutger Hauer" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Rutger Hauer</span>
    <span class="cast-card-character">Batty</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/586.jpg" alt="Sean Young" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Sean Young</span>
    <span class="cast-card-character">Rachael</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/587.jpg" alt="Edward James Olmos" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Edward James Olmos</span>
    <span class="cast-card-character">Gaff</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/588.jpg" alt="M. Emmet Walsh" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">M. Emmet Walsh</span>
    <span class="cast-card-character">Bryant</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/589.jpg" alt="Daryl Hannah" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Daryl Hannah</span>
    <span class="cast-card-character">Pris</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/590.jpg" alt="William Sanderson" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">William Sanderson</span>
    <span class="cast-card-character">Sebastian</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/591.jpg" alt="Brion James" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Brion James</span>
    <span class="cast-card-character">Leon</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/592.jpg" alt="Joe Turkel" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Joe Turkel</span>
    <span class="cast-card-character">Tyrell</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/593.jpg" alt="Joanna Cassidy" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Joanna Cassidy</span>
    <span class="cast-card-character">Zhora</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/20904.jpg" alt="James Hong" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">James Hong</span>
    <span class="cast-card-character">Chew</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/58495.jpg" alt="Morgan Paull" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Morgan Paull</span>
    <span class="cast-card-character">Holden</span>
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
    <li><a href="../apocalypse-now-1979">Apocalypse Now</a> because of Harrison Ford</li>
<li><a href="../kill-bill-the-whole-bloody-affair-2011">Kill Bill: The Whole Bloody Affair</a> because of Daryl Hannah</li>
<li><a href="../black-hawk-down-2001">Black Hawk Down</a> and <a href="../house-of-gucci-2021">House of Gucci</a> because of Ridley Scott</li>
  </ul>
</section>

</article>
