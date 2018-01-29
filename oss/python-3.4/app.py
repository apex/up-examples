from http.server import BaseHTTPRequestHandler, HTTPServer
import os

class handler(BaseHTTPRequestHandler):
  def do_GET(self):
    self.send_response(200)
    self.send_header('Content-type', 'text/plain')
    self.end_headers()
    self.wfile.write(bytes("Hello World from Python 3.4", "utf8"))
    return

addr = ('127.0.0.1', int(os.environ['PORT']))
httpd = HTTPServer(addr, handler)
httpd.serve_forever()
