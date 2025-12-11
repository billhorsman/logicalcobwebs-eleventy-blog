---
title: "Nouvelle Vague"
layout: layouts/films.njk
slug: nouvelle-vague-2025
ogImage: content/bill/films/backdrops/nouvelle-vague-2025.jpg
description: "After writing for Cahiers du cinéma, a young Jean-Luc Godard decides making films is the best film criticism. He convinces producer Georges de Beauregard to fund a low-budget feature, and creates a treatment with fellow New Wave filmmaker François Truffaut about a gangster couple. The result? Breathless, one of the first features of the Nouvelle Vague era of French cinema."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../the-ballad-of-wallis-island-2025"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
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
      The Ballad of Wallis Island
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      End of list
    </span>
  </div>
</nav>

<article class="film slug-nouvelle-vague-2025">
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

  <section class="related-films">
  <h2>Related films</h2>
  <ul>
    <li><a href="../anatomy-of-a-fall-2023">Anatomy of a Fall</a> because of Pierre-François Garel</li>
  </ul>
</section>

</article>
