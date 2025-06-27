---
title: "The Lighthouse"
layout: layouts/films.njk
slug: the-lighthouse-2019
ogImage: content/bill/films/backdrops/the-lighthouse-2019.jpg
description: "Two lighthouse keepers try to maintain their sanity while living on a remote and mysterious New England island in the 1890s."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../portrait-of-a-lady-on-fire-2019"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">73 / 100</a>
  </div>
  <div class="next">
    <a href="../eternal-beauty-2020">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      Portrait of a Lady on Fire
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      Eternal Beauty
    </span>
  </div>
</nav>

<article class="film slug-the-lighthouse-2019">
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
  <li><a href="../fantastic-mr-fox-2009">Fantastic Mr. Fox</a> (2009) by Willem Dafoe</li>
<li><a href="../the-grand-budapest-hotel-2014">The Grand Budapest Hotel</a> (2014) by Willem Dafoe</li>
<li><a href="../the-french-dispatch-2021">The French Dispatch</a> (2021) by Willem Dafoe</li>
<li><a href="../asteroid-city-2023">Asteroid City</a> (2023) by Willem Dafoe</li>
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
