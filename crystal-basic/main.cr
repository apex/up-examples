require "http/server"

port = ENV["PORT"].to_i

server = HTTP::Server.new(port) do |ctx|
  ctx.response.content_type = "text/plain"
  ctx.response.print "Hello world from Crystal"
end

server.listen
