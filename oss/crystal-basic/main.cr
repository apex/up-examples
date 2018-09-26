require "http/server"

port = ENV["PORT"].to_i

server = HTTP::Server.new do |ctx|
  ctx.response.content_type = "text/plain"
  ctx.response.print "Hello World from Crystal"
end

address = server.bind_tcp port

puts "Listening on http://#{address}"
server.listen
