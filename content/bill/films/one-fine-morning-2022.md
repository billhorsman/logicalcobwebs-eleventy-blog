---
title: "One Fine Morning"
layout: layouts/home.njk
slug: one-fine-morning-2022
ogImage: content/bill/films/backdrops/one-fine-morning-2022.jpg
description: "With a father suffering from neurodegenerative disease, a young woman lives with her eight-year-old daughter. While struggling to secure a decent nursing home, she runs into an unavailable friend with whom she embarks on an affair."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../empire-of-light-2022"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a href="../">Film list</a>
  </div>
  <div class="next">
    <a href="../the-banshees-of-inisherin-2022">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
</nav>

<p>94 / 100</p>

<article class="film slug-one-fine-morning-2022">
  <div class="backdrop-and-poster">
    <img class="poster" src="../films/posters/{{ slug }}.jpg" alt="">
    <img class="backdrop" src="../films/backdrops/{{ slug }}.jpg" alt="">
  </div>

  <h1>{{ film.title }} ({{ film | filmYear }})</h1>

  <p>
    {%- if film.language -%}Language: {{ film.language }}.{% endif %}
    Also known as Un beau matin.
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
