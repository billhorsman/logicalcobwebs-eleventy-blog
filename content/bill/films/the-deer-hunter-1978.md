---
title: "The Deer Hunter"
layout: layouts/films.njk
slug: the-deer-hunter-1978
ogImage: content/bill/films/backdrops/the-deer-hunter-1978.jpg
description: "A group of working-class friends decide to enlist in the Army during the Vietnam War and finds it to be hellish chaos -- not the noble venture they imagined. Before they left, Steven married his pregnant girlfriend -- and Michael and Nick were in love with the same woman. But all three are different men upon their return."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../the-man-who-fell-to-earth-1976"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">17 / 100</a>
  </div>
  <div class="next">
    <a href="../apocalypse-now-1979">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      The Man Who Fell to Earth
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      Apocalypse Now
    </span>
  </div>
</nav>

<article class="film slug-the-deer-hunter-1978">
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
    <li><a href="../dog-day-afternoon-1975">Dog Day Afternoon</a> because of John Cazale</li>
<li><a href="../brazil-1985">Brazil</a> and <a href="../killers-of-the-flower-moon-2023">Killers of the Flower Moon</a> because of Robert De Niro</li>
<li><a href="../fantastic-mr-fox-2009">Fantastic Mr. Fox</a> and <a href="../little-women-2019">Little Women</a> because of Meryl Streep</li>
  </ul>
</section>

</article>
