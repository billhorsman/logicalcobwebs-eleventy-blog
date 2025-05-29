require 'uri'
require 'net/http'
require "json"

# See https://developer.themoviedb.org/docs/getting-started

# 1. Fetch all pages from TMDB
# See https://developer.themoviedb.org/reference/list-details

list_id = 8291691
api_key = "eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiI3ODk1NmI5NjEzMGU5NmVhZTYwNDNjMDZiYjBjMTliNiIsIm5iZiI6MTcwODUwODMwMS4yNDg5OTk4LCJzdWIiOiI2NWQ1YzQ4ZGYyOWQ2NjAxN2RlOTRmMzgiLCJzY29wZXMiOlsiYXBpX3JlYWQiXSwidmVyc2lvbiI6MX0.TU6KG5MWHtgbsGRNDG786tlP22TY3y86-Q02sDUUkEs"
list = []

(1..10).each do |page|
  puts "Fetching page #{page}"
  url = URI("https://api.themoviedb.org/3/list/#{list_id}?language=en-US&page=#{page}")

  http = Net::HTTP.new(url.host, url.port)
  http.use_ssl = true

  request = Net::HTTP::Get.new(url)
  request["accept"] = "application/json"
  request["Authorization"] = "Bearer #{api_key}"

  response = http.request(request)
  items = JSON.parse(response.read_body)["items"]
  if items.any?
    list += items
  else
    puts "Page #{page} is empty"
    break
  end
end

puts "Found #{list.length} films"

# 2. Sort the films by date
list.sort_by! { |film| film["release_date"] }

# 3. Generate a slug for each item
list.each do |film|
  film["slug"] = film["title"].downcase.gsub(/[^a-z0-9 ]+/, "").gsub(" ", "-")
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
