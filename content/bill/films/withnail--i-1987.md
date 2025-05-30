---
title: "Withnail & I"
layout: layouts/home.njk
slug: withnail--i-1987
---

{% set film = films[slug] %}

<nav class="films">
  <a class="prev" href="../brazil-1985">Previous</a>
  <a href="../">Film list</a>
  <a class="next" href="../delicatessen-1991">Next</a>
</nav>

<p>29 / 100</p>

<article class="film">
  <div class="backdrop-and-poster">
    <img class="poster" src="../films/posters/withnail--i-1987.jpg" alt="">
    <img class="backdrop" src="../films/backdrops/withnail--i-1987.jpg" alt="">
  </div>

  <h1>Withnail & I ({{ film | filmYear }})</h1>

  <p class="director">
    Directed by <strong>{{ film | directors }}</strong>
  </p>


  <h2>
    Cast
  </h2>
  <ul>
            <li><strong>Richard E. Grant</strong> as <em>Withnail</em></li>
        <li><strong>Paul McGann</strong> as <em>...& I</em></li>
        <li><strong>Richard Griffiths</strong> as <em>Monty</em></li>
        <li><strong>Ralph Brown</strong> as <em>Danny</em></li>
        <li><strong>Michael Elphick</strong> as <em>Jake</em></li>
        <li><strong>Daragh O'Malley</strong> as <em>Irishman</em></li>
        <li><strong>Michael Wardle</strong> as <em>Isaac Parkin</em></li>
        <li><strong>Una Brandon-Jones</strong> as <em>Mrs. Parkin</em></li>
        <li><strong>Noel Johnson</strong> as <em>General</em></li>
        <li><strong>Irene Sutcliffe</strong> as <em>Waitress</em></li>
        <li><strong>Llewellyn Rees</strong> as <em>Tea Shop Proprietor</em></li>
        <li><strong>Robert Oates</strong> as <em>Policeman 1</em></li>
        <li><strong>Anthony Wise</strong> as <em>Policeman 2</em></li>
        <li><strong>Eddie Tagoe</strong> as <em>Presuming Ed</em></li>
  </ul>
</article>
<footer>
  <a href="../about">About this list</a>
</footer>
