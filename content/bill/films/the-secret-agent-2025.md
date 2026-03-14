---
title: "The Secret Agent"
layout: layouts/films.njk
slug: the-secret-agent-2025
ogImage: content/bill/films/backdrops/the-secret-agent-2025.jpg
description: "Brazil, 1977. Marcelo, a technology expert in his early 40s, is on the run. Hoping to reunite with his son, he travels to Recife during Carnival but soon realizes that the city is not the safe haven he was expecting."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../nouvelle-vague-2025"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">100 / 100</a>
  </div>
  <div class="next">
    <span>Next <i class="fa-solid fa-chevron-right fa-xs"></i></span>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      Nouvelle Vague
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      End of list
    </span>
  </div>
</nav>

<article class="film slug-the-secret-agent-2025">
  <div class="backdrop-and-poster">
    <img class="poster" src="../films/posters/{{ slug }}.jpg" alt="">
    <img class="backdrop" src="../films/backdrops/{{ slug }}.jpg" alt="">
  </div>

  <h1>{{ film.title }} ({{ film | filmYear }})</h1>

  <p>
    {%- if film.language -%}Language: {{ film.language }}.{% endif %}
    Also known as O Agente Secreto.
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
