---
title: "In the Heat of the Night"
layout: layouts/films.njk
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
    <a class="simple" href="../">9 / 100</a>
  </div>
  <div class="next">
    <a href="../2001-a-space-odyssey-1968">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      Barefoot in the Park
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      2001: A Space Odyssey
    </span>
  </div>
</nav>

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

  <p class="related-films">Related to:</p>
  <ul class="related-films">
  <li><a href="../the-sting-1973">The Sting</a> (1973) by Larry D. Mann</li>
<li><a href="../butch-cassidy-and-the-sundance-kid-1969">Butch Cassidy and the Sundance Kid</a> (1969) by Timothy Scott</li>
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
