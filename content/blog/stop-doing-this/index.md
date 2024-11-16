---
title: Stop doing this
description: Remembering to fail later (in your automatic tests)
excerpt: Remembering to fail later (in your automatic tests)
canonical: https://logicalcobwebs.com/blog/stop-doing-this
ogImage: content/blog/stop-doing-this/stop-sign.jpg
date: 2017-11-27
author: Bill Horsman
tags:
  - coding
  - testing
  - Ruby on Rails
---

That quick fix that you hate to do but can’t avoid. Usually it’s because there’s a problem in production that will take a few days to fix properly but waiting that long will cause problems. Or a workaround for a broken third party component while you wait for them to fix it. Whatever the reason, you’ve written some code and it’s temporary.

The trouble is, temporary code has a habit of lasting a long time. So when I do this, I think of a date by which it should be fixed by and build in a failure in my tests based on that date. In Ruby on Rails, I do it like this:

```rb
# config/initializers/_stop_doing_this.rb
def stop_doing_this(year, month, day)
  return unless Rails.env.test?
  if Date.today > Date.civil(year, month, day)
    raise “Stop doing this”
  end
end
```

Note: I use an underscore prefix on the initializer name so that it is available to other initializers (they are loaded alphabetically).

And then, in my code, I do something like:

```rb
class Foo
  def add_widget(widget)
    stop_doing_this 2017, 12, 25
    # My bad fix
  end
end
```

That way, on 25 Dec, your tests will fail and you will be reminded to do something about it. It only fails in the tests, not in production. At this point, you either get to fix it properly or advance the date forward a little bit to buy you some time.

It’s a pain to advance the date forward. I deliberately only move it by a couple of weeks or so because that pain is what’s going to motivate you to do the right thing.
