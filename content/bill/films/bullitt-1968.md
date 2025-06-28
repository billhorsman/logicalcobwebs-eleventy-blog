---
title: "Bullitt"
layout: layouts/films.njk
slug: bullitt-1968
ogImage: content/bill/films/backdrops/bullitt-1968.jpg
description: "Senator Walter Chalmers is aiming to take down mob boss Pete Ross with the help of testimony from the criminal's hothead brother Johnny, who is in protective custody in San Francisco under the watch of police lieutenant Frank Bullitt. When a pair of mob hitmen enter the scene, Bullitt follows their trail through a maze of complications and double-crosses. This thriller includes one of the most famous car chases ever filmed."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../2001-a-space-odyssey-1968"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">11 / 100</a>
  </div>
  <div class="next">
    <a href="../once-upon-a-time-in-the-west-1968">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      2001: A Space Odyssey
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      Once Upon a Time in the West
    </span>
  </div>
</nav>

<article class="film slug-bullitt-1968">
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

  <p class="related-films">Related to:</p>
  <ul class="related-films">
  <li><a href="../north-by-northwest-1959">North by Northwest</a> by Paul Genge</li>
<li><a href="../apocalypse-now-1979">Apocalypse Now</a> by Robert Duvall</li>
  </ul>

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
