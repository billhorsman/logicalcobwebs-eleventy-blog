---
title: Vanilla Marathon
description: An animation of a marathon distance using only HTML, Modern CSS and Vanilla JavaScript. Vanilla JavaScript means using no libraries of frameworks.
excerpt: Running a marathon, virtually.
date: 2024-05-23
author: Bill Horsman
tags:
  - coding
  - marathon
---

## How it started

My colleague Vicky donated to Macmillan Cancer Support on my [JustGiving page](https://www.justgiving.com/page/bill-runs-edinburgh-marathon-2024) for the Edinburgh Marathon on 26 May 2024. Her accompanying comment was a six-line C program, which was pretty cool.

```
int distance = 26;
int position = 0;
while (position < distance) {
  position ++;
}
printf("Finished!\n");
```

I could do that in JavaScript and put it on the Internet, I thought — so I built it and put it up on [Netlify](https://www.netlify.com/).

[Step 1](https://vanilla-marathon.logicalcobwebs.com/1/)

<iframe title="Vanilla Marathon Step 1" src="https://vanilla-marathon.logicalcobwebs.com/1/"></iframe>

There are no libraries or frameworks needed, just HTML, CSS and some [Vanilla JavaScript](http://vanilla-js.com/) (a parody project that has no code). It was quite refreshing and I wanted to keep it simple but maybe a bit more intetesting&hellip;

## Animation

I liked the idea of a little [Sprite](<https://en.wikipedia.org/wiki/Sprite_(computer_graphics)>) that moved across the bottom of the page. So I added a black circle, increased the number of steps and made it move from left to right.

[Step 2](https://vanilla-marathon.logicalcobwebs.com/2/)

<iframe title="Vanilla Marathon Step 2" src="https://vanilla-marathon.logicalcobwebs.com/2/"></iframe>

## Rolling

That sprite looks like it _sliding_ not _rolling_. Let's make it look like a cog and rotate it as it moves.

[Step 3](https://vanilla-marathon.logicalcobwebs.com/3/)

<iframe title="Vanilla Marathon Step 3" src="https://vanilla-marathon.logicalcobwebs.com/3/"></iframe>

## Icons

I shared this version with Vicky and she asked what I'd used to build it. Just HTML, CSS and JavaScript I said, proudly. Oh, and [Fontawesome](https://fontawesome.com/) for the icon &hellip; I'm going to have to do my own icons aren't I?

That took longer than I anticipated. I did them in Affinity Designer, exported to SVG and then a small Ruby script exported them into a javascript file, which would replace HTML like this with an SVG icon:

```
<i class="
  vanilla-icon
  va-person-running">
</i>
```

That's a simlar approach to Fontawesome, but a lot less elegant. Drawing icons is hard. I definitely took inspiration from the Fontawesome one but I started from scratch. I guess the best you can say about mine is that it's unique.

<section class="icon-gallery">
  <div>
    <i>
      <i class="fa fa-fw fa-person-running"></i>
    </i>
    <span>Fontawesome</span>
  </div>
  <div>
    <i>
      <svg viewBox="0 0 448 512" version="1.1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" xml:space="preserve" xmlns:serif="http://www.serif.com/">
        <g>
          <g transform="matrix(-1.27137e-16,0.613643,-1.199,-4.16334e-17,751.741,-51.8199)">
            <path d="M361.481,260L302.36,260C410.971,298.61 391.234,335.598 286.518,376.764L361.481,393.444C462.517,330.058 459.11,296.243 361.481,260Z"/>
          </g>
          <g transform="matrix(1.33929,0,0,1,-106.964,11.3745)">
            <path d="M337.216,468.625L292.667,459.625L274,488.284L330,488.284L337.216,468.625Z"/>
          </g>
          <g transform="matrix(0.965926,-0.258819,0.258819,0.965926,-80.3087,72.7921)">
            <path d="M272.8,500.284L217.8,500.284C349.104,409.806 231.714,357.356 203.556,250.452C233.493,278.974 259.702,277.893 270.601,266.319C333.57,344.273 362.922,416.325 272.8,500.284Z"/>
          </g>
          <g transform="matrix(0.419562,0.4478,0.99207,-0.929511,-263.008,327.83)">
            <path d="M360.485,263.854L302.36,260C197.349,287.197 191.596,329.776 280.375,382.405L331.141,366.984C217.276,296.667 286.305,301.274 360.485,263.854Z"/>
          </g>
          <g transform="matrix(-8.20076e-17,1.33929,-1,-6.12323e-17,502.284,-26.9643)">
            <path d="M330,466L282.96,460.284L274,488.284L326.267,488.284L330,466Z"/>
          </g>
          <g transform="matrix(-0.5,-0.866025,-0.866025,0.5,554.987,338.478)">
            <path d="M269.175,469.269L234.534,489.269C191.284,410.831 68.762,392.143 174.836,285.868C207.388,298.25 229.111,301.875 251.797,299.169C179.592,391.469 172.147,370.435 269.175,469.269Z"/>
          </g>
          <g transform="matrix(0.937977,0.346697,-0.346697,0.937977,134.116,-47.0363)">
            <path d="M220,111.279C220,107.877 218.27,104.708 215.409,102.868C212.548,101.028 208.947,100.769 205.852,102.18C187.486,110.896 169.243,111.143 151.183,101.367C148.839,100.123 146.014,100.199 143.739,101.567C141.465,102.935 140.073,105.395 140.073,108.049C140,108.05 140,108.05 140,108.05L140,245C140,267.091 157.909,285 180,285L180,285C202.091,285 220,267.091 220,245L220,111.279Z"/>
          </g>
          <g transform="matrix(1,0,0,1,112,1)">
            <circle cx="174" cy="65" r="40"/>
          </g>
        </g>
      </svg>
    </i>
    Not-so-awesome
  </div>
</section>

## All the edges

I decided it would be more interesting if the sprite did a tour of all four sides of the browser. So I added some more maths to change the `left` and `top` and made the person rotate each time they changed sides.

[Step 4](https://vanilla-marathon.logicalcobwebs.com/4/)

<iframe title="Vanilla Marathon Step 4" src="https://vanilla-marathon.logicalcobwebs.com/4/"></iframe>

## Avoid the sprite!

I didn't like how the sprite trampled over the heading.

[Step 5](https://vanilla-marathon.logicalcobwebs.com/5/)

<iframe title="Vanilla Marathon Step 5" src="https://vanilla-marathon.logicalcobwebs.com/5/"></iframe>

## How it ended

Moar icons!

<section class="icon-gallery">
  <div>
    <i>
      <svg viewBox="0 0 448 512" version="1.1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" xml:space="preserve" xmlns:serif="http://www.serif.com/">
        <g>
          <path d="M293.167,241.584C298.249,244.124 301.459,249.318 301.459,255C301.459,260.682 298.249,265.876 293.167,268.416C255.226,287.387 186.637,321.682 151.708,339.146C147.058,341.471 141.536,341.222 137.114,338.489C132.692,335.756 130,330.928 130,325.729C130,289.263 130,220.737 130,184.271C130,179.072 132.692,174.244 137.114,171.511C141.536,168.778 147.058,168.529 151.708,170.854C186.637,188.318 255.226,222.613 293.167,241.584Z"/>
        </g>
      </svg>
    </i>
    Play

  </div>
  <div>
    <i>
      <svg viewBox="0 0 448 512" version="1.1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" xml:space="preserve" xmlns:serif="http://www.serif.com/">
        <g>
          <path d="M402.88,135.029L402.88,115L259,101.279L259,135.029L402.88,135.029Z"/>
          <path d="M335,477.375L260,471.375L260,499.659L335,499.659L335,477.375Z"/>
          <path d="M312.679,485.424L259.553,499.659L202.783,287.791L203.116,287.702C212.049,289.998 221.674,290.094 231.208,287.54C240.742,284.985 249.029,280.089 255.618,273.634L255.909,273.556L312.679,485.424Z"/>
          <path d="M37.12,135.029L37.12,115L181,101.279L181,135.029L37.12,135.029Z"/>
          <path d="M103,477.375L178,471.375L178,499.659L103,499.659L103,477.375Z"/>
          <path d="M260,111.279C260,107.877 258.27,104.708 255.409,102.868C252.548,101.028 248.947,100.769 245.852,102.18C228.495,110.421 211.245,111.09 194.157,102.885C191.091,101.467 187.514,101.71 184.668,103.529C181.822,105.348 180.1,108.492 180.1,111.87C180,142.091 180,203.997 180,245C180,267.091 197.909,285 220,285L220,285C242.091,285 260,267.091 260,245L260,111.279Z"/>
          <rect x="181" y="101.279" width="80" height="30"/>
          <circle cx="220" cy="60" r="40"/>
          <path d="M216.242,357.408L196.527,283.832C191.1,281.193 186.239,277.727 182.061,273.634L181.77,273.556L125,485.424L178.126,499.659L216.242,357.408Z"/>
        </g>
      </svg>
    </i>
    Paused
  </div>
  <div>
    <i>
      <svg viewBox="0 0 448 512" version="1.1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" xml:space="preserve" xmlns:serif="http://www.serif.com/">
        <g>
          <g transform="matrix(-1.27137e-16,0.613643,-1.199,-4.16334e-17,751.741,-51.8199)">
            <path d="M361.481,260L302.36,260C410.971,298.61 391.234,335.598 286.518,376.764L361.481,393.444C462.517,330.058 459.11,296.243 361.481,260Z"/>
          </g>
          <g transform="matrix(1.33929,0,0,1,-106.964,11.3745)">
            <path d="M337.216,468.625L292.667,459.625L274,488.284L330,488.284L337.216,468.625Z"/>
          </g>
          <g transform="matrix(0.965926,-0.258819,0.258819,0.965926,-80.3087,72.7921)">
            <path d="M272.8,500.284L217.8,500.284C349.104,409.806 231.714,357.356 203.556,250.452C233.493,278.974 259.702,277.893 270.601,266.319C333.57,344.273 362.922,416.325 272.8,500.284Z"/>
          </g>
          <g transform="matrix(0.419562,0.4478,0.99207,-0.929511,-263.008,327.83)">
            <path d="M360.485,263.854L302.36,260C197.349,287.197 191.596,329.776 280.375,382.405L331.141,366.984C217.276,296.667 286.305,301.274 360.485,263.854Z"/>
          </g>
          <g transform="matrix(-8.20076e-17,1.33929,-1,-6.12323e-17,502.284,-26.9643)">
            <path d="M330,466L282.96,460.284L274,488.284L326.267,488.284L330,466Z"/>
          </g>
          <g transform="matrix(-0.5,-0.866025,-0.866025,0.5,554.987,338.478)">
            <path d="M269.175,469.269L234.534,489.269C191.284,410.831 68.762,392.143 174.836,285.868C207.388,298.25 229.111,301.875 251.797,299.169C179.592,391.469 172.147,370.435 269.175,469.269Z"/>
          </g>
          <g transform="matrix(0.937977,0.346697,-0.346697,0.937977,134.116,-47.0363)">
            <path d="M220,111.279C220,107.877 218.27,104.708 215.409,102.868C212.548,101.028 208.947,100.769 205.852,102.18C187.486,110.896 169.243,111.143 151.183,101.367C148.839,100.123 146.014,100.199 143.739,101.567C141.465,102.935 140.073,105.395 140.073,108.049C140,108.05 140,108.05 140,108.05L140,245C140,267.091 157.909,285 180,285L180,285C202.091,285 220,267.091 220,245L220,111.279Z"/>
          </g>
          <g transform="matrix(1,0,0,1,112,1)">
            <circle cx="174" cy="65" r="40"/>
          </g>
        </g>
      </svg>
    </i>
    Running
  </div>
  <div>
    <i>
      <svg viewBox="0 0 448 512" version="1.1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" xml:space="preserve" xmlns:serif="http://www.serif.com/">
        <g>
          <path d="M12,472.716L87,466.716L87,495L12,495L12,472.716Z"/>
          <path d="M100,340C108.011,322.578 113.835,322.052 129.818,321.14C139.008,320.521 147.642,325.587 151.586,333.912C164.087,360.295 190,415 190,415C168,421 158.317,452.071 165,465L100,340Z"/>
          <path d="M100,340L129,372L87.044,495C87.044,495 58.155,476.944 50,475L100,340Z"/>
          <path d="M170.038,471.951C167.626,466.805 166.279,461.06 166.279,455L166.279,455C166.279,432.909 184.187,415 206.279,415C247.281,415 309.188,415 339.409,415.1C342.786,415.1 345.931,416.822 347.75,419.668C349.569,422.514 349.812,426.091 348.393,429.157C340.188,446.245 340.858,463.495 349.099,480.852C350.51,483.947 350.251,487.548 348.411,490.409C346.571,493.27 343.402,495 340,495L206.279,495C200.179,495 194.399,493.635 189.227,491.194L314.553,476.745C318.65,476.272 322.391,474.192 324.954,470.961C327.517,467.73 328.691,463.614 328.219,459.517C328.194,459.296 328.168,459.074 328.142,458.852C327.625,454.362 325.256,450.292 321.608,447.624C317.96,444.956 313.364,443.932 308.928,444.799L170.038,471.951C173.991,480.389 180.804,487.219 189.227,491.194L156.213,495L153.919,475.103L170.038,471.951Z"/>
          <circle cx="390" cy="455" r="40"/>
        </g>
      </svg>
    </i>
    Resting

  </div>
</section>

I added keyboard shortcuts with keys that you can click. Switched to defaulting to kilometres be added a toggle so you can switch back to miles if that's what you're used to.

You can now get faster or slower too, with your pace displayed in minutes per kilometre/mile.

This is what I built: [Vanilla Marathon](https://vanilla-marathon.logicalcobwebs.com/)

<iframe title="Vanilla Marathon Final Version" src="https://vanilla-marathon.logicalcobwebs.com/"></iframe>

## Todo

- [ ] Publish source code
- [ ] Add those extra icons
