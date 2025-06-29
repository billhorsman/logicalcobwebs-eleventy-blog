---
title: "Rear Window"
layout: layouts/films.njk
slug: rear-window-1954
ogImage: content/bill/films/backdrops/rear-window-1954.jpg
description: "A wheelchair-bound photographer spies on his neighbors from his apartment window and becomes convinced one of them has committed murder."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../la-strada-1954"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">3 / 100</a>
  </div>
  <div class="next">
    <a href="../north-by-northwest-1959">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      La Strada
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      North by Northwest
    </span>
  </div>
</nav>

<article class="film slug-rear-window-1954">
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
    <li><a href="../its-a-wonderful-life-1946">It's a Wonderful Life</a> because of James Stewart</li>
<li><a href="../north-by-northwest-1959">North by Northwest</a> because of Jesslyn Fax, Bess Flowers, Len Hendry and Alfred Hitchcock</li>
<li><a href="../parasite-2019">Parasite</a> because of Alfred Hitchcock</li>
  </ul>
</section>

</article>
