import os
from bottle import route, run, template

@route('/')
def hello():
    return "hello"

@route('/hello/<name>')
def hello(name):
    return template('Hello <b>{{name}}</b>!', name=name)

run(port=int(os.environ.get('PORT', 8080)))
