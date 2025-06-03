---
title: "In the Heat of the Night"
layout: layouts/home.njk
slug: in-the-heat-of-the-night-1967
ogImage: content/bill/films/backdrops/in-the-heat-of-the-night-1967.jpg
description: "African-American Philadelphia police detective Virgil Tibbs is arrested on suspicion of murder by Bill Gillespie, the racist police chief of tiny Sparta, Mississippi. After Tibbs proves not only his own innocence but that of another man, he joins forces with Gillespie to track down the real killer. Their investigation takes them through every social level of the town, with Tibbs making enemies as well as unlikely friends as he hunts for the truth."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../barefoot-in-the-park-1967"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a href="../">Film list</a>
  </div>
  <div class="next">
    <a href="../2001-a-space-odyssey-1968">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
</nav>

<p>10 / 100</p>

<article class="film slug-in-the-heat-of-the-night-1967">
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
