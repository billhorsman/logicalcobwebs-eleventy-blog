---
title: "No Country for Old Men"
layout: layouts/films.njk
slug: no-country-for-old-men-2007
ogImage: content/bill/films/backdrops/no-country-for-old-men-2007.jpg
description: "Llewelyn Moss stumbles upon dead bodies, $2 million and a hoard of heroin in a Texas desert, but methodical killer Anton Chigurh comes looking for it, with local sheriff Ed Tom Bell hot on his trail. The roles of prey and predator blur as the violent pursuit of money and justice collide."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../hot-fuzz-2007"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">49 / 100</a>
  </div>
  <div class="next">
    <a href="../happygolucky-2008">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      Hot Fuzz
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      Happy-Go-Lucky
    </span>
  </div>
</nav>

<article class="film slug-no-country-for-old-men-2007">
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
    <li><a href="../fargo-1996">Fargo</a> and <a href="../the-big-lebowski-1998">The Big Lebowski</a> because of Joel Coen</li>
<li><a href="../the-tragedy-of-macbeth-2021">The Tragedy of Macbeth</a> because of Joel Coen and Stephen Root</li>
<li><a href="../trainspotting-1996">Trainspotting</a> because of Kelly Macdonald</li>
<li><a href="../dune-2021">Dune</a> because of Javier Bardem and Josh Brolin</li>
<li><a href="../killers-of-the-flower-moon-2023">Killers of the Flower Moon</a> because of Barry Corbin and Gene Jones</li>
<li><a href="../lucky-2017">Lucky</a> because of Beth Grant</li>
  </ul>
</section>

</article>
