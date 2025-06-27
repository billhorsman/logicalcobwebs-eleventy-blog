---
title: "Billy Elliot"
layout: layouts/films.njk
slug: billy-elliot-2000
ogImage: content/bill/films/backdrops/billy-elliot-2000.jpg
description: "County Durham, England, 1984. The miners' strike has started and the police have started coming up from Bethnal Green, starting a class war with the lower classes suffering. Caught in the middle of the conflict is 11-year old Billy Elliot, who, after leaving his boxing club for the day, stumbles upon a ballet class and finds out that he's naturally talented. He practices with his teacher Mrs. Wilkinson for an upcoming audition in Newcastle-upon Tyne for the royal Ballet school in London."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../the-talented-mr-ripley-1999"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">41 / 100</a>
  </div>
  <div class="next">
    <a href="../amlie-2001">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      The Talented Mr. Ripley
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      Amélie
    </span>
  </div>
</nav>

<article class="film slug-billy-elliot-2000">
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
  <li><a href="../shallow-grave-1994">Shallow Grave</a> (1994) by Gary Lewis</li>
<li><a href="../all-of-us-strangers-2023">All of Us Strangers</a> (2023) by Jamie Bell</li>
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
