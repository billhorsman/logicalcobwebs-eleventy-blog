---
title: "Shallow Grave"
layout: layouts/films.njk
slug: shallow-grave-1994
ogImage: content/bill/films/backdrops/shallow-grave-1994.jpg
description: "When David, Juliet, and Alex find their new roommate dead with a large sum of money, they agree to hide the body and keep the cash. However, this newfound fortune gradually corrodes their friendship."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../lon-the-professional-1994"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">31 / 100</a>
  </div>
  <div class="next">
    <a href="../fargo-1996">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      Léon: The Professional
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      Fargo
    </span>
  </div>
</nav>

<article class="film slug-shallow-grave-1994">
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
    <li><a href="../24-hour-party-people-2002">24 Hour Party People</a> because of Christopher Eccleston and Keith Allen</li>
<li><a href="../trainspotting-1996">Trainspotting</a> because of Ewan McGregor, Keith Allen, Peter Mullan, Victor Eadie, Billy Riddoch and Danny Boyle</li>
<li><a href="../black-hawk-down-2001">Black Hawk Down</a> because of Ewan McGregor</li>
<li><a href="../billy-elliot-2000">Billy Elliot</a> because of Gary Lewis</li>
  </ul>
</section>

</article>
