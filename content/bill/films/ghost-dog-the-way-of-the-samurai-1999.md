---
title: "Ghost Dog: The Way of the Samurai"
layout: layouts/films.njk
slug: ghost-dog-the-way-of-the-samurai-1999
ogImage: content/bill/films/backdrops/ghost-dog-the-way-of-the-samurai-1999.jpg
description: "An African-American Mafia hit man who models himself after the samurai of ancient Japan finds himself targeted for death by the mob."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../fight-club-1999"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">33 / 100</a>
  </div>
  <div class="next">
    <a href="../magnolia-1999">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      Fight Club
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      Magnolia
    </span>
  </div>
</nav>

<article class="film slug-ghost-dog-the-way-of-the-samurai-1999">
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
  <img src="../films/profiles/2178.jpg" alt="Forest Whitaker" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Forest Whitaker</span>
    <span class="cast-card-character">Ghost Dog</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/28002.jpg" alt="John Tormey" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">John Tormey</span>
    <span class="cast-card-character">Louie</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/19183.jpg" alt="Cliff Gorman" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Cliff Gorman</span>
    <span class="cast-card-character">Sonny Valerio</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/39598.jpg" alt="Frank Minucci" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Frank Minucci</span>
    <span class="cast-card-character">Big Angie</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/4255.jpg" alt="Richard Portnow" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Richard Portnow</span>
    <span class="cast-card-character">Handsome Frank</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/39596.jpg" alt="Tricia Vessey" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Tricia Vessey</span>
    <span class="cast-card-character">Louise Vargo</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/14731.jpg" alt="Henry Silva" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Henry Silva</span>
    <span class="cast-card-character">Ray Vargo</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/39599.jpg" alt="Gene Ruffini" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Gene Ruffini</span>
    <span class="cast-card-character">Old Consigliere</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/4692.jpg" alt="Frank Adonis" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Frank Adonis</span>
    <span class="cast-card-character">Valerio's Bodyguard</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/2561.jpg" alt="Victor Argo" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Victor Argo</span>
    <span class="cast-card-character">Vinny</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/4812.jpg" alt="Isaach de Bankolé" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Isaach de Bankolé</span>
    <span class="cast-card-character">Raymond</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1222948.jpg" alt="Camille Winbush" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Camille Winbush</span>
    <span class="cast-card-character">Pearline</span>
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
    <li><a href="../night-on-earth-1991">Night on Earth</a> because of Isaach de Bankolé and Jim Jarmusch</li>
<li><a href="../phone-booth-2003">Phone Booth</a> because of Forest Whitaker</li>
  </ul>
</section>

</article>
