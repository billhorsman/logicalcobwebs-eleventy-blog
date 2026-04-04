---
title: "The Talented Mr. Ripley"
layout: layouts/films.njk
slug: the-talented-mr-ripley-1999
ogImage: content/bill/films/backdrops/the-talented-mr-ripley-1999.jpg
description: "Tom Ripley is a calculating young man who believes it's better to be a fake somebody than a real nobody. Opportunity knocks in the form of a wealthy U.S. shipbuilder who hires Tom to travel to Italy to bring back his playboy son, Dickie. Ripley worms his way into the idyllic lives of Dickie and his girlfriend, plunging into a daring scheme of duplicity, lies and murder."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../the-straight-story-1999"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">36 / 100</a>
  </div>
  <div class="next">
    <a href="../billy-elliot-2000">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      The Straight Story
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      Billy Elliot
    </span>
  </div>
</nav>

<article class="film slug-the-talented-mr-ripley-1999">
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
    <span class="cast-card-character">Tom Ripley</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/12052.jpg" alt="Gwyneth Paltrow" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Gwyneth Paltrow</span>
    <span class="cast-card-character">Marge Sherwood</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/9642.jpg" alt="Jude Law" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Jude Law</span>
    <span class="cast-card-character">Dickie Greenleaf</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/112.jpg" alt="Cate Blanchett" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Cate Blanchett</span>
    <span class="cast-card-character">Meredith Logue</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1233.jpg" alt="Philip Seymour Hoffman" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Philip Seymour Hoffman</span>
    <span class="cast-card-character">Freddie Miles</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1709.jpg" alt="Jack Davenport" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Jack Davenport</span>
    <span class="cast-card-character">Peter Smith-Kingsley</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/8986.jpg" alt="James Rebhorn" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">James Rebhorn</span>
    <span class="cast-card-character">Herbert Greenleaf</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/15915.jpg" alt="Sergio Rubini" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Sergio Rubini</span>
    <span class="cast-card-character">Inspector Roverini</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/4492.jpg" alt="Philip Baker Hall" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Philip Baker Hall</span>
    <span class="cast-card-character">Alvin MacCarron</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1989.jpg" alt="Celia Weston" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Celia Weston</span>
    <span class="cast-card-character">Aunt Joan</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/15916.jpg" alt="Fiorello" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Fiorello</span>
    <span class="cast-card-character">Fausto</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/15917.jpg" alt="Stefania Rocca" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Stefania Rocca</span>
    <span class="cast-card-character">Silvana</span>
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
    <li><a href="../good-will-hunting-1997">Good Will Hunting</a> and <a href="../the-bourne-identity-2002">The Bourne Identity</a> because of Matt Damon</li>
<li><a href="../the-big-lebowski-1998">The Big Lebowski</a> because of Philip Seymour Hoffman</li>
<li><a href="../magnolia-1999">Magnolia</a> because of Philip Seymour Hoffman and Philip Baker Hall</li>
<li><a href="../the-grand-budapest-hotel-2014">The Grand Budapest Hotel</a> because of Jude Law</li>
<li><a href="../hot-fuzz-2007">Hot Fuzz</a> because of Cate Blanchett</li>
<li><a href="../house-of-gucci-2021">House of Gucci</a> because of Pietro Ragusa</li>
  </ul>
</section>

</article>
