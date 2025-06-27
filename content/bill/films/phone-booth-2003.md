---
title: "Phone Booth"
layout: layouts/films.njk
slug: phone-booth-2003
ogImage: content/bill/films/backdrops/phone-booth-2003.jpg
description: "A slick New York publicist who picks up a ringing receiver in a phone booth is told that if he hangs up, he'll be killed... and the little red light from a laser rifle sight is proof that the caller isn't kidding."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../the-bourne-identity-2002"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">47 / 100</a>
  </div>
  <div class="next">
    <a href="../the-motorcycle-diaries-2004">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      The Bourne Identity
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      The Motorcycle Diaries
    </span>
  </div>
</nav>

<article class="film slug-phone-booth-2003">
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
  <li><a href="../fight-club-1999">Fight Club</a> (1999) by Jared Leto</li>
<li><a href="../dallas-buyers-club-2013">Dallas Buyers Club</a> (2013) by Jared Leto</li>
<li><a href="../house-of-gucci-2021">House of Gucci</a> (2021) by Jared Leto</li>
<li><a href="../ghost-dog-the-way-of-the-samurai-1999">Ghost Dog: The Way of the Samurai</a> (1999) by Forest Whitaker</li>
<li><a href="../in-bruges-2008">In Bruges</a> (2008) by Colin Farrell</li>
<li><a href="../the-banshees-of-inisherin-2022">The Banshees of Inisherin</a> (2022) by Colin Farrell</li>
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
