---
title: "The Bourne Identity"
layout: layouts/films.njk
slug: the-bourne-identity-2002
ogImage: content/bill/films/backdrops/the-bourne-identity-2002.jpg
description: "Wounded to the brink of death and suffering from amnesia, Jason Bourne is rescued at sea by a fisherman. With nothing to go on but a Swiss bank account number, he starts to reconstruct his life, but finds that many people he encounters want him dead. However, Bourne realizes that he has the combat and mental skills of a world-class spy—but who does he work for?"
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../man-on-the-train-2002"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">42 / 100</a>
  </div>
  <div class="next">
    <a href="../phone-booth-2003">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      Man on the Train
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      Phone Booth
    </span>
  </div>
</nav>

<article class="film slug-the-bourne-identity-2002">
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
  <img src="../films/profiles/1892.jpg" alt="Matt Damon" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Matt Damon</span>
    <span class="cast-card-character">Jason Bourne</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/679.jpg" alt="Franka Potente" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Franka Potente</span>
    <span class="cast-card-character">Marie Helena Kreutz</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/2955.jpg" alt="Chris Cooper" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Chris Cooper</span>
    <span class="cast-card-character">Alexander Conklin</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/2296.jpg" alt="Clive Owen" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Clive Owen</span>
    <span class="cast-card-character">The Professor</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1248.jpg" alt="Brian Cox" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Brian Cox</span>
    <span class="cast-card-character">Ward Abbott</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/31164.jpg" alt="Adewale Akinnuoye-Agbaje" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Adewale Akinnuoye-Agbaje</span>
    <span class="cast-card-character">Nykwana Wombosi</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/32458.jpg" alt="Gabriel Mann" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Gabriel Mann</span>
    <span class="cast-card-character">Danny Zorn</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/12041.jpg" alt="Julia Stiles" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Julia Stiles</span>
    <span class="cast-card-character">Nicky Parsons</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/27740.jpg" alt="Walton Goggins" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Walton Goggins</span>
    <span class="cast-card-character">Research Tech</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/52419.jpg" alt="Josh Hamilton" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Josh Hamilton</span>
    <span class="cast-card-character">Research Tech</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/27198.jpg" alt="Orso Maria Guerrini" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Orso Maria Guerrini</span>
    <span class="cast-card-character">Giancarlo</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/49965.jpg" alt="Tim Dutton" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Tim Dutton</span>
    <span class="cast-card-character">Eamon</span>
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
    <li><a href="../good-will-hunting-1997">Good Will Hunting</a> and <a href="../the-talented-mr-ripley-1999">The Talented Mr. Ripley</a> because of Matt Damon</li>
<li><a href="../little-women-2019">Little Women</a> because of Chris Cooper</li>
<li><a href="../fantastic-mr-fox-2009">Fantastic Mr. Fox</a> because of Brian Cox</li>
<li><a href="../mr-turner-2014">Mr. Turner</a> because of Vincent Franklin</li>
  </ul>
</section>

</article>
