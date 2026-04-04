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
    <a class="simple" href="../">27 / 100</a>
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

  <section class="cast-grid">
  <div class="cast-grid-cards">
    <div class="cast-card">
  <img src="../films/profiles/17258.jpg" alt="Kerry Fox" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Kerry Fox</span>
    <span class="cast-card-character">Juliet Miller</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/2040.jpg" alt="Christopher Eccleston" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Christopher Eccleston</span>
    <span class="cast-card-character">David Stephens</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/3061.jpg" alt="Ewan McGregor" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Ewan McGregor</span>
    <span class="cast-card-character">Alex Law</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/25136.jpg" alt="Ken Stott" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Ken Stott</span>
    <span class="cast-card-character">Detective Inspector McCall</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/20056.jpg" alt="Keith Allen" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Keith Allen</span>
    <span class="cast-card-character">Hugo</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/149359.jpg" alt="Colin McCredie" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Colin McCredie</span>
    <span class="cast-card-character">Cameron</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/3064.jpg" alt="Peter Mullan" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Peter Mullan</span>
    <span class="cast-card-character">Andy</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/480.jpg" alt="Gary Lewis" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Gary Lewis</span>
    <span class="cast-card-character">Visitor</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">Frances Low</span>
    <span class="cast-card-character">Doctor</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">Robert David MacDonald</span>
    <span class="cast-card-character">Lumsden</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/2220.jpg" alt="Tony Curran" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Tony Curran</span>
    <span class="cast-card-character">Travel Agent</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/8999.jpg" alt="John Hodge" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">John Hodge</span>
    <span class="cast-card-character">DC Mitchell</span>
  </div>
</div>
  </div>
</section>

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
