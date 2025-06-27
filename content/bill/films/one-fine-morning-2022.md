---
title: "One Fine Morning"
layout: layouts/films.njk
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
    <a class="simple" href="../">93 / 100</a>
  </div>
  <div class="next">
    <a href="../the-banshees-of-inisherin-2022">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      Empire of Light
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      The Banshees of Inisherin
    </span>
  </div>
</nav>

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

  <p class="related-films">Related to:</p>
  <ul class="related-films">
  <li><a href="../the-grand-budapest-hotel-2014">The Grand Budapest Hotel</a> and <a href="../the-french-dispatch-2021">The French Dispatch</a> by Léa Seydoux</li>
<li><a href="../the-french-dispatch-2021">The French Dispatch</a> by Sharif Andoura</li>
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
