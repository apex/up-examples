up

curl -s `up url`
# contains: A Programming Language

curl -s -D- `up url`/home.css
# contains: Content-Type: text/css; charset=utf-8

curl -s -D- `up url`/guide/guide.html
# contains: Installing
