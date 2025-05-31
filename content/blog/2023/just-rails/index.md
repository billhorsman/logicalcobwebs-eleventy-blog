---
title: Just Rails
description: Building an app with Rails 7 and Hotwire
excerpt: Stepping away from Webpack and JSON API driven UIs
ogImage: content/blog/2023/just-rails/rails.png
date: 2023-12-15
author: Bill Horsman
tags:
  - coding
  - rails
  - hotwire
---

We needed to build an app from scratch a couple of months ago which is always an opportunity to try out some new ideas. After a few years of building Vue apps (living inside Ruby on Rails) we decided to see what would happen if we stuck to Rails' recommendations &mdash; simplifying everything as much as possible.

### Goals

- Faster to develop, build and use
- Simpler to maintain
- No compromise on user experience

### How we did it

- [Ruby on Rails](https://rubyonrails.org/)
- [Hotwire](https://hotwired.dev/) &mdash; Turbo and Stimulus
- That's about it

### Conclusion

Two months in and we have no regrets. We've got a fast, simple app that was a joy to develop. It felt like adding new features was 10x easier than with Vue. That could be because of a few things:

- We had a pretty good knowledge of the domain and knew from the outset how most things should work.
- Every new app has a honeymoon period where features are easy to add and tests run fast.
- We were focused on building this app over a short period of time and it was a very productive few weeks.

But I do think that the most significant reasons were:

- Not having to build an API **and** and client UI separately.
- [Turbo](https://turbo.hotwired.dev/) really does make development fast.
- No build step or transpiling of Javascript (e.g. no Webpack).

And a lot of that is possible because modern browsers are so good:

- Modern Javascript support is widespread.
- HTTP2 means you can download many small files without overhead (no need to bundle things into one big file).
