---
title: "Kill Bill: The Whole Bloody Affair"
layout: layouts/films.njk
slug: kill-bill-the-whole-bloody-affair-2011
ogImage: content/bill/films/backdrops/kill-bill-the-whole-bloody-affair-2011.jpg
description: "An assassin is shot and almost killed by her ruthless employer, Bill, and other members of their assassination circle – but she lives to plot her vengeance."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../tomboy-2011"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">52 / 100</a>
  </div>
  <div class="next">
    <a href="../all-is-lost-2013">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      Tomboy
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      All Is Lost
    </span>
  </div>
</nav>

<article class="film slug-kill-bill-the-whole-bloody-affair-2011">
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
  <img src="../films/profiles/139.jpg" alt="Uma Thurman" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Uma Thurman</span>
    <span class="cast-card-character">The Bride / Beatrix Kiddo (Black Mamba) / Mommy</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/140.jpg" alt="Lucy Liu" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Lucy Liu</span>
    <span class="cast-card-character">O-Ren Ishii (Cottonmouth)</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/2535.jpg" alt="Vivica A. Fox" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Vivica A. Fox</span>
    <span class="cast-card-character">Vernita Green (Copperhead)</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/147.jpg" alt="Michael Madsen" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Michael Madsen</span>
    <span class="cast-card-character">Budd (Sidewinder)</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/589.jpg" alt="Daryl Hannah" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Daryl Hannah</span>
    <span class="cast-card-character">Elle Driver (California Mountain Snake)</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/141.jpg" alt="David Carradine" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">David Carradine</span>
    <span class="cast-card-character">Bill (Snake Charmer)</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/2537.jpg" alt="Sonny Chiba" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Sonny Chiba</span>
    <span class="cast-card-character">Hattori Hanzo</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/2539.jpg" alt="Julie Dreyfus" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Julie Dreyfus</span>
    <span class="cast-card-character">Sofie Fatale</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/2538.jpg" alt="Chiaki Kuriyama" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Chiaki Kuriyama</span>
    <span class="cast-card-character">Gogo Yubari</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/240171.jpg" alt="Gordon Liu Chia-Hui" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Gordon Liu Chia-Hui</span>
    <span class="cast-card-character">Johnny Mo / Pai Mei</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/2536.jpg" alt="Michael Parks" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Michael Parks</span>
    <span class="cast-card-character">Earl McGraw / Esteban Vihaio</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/4026.jpg" alt="Ambrosia Kelley" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Ambrosia Kelley</span>
    <span class="cast-card-character">Nikki Bell</span>
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
    <li><a href="../blade-runner-1982">Blade Runner</a> because of Daryl Hannah</li>
<li><a href="../magnolia-1999">Magnolia</a> because of Michael Bowen</li>
<li><a href="../phone-booth-2003">Phone Booth</a> because of Shu Lan Tuan</li>
  </ul>
</section>

</article>
