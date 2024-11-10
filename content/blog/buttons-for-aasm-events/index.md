---
title: Buttons for AASM events
description: Don't make your UI think harder than it has to
excerpt: Don't make your UI think harder than it has to
date: 2017-05-12
author: Bill Horsman
tags:
  - coding
  - Ruby on Rails
---

This might be of interest if you use the [AASM](https://github.com/aasm/aasm) state machine gem and [Ruby on Rails](http://rubyonrails.org/). If not, well, I warned you.

Let’s take the example, from the AASM documentation, of a Job. You can run it, sleep it or clean it. We want to have an UI that allows an operator to control the job with a series of buttons.

```rb
class Job
  include AASM

  aasm do
    state :sleeping, initial: true
    state :running, :cleaning

    event :run do
      transitions from: :sleeping, to: :running
    end

    event :clean do
      transitions from: :running, to: :cleaning
    end

    event :sleep do
      transitions from: [:running, :cleaning], to: :sleeping
    end
  end

end
```

The buttons that need to be shown depend on the state. You write some code that copes with that:

```erb
<% if @job.sleeping? %>
 <%= link_to "Run", run_job_path, method: :patch, class: "btn btn-success" %>
<% elsif @job.running? %>
 <%= link_to "Clean", clean_job_path, method: :patch, class: "btn btn-default" %>
 <%= link_to "Shut Down", sleep_job_path, method: :patch, class: "btn btn-danger" %>
<% elsif @job.cleaning? %>
 <%= link_to "Shut Down", sleep_job_path, method: :patch, class: "btn btn-danger" %>
<% end %>
```

And if the job is <em>running</em>:

<div class="btns">
  <button type="button" class="btn">Clean</button>
  <button type="button" class="btn btn-danger">Shutdown</button>
</div>


## Step 1 — Less to think about

But that’s hard to get right (you need to understand all the events and their transitions) and will need changing if the model changes. However, AASM allows you to inspect the state and see what is possible. Instead of working out which events are possible from which state we just ask what events are possible from the current state and check if our event is in that list. It makes things a little better:

```erb
<% @job.aasm.events(possible: true).map(&:name).tap do |possible| %>
  <% if possible.include? :run %>
    <%= link_to "Run", run_job_path, method: :patch, class: "btn btn-success" %>
  <% end %>

  <% if possible.include? :clean %>
    <%= link_to "Clean", clean_job_path, method: :patch, class: "btn btn-default" %>
  <% end %>
  <% if possible.include? :sleep %>
    <%= link_to "Shut Down", sleep_job_path, method: :patch, class: "btn btn-danger" %>
  <% end %>
<% end %>
```

That’s still lots of typing but at least we don’t need to know about the relationship between state and events.

## Step 2 — Less code too

We can take this a step further though. Instead of checking whether a specific event is in the list of possible events why not just iterate over the list of events? We’ll need to adapt our code a little, especially the path:

```erb
<% @job.aasm.events(possible: true).map(&:name).each do |event| %>
  <%= link_to event, [event, @job], method: :patch, class: "btn btn-default" %>
<% end %>
```

This will display a button for each event that is possible from that state. We’ve only used three lines of code and it will automatically adapt to any code changes in the model. Nice.

## Step 3 — Making it look nicer

We want the button captions to be different from the event names and we might want to use different styling on different buttons. We can make use of [Rails’ I18n](http://guides.rubyonrails.org/i18n.html) to translate those values:

```erb
<% job.aasm.events(possible: true).map(&:name).each do |event| %>
  <%= link_to t(event, scope: [:job, :event]), [event, job], method: :patch, class: "btn btn-#{t event, scope: [:job, :btn_class]}" %>
<% end %>
```

And then in the en.yml file:

```yml
en:
  job:
    event:
      run: "Run"
      sleep: "Shut Down"
      clean: "Clean"
    btn_class:
      run: "success"
      sleep: "danger"
      clean: "default"
```

## Step 4 — Not All Events

Say we introduce a new event called alarm which gets called if something bad happens. We don’t want operators to press the alarm button so we introduce the concept of operator_events and clean up our view a little in the process:

```rb
class Job
  include AASM
  aasm do
    state :sleeping, initial: true
    state :running, :cleaning, :broken
    event :run do
      transitions from: :sleeping, to: :running
    end
    event :clean do
      transitions from: :running, to: :cleaning
    end
    event :sleep do
     transitions from: [:running, :cleaning], to: :sleeping
    end
    event :alarm do
      transitions from: [:running], to: :broken
    end
    event :fix do
      transitions from: [:broken], to: :sleeping
    end
  end
  def operator_events
    aasm.events(possible: true).map(&:name) & %i[run clean sleep fix]
  end
end
```

Then the view is slightly simpler and will never show the alarm button:

```erb
<% @job.operator_events.each do |event| %>
  <%= link_to t(event, scope: [:job, :event]), [event, @job], method: :patch, class: "btn btn-#{t event, scope: [:job, :btn_class]}" %>
<% end %>
```

### Step 5 — Sharing

This code is a candidate for being a shared partial.

```erb
render 'shared/events', model: job
```

And the partial that we can reuse for other models:

```erb
<% model.operator_events.each do |event| %>
  <%= link_to t(event, scope: [model.class.name.underscore, :event]), [event, model], method: :patch, class: "btn btn-#{t event, scope: [model.class.name.underscore, :btn_class]}" %>
<% end %>
```

If this feels like a bit too much abstraction for you, at least go for step 1 and don’t make your view think about what event is possible from what state. Here it is again:

```erb
<% @job.aasm.events(possible: true).map(&:name).tap do |possible| %>
  <% if possible.include? :run %>
    <%= link_to "Run", run_job_path, method: :patch, class: "btn btn-success" %>
  <% end %>
  <% if possible.include? :clean %>
    <%= link_to "Clean", clean_job_path, method: :patch, class: "btn btn-default" %>
  <% end %>
  <% if possible.include? :sleep %>
    <%= link_to "Shut Down", sleep_job_path, method: :patch, class: "btn btn-danger" %>
  <% end %>
<% end %>
```

### See the code

You can see the code for this article in this [GitHub repository](https://github.com/billhorsman/aasm_btns) with each step as a separate commit.
