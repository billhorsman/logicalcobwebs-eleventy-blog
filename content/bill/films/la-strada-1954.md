---
title: "La Strada"
layout: layouts/home.njk
slug: la-strada-1954
---

{% set film = films[slug] %}

<nav class="films">
  <a class="prev" href="../whisky-galore-1949">Previous</a>
  <a href="../">Film list</a>
  <a class="next" href="../north-by-northwest-1959">Next</a>
</nav>

<p>3 / 100</p>

<article class="film">
  <div class="backdrop-and-poster">
    <img class="poster" src="../films/posters/la-strada-1954.jpg" alt="">
    <img class="backdrop" src="../films/backdrops/la-strada-1954.jpg" alt="">
  </div>

  <h1>La Strada ({{ film | filmYear }})</h1>

  <p class="director">
    Directed by <strong>{{ film | directors }}</strong>
  </p>


  <h2>
    Cast
  </h2>
  <ul>
            <li><strong>Giulietta Masina</strong> as <em>Gelsomina</em></li>
        <li><strong>Anthony Quinn</strong> as <em>Zampanò</em></li>
        <li><strong>Richard Basehart</strong> as <em>Il 'Matto'</em></li>
        <li><strong>Aldo Silvani</strong> as <em>Il Signor Giraffa</em></li>
        <li><strong>Marcella Rovere</strong> as <em>La Vedova</em></li>
        <li><strong>Livia Venturini</strong> as <em>La Suorina</em></li>
        <li><strong>Pietro Ceccarelli</strong> as <em>Innkeeper (uncredited)</em></li>
        <li><strong>Giovanna Galli</strong> as <em>Prostitute at the Inn (uncredited)</em></li>
        <li><strong>Gustavo Giorgi</strong> as <em>(uncredited)</em></li>
        <li><strong>Yami Kamadeva</strong> as <em>Prostitute (uncredited)</em></li>
        <li><strong>Mario Passante</strong> as <em>Waiter (uncredited)</em></li>
        <li><strong>Anna Primula</strong> as <em>Gelsomina's Mother (uncredited)</em></li>
        <li><strong>Alexandre Trannoy</strong> as <em>Juggler (uncredited)</em></li>
        <li><strong>Goffredo Unger</strong> as <em>Man Restraining Zampano from Attacking (uncredited)</em></li>
        <li><strong>Nazzareno Zamperla</strong> as <em>Man Restraining Zampano from Attacking (uncredited)</em></li>
  </ul>
</article>
<footer>
  <a href="../about">About this list</a>
</footer>
