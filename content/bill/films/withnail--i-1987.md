---
title: "Withnail & I"
layout: layouts/home.njk
slug: withnail--i-1987
ogImage: content/bill/films/backdrops/withnail--i-1987.jpg
description: "Two out-of-work actors -- the anxious, luckless Marwood and his acerbic, alcoholic friend, Withnail -- spend their days drifting between their squalid flat, the unemployment office and the pub. When they take a holiday \"by mistake\" at the country house of Withnail's flamboyantly gay uncle, Monty, they encounter the unpleasant side of the English countryside: tedium, terrifying locals and torrential rain."
---

{% set film = films[slug] %}

<nav class="films">
  <a class="prev" href="../brazil-1985">Previous</a>
  <a href="../">Film list</a>
  <a class="next" href="../delicatessen-1991">Next</a>
</nav>

<p>29 / 100</p>

<article class="film slug-withnail--i-1987">
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
