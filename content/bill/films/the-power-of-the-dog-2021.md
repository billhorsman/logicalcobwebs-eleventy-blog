---
title: "The Power of the Dog"
layout: layouts/films.njk
slug: the-power-of-the-dog-2021
ogImage: content/bill/films/backdrops/the-power-of-the-dog-2021.jpg
description: "A domineering but charismatic rancher wages a war of intimidation on his brother's new wife and her teen son, until long-hidden secrets come to light."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../the-french-dispatch-2021"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">87 / 100</a>
  </div>
  <div class="next">
    <a href="../the-tragedy-of-macbeth-2021">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      The French Dispatch
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      The Tragedy of Macbeth
    </span>
  </div>
</nav>

<article class="film slug-the-power-of-the-dog-2021">
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
    <li><a href="../1917-2019">1917</a> because of Benedict Cumberbatch</li>
<li><a href="../killers-of-the-flower-moon-2023">Killers of the Flower Moon</a> because of Jesse Plemons</li>
  </ul>
</section>

</article>
