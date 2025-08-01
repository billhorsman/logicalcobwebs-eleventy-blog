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
    <a class="simple" href="../">21 / 100</a>
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
