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
    <a class="simple" href="../">12 / 100</a>
  </div>
  <div class="next">
    <a href="../once-upon-a-time-in-the-west-1968">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
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
