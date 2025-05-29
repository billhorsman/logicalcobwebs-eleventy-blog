require 'uri'
require 'net/http'
require "json"

# See https://developer.themoviedb.org/docs/getting-started

# 1. Fetch all pages from TMDB
# See https://developer.themoviedb.org/reference/list-details

list_id = 8291691
list = []

def get_api_response(endpoint:, page: nil, language: "en-US")
  api_key = "eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiI3ODk1NmI5NjEzMGU5NmVhZTYwNDNjMDZiYjBjMTliNiIsIm5iZiI6MTcwODUwODMwMS4yNDg5OTk4LCJzdWIiOiI2NWQ1YzQ4ZGYyOWQ2NjAxN2RlOTRmMzgiLCJzY29wZXMiOlsiYXBpX3JlYWQiXSwidmVyc2lvbiI6MX0.TU6KG5MWHtgbsGRNDG786tlP22TY3y86-Q02sDUUkEs"

  url = URI("https://api.themoviedb.org/3/#{endpoint}?language=#{language}#{"&page=#{page}" if page}")

  http = Net::HTTP.new(url.host, url.port)
  http.use_ssl = true

  request = Net::HTTP::Get.new(url)
  request["accept"] = "application/json"
  request["Authorization"] = "Bearer #{api_key}"

  http.request(request)
end

if ARGV[0] == "refresh"
  (1..10).each do |page|
    puts "Fetching page #{page}"
    response = get_api_response(endpoint: "list/#{list_id}", page:)
    items = JSON.parse(response.read_body)["items"]
    if items.any?
      list += items
    else
      puts "Page #{page} is empty"
      break
    end
  end

  # 2. Sort the films by date
  list.sort_by! { |film| film["release_date"] }

  # 3. Generate a slug for each item
  list.each do |film|
    film["slug"] = film["title"].downcase.gsub(/[^a-z0-9 ]+/, "").gsub(" ", "-")
  end

  # 3.5. Get the major credits for each film
  list.each do |film|
    filename = "tmp/#{film["slug"]}-credits.json"
    if !File.exist?(filename)
      response = get_api_response(endpoint: "movie/#{film["id"]}/credits")
      IO.write(filename, response.read_body)
    end
    credits = JSON.parse(IO.read(filename))
    film["director"] = credits.dig("crew").select { |crew| crew["job"] == "Director" }.map { |crew| crew["name"] }.join(", ")
    film["cast"] = credits.dig("cast").map { |cast| 
      {
        name: cast["name"],
        character: cast["character"]
      }
    }
  end

  # 4. Download backdrop and poster images
  list.each do |film|
    %w[poster backdrop].each do |type|
      if image_slug = film["#{type}_path"]
        path = "#{type}s/#{film["slug"]}.jpg"
        full_path = "../content/bill/films/#{path}"
        if !File.exist?(full_path)
          url = "https://image.tmdb.org/t/p/w500#{image_slug}"
          response = Net::HTTP.get(URI(url))
          puts "Writing #{response.length} bytes to #{path}"
          File.write(full_path, response)    
        end
        film[type] = path
      end
    end
  end

  # 5. Write the list to global data
  IO.write("../_data/films.json", JSON.pretty_generate(list:))

else
  list = JSON.parse(File.read("../_data/films.json"))["list"]
end

# 6. Write a page for each film

list.each do |film|
  path = "../content/bill/films/#{film["slug"]}.md"
  File.write(path, <<~MD)
    ---
    title: "#{film["title"]}"
    layout: layouts/home.njk
    ---

    <a href="../">Back to all films</a>
    
    <article class="film">
      <h1>#{film["title"]} (#{film["release_date"].split("-").first})</h1>

      <p class="director">
        Directed by <strong>#{film["director"]}</strong>
      </p>

      <img src="../films/#{film["poster"]}" alt="">

      <h2>
        Cast
      </h2>
      <ul>
        #{film["cast"].map { |cast| "<li><strong>#{cast["name"]}</strong> as <em>#{cast["character"]}</em></li>" }.join("\n")}
      </ul>
    </article>
  MD
end
