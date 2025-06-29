---
title: "It's a Wonderful Life"
layout: layouts/films.njk
slug: its-a-wonderful-life-1946
ogImage: content/bill/films/backdrops/its-a-wonderful-life-1946.jpg
description: "A holiday favourite for generations...  George Bailey has spent his entire life giving to the people of Bedford Falls.  All that prevents rich skinflint Mr. Potter from taking over the entire town is George's modest building and loan company.  But on Christmas Eve the business's $8,000 is lost and George's troubles begin."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <span><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</span>
  </div>
  <div>
    <a class="simple" href="../">1 / 100</a>
  </div>
  <div class="next">
    <a href="../whisky-galore-1949">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      Start of list
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      Whisky Galore!
    </span>
  </div>
</nav>

<article class="film slug-its-a-wonderful-life-1946">
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
    <li><a href="../barefoot-in-the-park-1967">Barefoot in the Park</a> because of John Indrisano</li>
<li><a href="../north-by-northwest-1959">North by Northwest</a> because of Finn Zirzow</li>
  </ul>
</section>

</article>
