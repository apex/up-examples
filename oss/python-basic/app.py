from BaseHTTPServer import BaseHTTPRequestHandler,HTTPServer
import os

class myHandler(BaseHTTPRequestHandler):
	def do_GET(self):
		self.send_response(200)
		self.send_header('Content-type','text/html')
		self.end_headers()
		self.wfile.write("Hello World from Python\n")
		return

server = HTTPServer(('', int(os.environ['PORT'])), myHandler)
server.serve_forever()
