---
title: "The Secret Agent"
layout: layouts/films.njk
slug: the-secret-agent-2025
ogImage: content/bill/films/backdrops/the-secret-agent-2025.jpg
description: "Brazil, 1977. Marcelo, a technology expert in his early 40s, is on the run. Hoping to reunite with his son, he travels to Recife during Carnival but soon realizes that the city is not the safe haven he was expecting."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../nouvelle-vague-2025"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">100 / 100</a>
  </div>
  <div class="next">
    <span>Next <i class="fa-solid fa-chevron-right fa-xs"></i></span>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      Nouvelle Vague
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      End of list
    </span>
  </div>
</nav>

<article class="film slug-the-secret-agent-2025">
  <div class="backdrop-and-poster">
    <img class="poster" src="../films/posters/{{ slug }}.jpg" alt="" eleventy:ignore>
    <img class="backdrop" src="../films/backdrops/{{ slug }}.jpg" alt="" eleventy:ignore>
  </div>

  <h1>{{ film.title }} ({{ film | filmYear }})</h1>

  <p>
    {%- if film.language -%}Language: {{ film.language }}.{% endif %}
    Also known as O Agente Secreto.
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
  <img src="../films/profiles/52583.jpg" alt="Wagner Moura" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Wagner Moura</span>
    <span class="cast-card-character">Armando Solimões / Marcelo Alves / Adult Fernando Solimões</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/4497686.jpg" alt="Tânia Maria" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Tânia Maria</span>
    <span class="cast-card-character">Dona Sebastiana</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/2652283.jpg" alt="Alice Carvalho" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Alice Carvalho</span>
    <span class="cast-card-character">Fátima</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/228057.jpg" alt="Maria Fernanda Cândido" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Maria Fernanda Cândido</span>
    <span class="cast-card-character">Elza / Sara Guébert</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1475555.jpg" alt="Gabriel Leone" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Gabriel Leone</span>
    <span class="cast-card-character">Bobbi Borba</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1646.jpg" alt="Udo Kier" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Udo Kier</span>
    <span class="cast-card-character">Hans</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/2346588.jpg" alt="Carlos Francisco" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Carlos Francisco</span>
    <span class="cast-card-character">Seu Alexandre</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/2346576.jpg" alt="Thomás Aquino" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Thomás Aquino</span>
    <span class="cast-card-character">Valdemar</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/583446.jpg" alt="Hermila Guedes" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Hermila Guedes</span>
    <span class="cast-card-character">Claudia</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/2375667.jpg" alt="Robério Diógenes" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Robério Diógenes</span>
    <span class="cast-card-character">Euclides Cavalcante</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/143500.jpg" alt="Roney Villela" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Roney Villela</span>
    <span class="cast-card-character">Augusto Borba</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1691901.jpg" alt="Isabél Zuaa" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Isabél Zuaa</span>
    <span class="cast-card-character">Thereza Vitória</span>
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

  
</article>
