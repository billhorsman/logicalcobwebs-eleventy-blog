---
title: "Bullitt"
layout: layouts/home.njk
slug: bullitt-1968
ogImage: content/bill/films/backdrops/bullitt-1968.jpg
description: "Senator Walter Chalmers is aiming to take down mob boss Pete Ross with the help of testimony from the criminal's hothead brother Johnny, who is in protective custody in San Francisco under the watch of police lieutenant Frank Bullitt. When a pair of mob hitmen enter the scene, Bullitt follows their trail through a maze of complications and double-crosses. This thriller includes one of the most famous car chases ever filmed."
---

{% set film = films[slug] %}

<nav class="films">
  <a class="prev" href="../2001-a-space-odyssey-1968">Previous</a>
  <a href="../">Film list</a>
  <a class="next" href="../once-upon-a-time-in-the-west-1968">Next</a>
</nav>

<p>12 / 100</p>

<article class="film">
  <div class="backdrop-and-poster">
    <img class="poster" src="../films/posters/{{ slug }}.jpg" alt="">
    <img class="backdrop" src="../films/backdrops/{{ slug }}.jpg" alt="">
  </div>

  <h1>{{ film.title }} ({{ film | filmYear }})</h1>

  

  <p class="director">
    Directed by <strong>{{ film | directors }}</strong>
  </p>

  {% if films.reviews[slug] %}
    <blockquote> 
      {{ films.reviews[slug] | safe }} <em>— Bill</em>
    </blockquote> 
  {% endif %}

  <h2>
    Cast
  </h2>
  <ul>
    {%- for cast in film.credits.cast -%}
      <li>
        {{ cast.name }} as <em>{{ cast.character }}</em>
      </li>
    {%- endfor -%}
  </ul>

  <h2>
    Crew
  </h2>
  <ul>
    {%- for crew in film.credits.crew -%}
      <li>
        {{ crew.name }} &mdash; <em>{{ crew.job }}</em>
      </li>
    {%- endfor -%}
  </ul>
</article>
<footer>
  <a href="../about">About this list</a>
</footer>
