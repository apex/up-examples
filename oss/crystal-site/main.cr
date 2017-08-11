require "http/server"
require "crouter"

class App < Crouter::Router
  get "/" do
    context.response.content_type = "text/html"
    context.response << <<-EOF
      <!DOCTYPE html>
      <html>
        <head>
          <meta charset="utf-8">
          <title>Hello</title>
          <link rel="stylesheet" href="style.css">
        </head>
        <body>
          <p>Hello World</p>
        </body>
      </html>
    EOF
  end

  get "/style.css" do
    context.response.content_type = "text/css"
    context.response << <<-EOF
      body {
        padding: 50px;
        font: 14px Helvetica;
      }
    EOF
  end
end

port = ENV["PORT"].to_i

server = HTTP::Server.new(port, [HTTP::LogHandler.new, App.new])
server.listen
