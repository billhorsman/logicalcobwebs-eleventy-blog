---
title: "Three Days of the Condor"
layout: layouts/films.njk
slug: three-days-of-the-condor-1975
ogImage: content/bill/films/backdrops/three-days-of-the-condor-1975.jpg
description: "When bookish CIA researcher Joe Turner finds all his co-workers dead, he, together with a woman he has kidnapped, must work together to outwit those responsible until he determines who he can really trust."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../dog-day-afternoon-1975"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">17 / 100</a>
  </div>
  <div class="next">
    <a href="../the-man-who-fell-to-earth-1976">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      Dog Day Afternoon
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      The Man Who Fell to Earth
    </span>
  </div>
</nav>

<article class="film slug-three-days-of-the-condor-1975">
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
</article>
