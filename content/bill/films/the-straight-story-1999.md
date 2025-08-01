---
title: "The Straight Story"
layout: layouts/films.njk
slug: the-straight-story-1999
ogImage: content/bill/films/backdrops/the-straight-story-1999.jpg
description: "A retired farmer and widower in his 70s, Alvin Straight learns one day that his distant brother Lyle has suffered a stroke and may not recover. Alvin is determined to make things right with Lyle while he still can, but his brother lives in Wisconsin, while Alvin is stuck in Iowa with no car and no driver's license. Then he hits on the idea of making the trip on his old lawnmower, thus beginning a picturesque and at times deeply spiritual odyssey."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../magnolia-1999"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">37 / 100</a>
  </div>
  <div class="next">
    <a href="../the-talented-mr-ripley-1999">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      Magnolia
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      The Talented Mr. Ripley
    </span>
  </div>
</nav>

<article class="film slug-the-straight-story-1999">
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
    <li><a href="../paris-texas-1984">Paris, Texas</a> because of Harry Dean Stanton</li>
<li><a href="../lucky-2017">Lucky</a> because of Harry Dean Stanton and David Lynch</li>
<li><a href="../fargo-1996">Fargo</a> because of Sally Wingert</li>
<li><a href="../the-fabelmans-2022">The Fabelmans</a> because of David Lynch</li>
  </ul>
</section>

</article>
