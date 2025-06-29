---
title: "Fargo"
layout: layouts/films.njk
slug: fargo-1996
ogImage: content/bill/films/backdrops/fargo-1996.jpg
description: "Jerry, a small-town Minnesota car salesman is bursting at the seams with debt... but he's got a plan. He's going to hire two thugs to kidnap his wife in a scheme to collect a hefty ransom from his wealthy father-in-law. It's going to be a snap and nobody's going to get hurt... until people start dying. Enter Police Chief Marge, a coffee-drinking, parka-wearing - and extremely pregnant - investigator who'll stop at nothing to get her man. And if you think her small-time investigative skills will give the crooks a run for their ransom... you betcha!"
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../shallow-grave-1994"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">30 / 100</a>
  </div>
  <div class="next">
    <a href="../trainspotting-1996">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      Shallow Grave
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      Trainspotting
    </span>
  </div>
</nav>

<article class="film slug-fargo-1996">
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
    <li><a href="../nomadland-2021">Nomadland</a> because of Frances McDormand and Warren Keith</li>
<li><a href="../the-french-dispatch-2021">The French Dispatch</a> because of Frances McDormand and Steve Park</li>
<li><a href="../the-tragedy-of-macbeth-2021">The Tragedy of Macbeth</a> because of Frances McDormand and Joel Coen</li>
<li><a href="../magnolia-1999">Magnolia</a> because of William H. Macy</li>
<li><a href="../the-big-lebowski-1998">The Big Lebowski</a> because of Steve Buscemi, Peter Stormare, Warren Keith and Joel Coen</li>
<li><a href="../lucky-2017">Lucky</a> because of John Carroll Lynch</li>
<li><a href="../asteroid-city-2023">Asteroid City</a> because of Steve Park</li>
<li><a href="../the-straight-story-1999">The Straight Story</a> because of Sally Wingert</li>
<li><a href="../no-country-for-old-men-2007">No Country for Old Men</a> because of Joel Coen</li>
  </ul>
</section>

</article>
