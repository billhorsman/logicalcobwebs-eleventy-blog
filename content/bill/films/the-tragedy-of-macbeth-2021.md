---
title: "The Tragedy of Macbeth"
layout: layouts/films.njk
slug: the-tragedy-of-macbeth-2021
ogImage: content/bill/films/backdrops/the-tragedy-of-macbeth-2021.jpg
description: "Macbeth, the Thane of Glamis, receives a prophecy from a trio of witches that one day he will become King of Scotland. Consumed by ambition and spurred to action by his wife, Macbeth murders his king and takes the throne for himself."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../the-power-of-the-dog-2021"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">86 / 100</a>
  </div>
  <div class="next">
    <a href="../between-two-worlds-2022">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      The Power of the Dog
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      Between Two Worlds
    </span>
  </div>
</nav>

<article class="film slug-the-tragedy-of-macbeth-2021">
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
  <img src="../films/profiles/5292.jpg" alt="Denzel Washington" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Denzel Washington</span>
    <span class="cast-card-character">Macbeth</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/3910.jpg" alt="Frances McDormand" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Frances McDormand</span>
    <span class="cast-card-character">Lady Macbeth</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/178634.jpg" alt="Alex Hassell" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Alex Hassell</span>
    <span class="cast-card-character">Ross</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/216425.jpg" alt="Bertie Carvel" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Bertie Carvel</span>
    <span class="cast-card-character">Banquo</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/2039.jpg" alt="Brendan Gleeson" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Brendan Gleeson</span>
    <span class="cast-card-character">Duncan</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1154054.jpg" alt="Corey Hawkins" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Corey Hawkins</span>
    <span class="cast-card-character">Macduff</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/10982.jpg" alt="Harry Melling" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Harry Melling</span>
    <span class="cast-card-character">Malcolm</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/65885.jpg" alt="Miles Anderson" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Miles Anderson</span>
    <span class="cast-card-character">Lennox</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/72309.jpg" alt="Kathryn Hunter" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Kathryn Hunter</span>
    <span class="cast-card-character">Witches / Old Man</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/2684944.jpg" alt="Matt Helm" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Matt Helm</span>
    <span class="cast-card-character">Donalbain</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/2042908.jpg" alt="Moses Ingram" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Moses Ingram</span>
    <span class="cast-card-character">Lady Macduff</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/60062.jpg" alt="Scott Subiono" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Scott Subiono</span>
    <span class="cast-card-character">Murderer</span>
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
    <li><a href="../fargo-1996">Fargo</a> because of Frances McDormand and Joel Coen</li>
<li><a href="../nomadland-2021">Nomadland</a> and <a href="../the-french-dispatch-2021">The French Dispatch</a> because of Frances McDormand</li>
<li><a href="../the-big-lebowski-1998">The Big Lebowski</a> because of Joel Coen</li>
<li><a href="../no-country-for-old-men-2007">No Country for Old Men</a> because of Joel Coen and Stephen Root</li>
<li><a href="../in-bruges-2008">In Bruges</a> and <a href="../the-banshees-of-inisherin-2022">The Banshees of Inisherin</a> because of Brendan Gleeson</li>
  </ul>
</section>

</article>
