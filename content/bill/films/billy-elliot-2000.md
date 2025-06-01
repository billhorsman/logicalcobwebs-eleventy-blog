---
title: "Billy Elliot"
layout: layouts/home.njk
slug: billy-elliot-2000
ogImage: content/bill/films/backdrops/billy-elliot-2000.jpg
description: "County Durham, England, 1984. The miners' strike has started and the police have started coming up from Bethnal Green, starting a class war with the lower classes suffering. Caught in the middle of the conflict is 11-year old Billy Elliot, who, after leaving his boxing club for the day, stumbles upon a ballet class and finds out that he's naturally talented. He practices with his teacher Mrs. Wilkinson for an upcoming audition in Newcastle-upon Tyne for the royal Ballet school in London."
---

{% set film = films[slug] %}

<nav class="films">
  <a class="prev" href="../the-talented-mr-ripley-1999">Previous</a>
  <a href="../">Film list</a>
  <a class="next" href="../amlie-2001">Next</a>
</nav>

<p>43 / 100</p>

<article class="film slug-billy-elliot-2000">
  <div class="backdrop-and-poster">
    <img class="poster" src="../films/posters/{{ slug }}.jpg" alt="">
    <img class="backdrop" src="../films/backdrops/{{ slug }}.jpg" alt="">
  </div>

  <h1>{{ film.title }} ({{ film | filmYear }})</h1>

  

  <p class="director">
    Directed by <strong>{{ film | directors }}</strong>
  </p>

  {% if films.reviews[slug] %}
    <blockquote> 
      {{ films.reviews[slug] | safe }} <em>— Bill</em>
    </blockquote> 
  {% endif %}

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
