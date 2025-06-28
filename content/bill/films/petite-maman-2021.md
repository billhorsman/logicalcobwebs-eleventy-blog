---
title: "Petite Maman"
layout: layouts/films.njk
slug: petite-maman-2021
ogImage: content/bill/films/backdrops/petite-maman-2021.jpg
description: "After the death of her beloved grandmother, eight-year-old Nelly meets a strangely familiar girl her own age in the woods. Instantly forming a connection with this mysterious new friend, Nelly embarks on a fantastical journey of discovery which helps her come to terms with this newfound loss."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../nomadland-2021"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">85 / 100</a>
  </div>
  <div class="next">
    <a href="../sweetheart-2021">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      Nomadland
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      Sweetheart
    </span>
  </div>
</nav>

<article class="film slug-petite-maman-2021">
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
  <li><a href="../tomboy-2011">Tomboy</a> and <a href="../portrait-of-a-lady-on-fire-2019">Portrait of a Lady on Fire</a> by Céline Sciamma</li>
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
