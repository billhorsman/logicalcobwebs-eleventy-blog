---
title: "The Fugitive"
layout: layouts/films.njk
slug: the-fugitive-1993
ogImage: content/bill/films/backdrops/the-fugitive-1993.jpg
description: "Wrongfully convicted of murdering his wife and sentenced to death, Richard Kimble escapes from the law in an attempt to find the real killer and clear his name."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../night-on-earth-1991"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">31 / 100</a>
  </div>
  <div class="next">
    <a href="../whats-eating-gilbert-grape-1993">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      Night on Earth
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      What's Eating Gilbert Grape
    </span>
  </div>
</nav>

<article class="film slug-the-fugitive-1993">
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

  <details>
    <summary>
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

</article>
<footer>
  <a href="../about">About this list</a>
</footer>
