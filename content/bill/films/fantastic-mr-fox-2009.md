---
title: "Fantastic Mr. Fox"
layout: layouts/films.njk
slug: fantastic-mr-fox-2009
ogImage: content/bill/films/backdrops/fantastic-mr-fox-2009.jpg
description: "The Fantastic Mr. Fox, bored with his current life, plans a heist against the three local farmers. The farmers, tired of sharing their chickens with the sly fox, seek revenge against him and his family."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../district-9-2009"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">53 / 100</a>
  </div>
  <div class="next">
    <a href="../micmacs-2009">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      District 9
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      Micmacs
    </span>
  </div>
</nav>

<article class="film slug-fantastic-mr-fox-2009">
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
    <li><a href="../the-deer-hunter-1978">The Deer Hunter</a> because of Meryl Streep</li>
<li><a href="../the-bourne-identity-2002">The Bourne Identity</a> because of Brian Cox</li>
<li><a href="../hot-fuzz-2007">Hot Fuzz</a> because of Garth Jennings</li>
<li><a href="../the-grand-budapest-hotel-2014">The Grand Budapest Hotel</a> because of Jason Schwartzman, Wallace Wolodarsky, Willem Dafoe, Bill Murray, Robin Hurlstone, Wes Anderson, Owen Wilson and Adrien Brody</li>
<li><a href="../the-french-dispatch-2021">The French Dispatch</a> because of Jason Schwartzman, Wallace Wolodarsky, Willem Dafoe, Bill Murray, Wes Anderson, Jarvis Cocker, Owen Wilson and Adrien Brody</li>
<li><a href="../asteroid-city-2023">Asteroid City</a> because of Jason Schwartzman, Willem Dafoe, Wes Anderson, Jarvis Cocker and Adrien Brody</li>
<li><a href="../the-lighthouse-2019">The Lighthouse</a> because of Willem Dafoe</li>
  </ul>
</section>

</article>
