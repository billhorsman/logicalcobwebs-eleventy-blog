---
title: "Phone Booth"
layout: layouts/home.njk
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
    <a href="../">Film list</a>
  </div>
  <div class="next">
    <a href="../the-motorcycle-diaries-2004">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
</nav>

<p>49 / 100</p>

<article class="film slug-phone-booth-2003">
  <div class="backdrop-and-poster">
    <img class="poster" src="../films/posters/{{ slug }}.jpg" alt="">
    <img class="backdrop" src="../films/backdrops/{{ slug }}.jpg" alt="">
  </div>

  <h1>{{ film.title }} ({{ film | filmYear }})</h1>

  

  <p class="director">
    Directed by <strong>{{ film | directors }}</strong>
  </p>

  <blockquote> 
    {{ film.overview }} <em>—&nbsp;<a href="https://www.themoviedb.org/movie/{{ film.id }}">TMDB</a></em>
  </blockquote> 

  {%- if films.reviews[slug] -%}
    <blockquote> 
      {{ films.reviews[slug] | safe }} <em>—&nbsp;<a href="/bill">Bill</a></em>
    </blockquote> 
  {%- endif -%}

  <h2>
    Cast
  </h2>
  <ul>
    {%- for cast in film.credits.cast -%}
      <li>
        {{ cast.name }} as <em>{{ cast.character }}</em>
      </li>
    {%- endfor -%}
  </ul>

  <h2>
    Crew
  </h2>
  <ul>
    {%- for crew in film.credits.crew -%}
      <li>
        {{ crew.name }} &mdash; <em>{{ crew.job }}</em>
      </li>
    {%- endfor -%}
  </ul>
</article>
<footer>
  <a href="../about">About this list</a>
</footer>
