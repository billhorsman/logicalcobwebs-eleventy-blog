---
title: "Dallas Buyers Club"
layout: layouts/films.njk
slug: dallas-buyers-club-2013
ogImage: content/bill/films/backdrops/dallas-buyers-club-2013.jpg
description: "Loosely based on the true-life tale of Ron Woodroof, a drug-taking, women-loving, homophobic man who in 1986 was diagnosed with HIV/AIDS and given thirty days to live."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../all-is-lost-2013"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">57 / 100</a>
  </div>
  <div class="next">
    <a href="../mr-turner-2014">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      All Is Lost
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      Mr. Turner
    </span>
  </div>
</nav>

<article class="film slug-dallas-buyers-club-2013">
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
    <li><a href="../fight-club-1999">Fight Club</a>, <a href="../phone-booth-2003">Phone Booth</a> and <a href="../house-of-gucci-2021">House of Gucci</a> because of Jared Leto</li>
<li><a href="../the-french-dispatch-2021">The French Dispatch</a> because of Griffin Dunne</li>
<li><a href="../killers-of-the-flower-moon-2023">Killers of the Flower Moon</a> because of Carl Palmer</li>
  </ul>
</section>

</article>
