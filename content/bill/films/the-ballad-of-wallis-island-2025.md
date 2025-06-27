---
title: "The Ballad of Wallis Island"
layout: layouts/films.njk
slug: the-ballad-of-wallis-island-2025
ogImage: content/bill/films/backdrops/the-ballad-of-wallis-island-2025.jpg
description: "Eccentric lottery winner Charles lives alone on a remote island but dreams of hiring his favourite musician, Herb McGwyer, to play an exclusive, private gig. Unbeknownst to Herb, Charles has also hired Herb’s ex-bandmate and ex-girlfriend, Nell, with her new husband in town, to perform the old favourites. As tempers flare and old tensions resurface, the stormy weather traps them all on the island and Charles desperately looks for a way to salvage his dream gig."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../killers-of-the-flower-moon-2023"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
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
      Killers of the Flower Moon
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      End of list
    </span>
  </div>
</nav>

<article class="film slug-the-ballad-of-wallis-island-2025">
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
</article>
