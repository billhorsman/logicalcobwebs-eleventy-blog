---
title: "The French Dispatch"
layout: layouts/films.njk
slug: the-french-dispatch-2021
ogImage: content/bill/films/backdrops/the-french-dispatch-2021.jpg
description: "The staff of an American magazine based in France puts out its last issue, with stories featuring an artist sentenced to life imprisonment, student riots, and a kidnapping resolved by a chef."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../sweetheart-2021"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">87 / 100</a>
  </div>
  <div class="next">
    <a href="../the-power-of-the-dog-2021">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      Sweetheart
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      The Power of the Dog
    </span>
  </div>
</nav>

<article class="film slug-the-french-dispatch-2021">
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
  <li><a href="../fargo-1996">Fargo</a> by Frances McDormand and Steve Park</li>
<li><a href="../nomadland-2021">Nomadland</a> and <a href="../the-tragedy-of-macbeth-2021">The Tragedy of Macbeth</a> by Frances McDormand</li>
<li><a href="../asteroid-city-2023">Asteroid City</a> by Steve Park, Edward Norton, Jason Schwartzman, Willem Dafoe, Wes Anderson, Jarvis Cocker, Adrien Brody, Tilda Swinton, Tony Revolori, Bob Balaban, Fisher Stevens, Jeffrey Wright, Mohamed Belhadjine, Nicolas Avinée, Rupert Friend, Tom Hudson, Stéphane Bak, Liev Schreiber, Damien Bonnard, Rodolphe Pauly and Eliel Ford</li>
<li><a href="../fight-club-1999">Fight Club</a> by Edward Norton</li>
<li><a href="../the-grand-budapest-hotel-2014">The Grand Budapest Hotel</a> by Edward Norton, Jason Schwartzman, Wallace Wolodarsky, Willem Dafoe, Bill Murray, Wes Anderson, Owen Wilson, Adrien Brody, Mathieu Amalric, Saoirse Ronan, Léa Seydoux, Tilda Swinton, Tony Revolori, Larry Pine, Bob Balaban and Fisher Stevens</li>
<li><a href="../fantastic-mr-fox-2009">Fantastic Mr. Fox</a> by Jason Schwartzman, Wallace Wolodarsky, Willem Dafoe, Bill Murray, Wes Anderson, Jarvis Cocker, Owen Wilson and Adrien Brody</li>
<li><a href="../the-lighthouse-2019">The Lighthouse</a> by Willem Dafoe</li>
<li><a href="../dallas-buyers-club-2013">Dallas Buyers Club</a> by Griffin Dunne</li>
<li><a href="../interstellar-2014">Interstellar</a> and <a href="../dune-2021">Dune</a> by Timothée Chalamet</li>
<li><a href="../sink-or-swim-2018">Sink or Swim</a> by Mathieu Amalric and Félix Moati</li>
<li><a href="../one-fine-morning-2022">One Fine Morning</a> by Léa Seydoux and Sharif Andoura</li>
<li><a href="../cest-la-vie-2017">C'est la vie!</a> by Benjamin Lavernhe</li>
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
