---
title: "24 Hour Party People"
layout: layouts/films.njk
slug: 24-hour-party-people-2002
ogImage: content/bill/films/backdrops/24-hour-party-people-2002.jpg
description: "Manchester, 1976. Tony Wilson is an ambitious but frustrated local TV news reporter looking for a way to make his mark. After witnessing a life-changing concert by a band known as the Sex Pistols, he persuades his station to televise one of their performances, and soon Manchester's punk groups are clamoring for him to manage them. Riding the wave of a musical revolution, Wilson and his friends create the legendary Factory Records label and The Hacienda club."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../black-hawk-down-2001"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">40 / 100</a>
  </div>
  <div class="next">
    <a href="../man-on-the-train-2002">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      Black Hawk Down
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      Man on the Train
    </span>
  </div>
</nav>

<article class="film slug-24-hour-party-people-2002">
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
  <img src="../films/profiles/4581.jpg" alt="Steve Coogan" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Steve Coogan</span>
    <span class="cast-card-character">Tony Wilson</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/14887.jpg" alt="Paddy Considine" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Paddy Considine</span>
    <span class="cast-card-character">Rob Gretton</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/16702.jpg" alt="Sean Harris" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Sean Harris</span>
    <span class="cast-card-character">Ian Curtis</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1120.jpg" alt="Lennie James" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Lennie James</span>
    <span class="cast-card-character">Alan Erasmus</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1834.jpg" alt="Shirley Henderson" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Shirley Henderson</span>
    <span class="cast-card-character">Lindsay Wilson</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1333.jpg" alt="Andy Serkis" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Andy Serkis</span>
    <span class="cast-card-character">Martin Hannett</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/42604.jpg" alt="John Simm" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">John Simm</span>
    <span class="cast-card-character">Bernard Sumner</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/96280.jpg" alt="Ralf Little" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Ralf Little</span>
    <span class="cast-card-character">Hooky</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1667.jpg" alt="Danny Cunningham" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Danny Cunningham</span>
    <span class="cast-card-character">Shaun Ryder</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/7318.jpg" alt="Peter Kay" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Peter Kay</span>
    <span class="cast-card-character">Don Tonay</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/7321.jpg" alt="John Thomson" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">John Thomson</span>
    <span class="cast-card-character">Charles</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/71282.jpg" alt="Kate Magowan" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Kate Magowan</span>
    <span class="cast-card-character">Yvette Livesay</span>
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
    <li><a href="../shallow-grave-1994">Shallow Grave</a> because of Christopher Eccleston and Keith Allen</li>
<li><a href="../trainspotting-1996">Trainspotting</a> because of Keith Allen and Shirley Henderson</li>
<li><a href="../hot-fuzz-2007">Hot Fuzz</a> because of Steve Coogan, Paddy Considine, Simon Pegg and Ron Cook</li>
<li><a href="../empire-of-light-2022">Empire of Light</a> because of Ron Cook</li>
<li><a href="../dune-2021">Dune</a> because of Neil Bell</li>
  </ul>
</section>

</article>
