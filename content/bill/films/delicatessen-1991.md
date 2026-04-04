---
title: "Delicatessen"
layout: layouts/films.njk
slug: delicatessen-1991
ogImage: content/bill/films/backdrops/delicatessen-1991.jpg
description: "In a post-apocalyptic world, the residents of an apartment above the butcher shop receive an occasional delicacy of meat, something that is in low supply. A young man new in town falls in love with the butcher's daughter, which causes conflicts in her family, who need the young man for other business-related purposes."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../withnail--i-1987"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">25 / 100</a>
  </div>
  <div class="next">
    <a href="../night-on-earth-1991">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      Withnail & I
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      Night on Earth
    </span>
  </div>
</nav>

<article class="film slug-delicatessen-1991">
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
    <li><a href="../diva-1981">Diva</a> because of Dominique Pinon</li>
<li><a href="../amlie-2001">Amélie</a> because of Dominique Pinon, Ticky Holgado, Rufus, Patrick Paroux and Jean-Pierre Jeunet</li>
<li><a href="../micmacs-2009">Micmacs</a> because of Dominique Pinon and Jean-Pierre Jeunet</li>
<li><a href="../black-hawk-down-2001">Black Hawk Down</a> because of Nikky Smedley</li>
  </ul>
</section>

</article>
