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
    <a class="simple" href="../">37 / 100</a>
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

  <section class="cast-grid">
  <div class="cast-grid-cards">
    <div class="cast-card">
  <img src="../films/profiles/478.jpg" alt="Jamie Bell" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Jamie Bell</span>
    <span class="cast-card-character">Billy Elliot</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/480.jpg" alt="Gary Lewis" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Gary Lewis</span>
    <span class="cast-card-character">Jackie Elliot</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/477.jpg" alt="Julie Walters" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Julie Walters</span>
    <span class="cast-card-character">Mrs Wilkinson</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/481.jpg" alt="Jean Heywood" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Jean Heywood</span>
    <span class="cast-card-character">Grandma</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/479.jpg" alt="Jamie Draven" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Jamie Draven</span>
    <span class="cast-card-character">Tony Elliot</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/482.jpg" alt="Stuart Wells" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Stuart Wells</span>
    <span class="cast-card-character">Michael Caffrey</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1241113.jpg" alt="Mike Elliot" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Mike Elliot</span>
    <span class="cast-card-character">George Watson</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/3182293.jpg" alt="Billy Fane" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Billy Fane</span>
    <span class="cast-card-character">Mr Braithwaite</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">Nicola Blackwell</span>
    <span class="cast-card-character">Debbie Wilkinson</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">Carol McGuigan</span>
    <span class="cast-card-character">Librarian</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/107381.jpg" alt="Joe Renton" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Joe Renton</span>
    <span class="cast-card-character">Gary Poulson</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">Colin MacLachlan</span>
    <span class="cast-card-character">Mr Tom Wilkinson</span>
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
    <li><a href="../shallow-grave-1994">Shallow Grave</a> because of Gary Lewis</li>
<li><a href="../all-of-us-strangers-2023">All of Us Strangers</a> because of Jamie Bell</li>
  </ul>
</section>

</article>
